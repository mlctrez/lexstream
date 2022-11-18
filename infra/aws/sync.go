package main

import (
	"archive/zip"
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	iamTypes "github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	lambdaTypes "github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	_ "github.com/aws/aws-sdk-go-v2/service/s3/types"
	s3Types "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/mlctrez/bolt"
	"github.com/mlctrez/lexstream/internal/settings"
	"io"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"time"
)

var assumeRoleDoc = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}`

const AWSManaged = "arn:aws:iam::aws:policy/"

var RolePolicies = []string{"CloudWatchLogsFullAccess", "AmazonS3FullAccess", "AmazonSQSFullAccess"}

const LambdaRole = "lexstream_role"

type Infra struct {
	settings *settings.Settings
	cfg      aws.Config
}

func (i *Infra) Setup() (err error) {

	if i.cfg, err = config.LoadDefaultConfig(context.TODO()); err != nil {
		return err
	}

	i.settings, err = settings.Load()

	return
}

func (i *Infra) SetupRole() (err error) {
	iac := iam.NewFromConfig(i.cfg)
	roleName := aws.String(LambdaRole)

	ctx := context.TODO()

	_, err = iac.GetRole(ctx, &iam.GetRoleInput{RoleName: roleName})

	var nse *iamTypes.NoSuchEntityException
	if errors.As(err, &nse) {
		_, err = iac.CreateRole(ctx, &iam.CreateRoleInput{
			AssumeRolePolicyDocument: &assumeRoleDoc,
			RoleName:                 roleName,
		})
	}
	if err != nil {
		return err
	}

	_, err = iac.UpdateAssumeRolePolicy(ctx, &iam.UpdateAssumeRolePolicyInput{RoleName: roleName, PolicyDocument: &assumeRoleDoc})
	if err != nil {
		return err
	}
	foundPolicyNames := make(map[string]bool)
	paginator := iam.NewListRolePoliciesPaginator(iac, &iam.ListRolePoliciesInput{RoleName: roleName})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		for _, name := range page.PolicyNames {
			foundPolicyNames[name] = true
		}
	}

	for _, policyName := range RolePolicies {
		if !foundPolicyNames[policyName] {
			policyInput := &iam.AttachRolePolicyInput{
				RoleName:  roleName,
				PolicyArn: aws.String(AWSManaged + policyName),
			}
			_, err = iac.AttachRolePolicy(ctx, policyInput)
			if err != nil {
				return err
			}
		}
	}

	if err != nil {
		return err
	}

	return nil

}

func (i *Infra) SetupLambda() (err error) {
	ctx := context.TODO()

	lc := lambda.NewFromConfig(i.cfg)
	var role *iam.GetRoleOutput
	role, err = iam.NewFromConfig(i.cfg).GetRole(ctx, &iam.GetRoleInput{RoleName: aws.String(LambdaRole)})
	if err != nil {
		return err
	}

	functionName := aws.String("lexstream_lambda")
	var functionOutput *lambda.GetFunctionOutput
	functionOutput, err = lc.GetFunction(ctx, &lambda.GetFunctionInput{FunctionName: functionName})
	var nfe *lambdaTypes.ResourceNotFoundException
	bucket := i.settings.Bucket

	if errors.As(err, &nfe) {
		err = nil
		fmt.Println("creating lambda")

		_, err = lc.CreateFunction(ctx, &lambda.CreateFunctionInput{
			Code: &lambdaTypes.FunctionCode{
				S3Bucket: aws.String(bucket),
				S3Key:    aws.String("lexstream.zip"),
			},
			Description: aws.String("LexStream handles alexa media skill requests"),
			Environment: &lambdaTypes.Environment{Variables: map[string]string{
				"LEXSTREAM_BUCKET_NAME": bucket,
			}},
			FunctionName: functionName,
			Handler:      aws.String("lexstream"),
			PackageType:  lambdaTypes.PackageTypeZip,
			Publish:      true,
			Role:         role.Role.Arn,
			Runtime:      lambdaTypes.RuntimeGo1x,
		})
		if err != nil {
			return
		}
	} else {

		environment := functionOutput.Configuration.Environment
		if lb, ok := environment.Variables["LEXSTREAM_BUCKET_NAME"]; ok {
			if lb != bucket {
				fmt.Println("updating function configuration")
				_, err = lc.UpdateFunctionConfiguration(ctx, &lambda.UpdateFunctionConfigurationInput{
					FunctionName: functionName,
					Environment: &lambdaTypes.Environment{Variables: map[string]string{
						"LEXSTREAM_BUCKET_NAME": bucket,
					}},
				})
				if err != nil {
					return
				}
			}
		}

		fmt.Println("updating function code")
		// this needs to re-try if above update occurred
		for t := 0; t < 10; t++ {
			_, err = lc.UpdateFunctionCode(ctx, &lambda.UpdateFunctionCodeInput{
				FunctionName: functionName,
				S3Bucket:     aws.String(bucket),
				S3Key:        aws.String("lexstream.zip"),
			})
			var rce *lambdaTypes.ResourceConflictException
			if errors.As(err, &rce) {
				fmt.Println("sleeping due to resource conflict exception")
				time.Sleep(time.Second * 4)
				continue
			} else {
				break
			}
		}

	}
	if err != nil {
		return
	}
	functionOutput, err = lc.GetFunction(ctx, &lambda.GetFunctionInput{FunctionName: functionName})
	if err != nil {
		return
	}

	_, err = lc.RemovePermission(ctx, &lambda.RemovePermissionInput{
		FunctionName: functionName,
		RevisionId:   functionOutput.Configuration.RevisionId,
		StatementId:  aws.String("alexa-skill-invoke-policy"),
	})

	var rnf *lambdaTypes.ResourceNotFoundException
	if err != nil && !errors.As(err, &rnf) {
		return
	}

	_, err = lc.AddPermission(ctx, &lambda.AddPermissionInput{
		Action:           aws.String("lambda:InvokeFunction"),
		EventSourceToken: aws.String(i.settings.SkillId),
		FunctionName:     functionOutput.Configuration.FunctionName,
		Principal:        aws.String("alexa-appkit.amazon.com"),
		StatementId:      aws.String("alexa-skill-invoke-policy"),
	})
	if err != nil {
		return
	}

	fmt.Println("lambda arn for alexa skills endpoint:", *functionOutput.Configuration.FunctionArn)

	return
}

func (i *Infra) SetupBucket() (err error) {
	s3c := s3.NewFromConfig(i.cfg)

	bucket := aws.String(i.settings.Bucket)
	_, err = s3c.CreateBucket(context.TODO(), &s3.CreateBucketInput{Bucket: bucket})
	if err != nil {
		return err
	}
	_, err = s3c.PutPublicAccessBlock(context.TODO(), &s3.PutPublicAccessBlockInput{
		Bucket: bucket,
		PublicAccessBlockConfiguration: &s3Types.PublicAccessBlockConfiguration{
			BlockPublicPolicy:     true,
			BlockPublicAcls:       true,
			RestrictPublicBuckets: true,
			IgnorePublicAcls:      true,
		},
	})
	if err != nil {
		return err
	}

	return

}

type Provider struct {
}

func (p *Provider) Key() bolt.Key          { return "keyOne" }
func (p *Provider) Value() ([]byte, error) { return []byte("valueOne"), nil }

func (i *Infra) BuildCode() (err error) {

	cmd := exec.Command("go", "build", "-o", "temp/lexstream", "skill/bin/lambda.go")
	cmd.Env = []string{"GOARCH=amd64", "GOOS=linux"}
	for _, s := range os.Environ() {
		cmd.Env = append(cmd.Env, s)
	}
	var out []byte
	if out, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(string(out))
		return
	}

	b, err := bolt.New("temp/bolt.db")
	if err != nil {
		return err
	}
	boltBucket := bolt.Key("sample")
	err = b.CreateBuckets(bolt.Keys{boltBucket})
	if err != nil {
		return err
	}
	err = b.Put(boltBucket, &bolt.Value{K: "keyOne", V: []byte("valueOne")})
	if err != nil {
		return err
	}
	err = b.Close()
	if err != nil {
		return err
	}

	var archive *os.File
	if archive, err = os.Create("temp/lexstream.zip"); err != nil {
		return
	}
	defer archive.Close()
	zipWriter := zip.NewWriter(archive)
	defer zipWriter.Close()

	if err = writeFile("lexstream", zipWriter); err != nil {
		return
	}
	if err = writeFile("bolt.db", zipWriter); err != nil {
		return
	}

	return err
}

func writeFile(name string, zipWriter *zip.Writer) (err error) {

	var stat os.FileInfo
	stat, err = os.Stat(fmt.Sprintf("temp/%s", name))
	if err != nil {
		return err
	}
	var header *zip.FileHeader
	header, err = zip.FileInfoHeader(stat)
	if err != nil {
		return err
	}

	var writer io.Writer
	writer, err = zipWriter.CreateHeader(header)
	if err != nil {
		return
	}
	var reader *os.File
	reader, err = os.Open(fmt.Sprintf("temp/%s", name))
	defer func() { _ = reader.Close() }()
	_, err = io.Copy(writer, reader)
	return err
}

func (i *Infra) UploadCode() (err error) {

	var archive *os.File
	if archive, err = os.Open("temp/lexstream.zip"); err != nil {
		return
	}
	defer archive.Close()

	s3c := s3.NewFromConfig(i.cfg)
	_, err = s3c.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(i.settings.Bucket),
		Key:    aws.String("lexstream.zip"),
		Body:   archive,
	})
	return
}

func (i *Infra) Cleanup() (err error) {
	return os.RemoveAll("temp")
}

func execute(functions ...func() error) {
	for _, f := range functions {
		fmt.Println("execute", runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name())
		err := f()
		if err != nil {
			panic(err)
		}
	}
}

func main() {

	i := &Infra{}

	execute(
		i.Setup,
		i.SetupBucket,
		i.BuildCode,
		i.UploadCode,
		i.SetupRole,
		i.SetupLambda,
		i.Cleanup,
	)
}
