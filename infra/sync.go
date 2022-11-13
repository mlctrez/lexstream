package main

import (
	"archive/zip"
	"context"
	"encoding/json"
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
	"github.com/mlctrez/lexstream/smapi"
	"io"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"strings"
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

var RolePolicies = []string{"CloudWatchLogsFullAccess", "AmazonS3FullAccess"}

const LambdaRole = "lexstream_role"

var settings = &SettingsJson{}

type SettingsJson struct {
	Bucket    string `json:"bucket"`
	SkillName string `json:"skill_name"`
	SkillId   string `json:"-"`
}

var cfg aws.Config

func Setup() (err error) {

	if cfg, err = config.LoadDefaultConfig(context.TODO()); err != nil {
		return err
	}

	var open *os.File
	if open, err = os.Open("settings.json"); err != nil {
		return
	}
	defer open.Close()
	err = json.NewDecoder(open).Decode(settings)
	if err != nil {
		return
	}
	if settings.Bucket == "" {
		return fmt.Errorf("no bucket name in settings.json")
	}

	if settings.SkillId == "" {
		id, getSkillIdErr := smapi.GetSkillIdForName(settings.SkillName)
		if getSkillIdErr != nil {
			return getSkillIdErr
		}
		settings.SkillId = id
	}
	if !strings.HasPrefix(settings.SkillId, "amzn1.ask.skill.") {
		return fmt.Errorf("skill id does not appear to be valid in settings.json")
	}

	return
}

func SetupRole() (err error) {
	iac := iam.NewFromConfig(cfg)
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
			_, err = iac.AttachRolePolicy(ctx, &iam.AttachRolePolicyInput{RoleName: roleName, PolicyArn: aws.String(AWSManaged + policyName)})
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

func SetupLambda() (err error) {
	ctx := context.TODO()

	lc := lambda.NewFromConfig(cfg)
	var role *iam.GetRoleOutput
	role, err = iam.NewFromConfig(cfg).GetRole(ctx, &iam.GetRoleInput{RoleName: aws.String(LambdaRole)})
	if err != nil {
		return err
	}

	functionName := aws.String("lexstream_lambda")
	var functionOutput *lambda.GetFunctionOutput
	functionOutput, err = lc.GetFunction(ctx, &lambda.GetFunctionInput{FunctionName: functionName})
	var nfe *lambdaTypes.ResourceNotFoundException
	if errors.As(err, &nfe) {
		err = nil
		fmt.Println("creating lambda")

		_, err = lc.CreateFunction(ctx, &lambda.CreateFunctionInput{
			Code: &lambdaTypes.FunctionCode{
				S3Bucket: aws.String(settings.Bucket),
				S3Key:    aws.String("lexstream.zip"),
			},
			Description: aws.String("LexStream handles alexa media skill requests"),
			Environment: &lambdaTypes.Environment{Variables: map[string]string{
				"LEXSTREAM_BUCKET_NAME": settings.Bucket,
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
			if lb != settings.Bucket {
				fmt.Println("updating function configuration")
				_, err = lc.UpdateFunctionConfiguration(ctx, &lambda.UpdateFunctionConfigurationInput{
					FunctionName: functionName,
					Environment: &lambdaTypes.Environment{Variables: map[string]string{
						"LEXSTREAM_BUCKET_NAME": settings.Bucket,
					}},
				})
				if err != nil {
					return
				}
			}
		}

		fmt.Println("updating function code")
		// this needs to re-try if above update occurred
		for i := 0; i < 10; i++ {
			_, err = lc.UpdateFunctionCode(ctx, &lambda.UpdateFunctionCodeInput{
				FunctionName: functionName,
				S3Bucket:     aws.String(settings.Bucket),
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
		EventSourceToken: aws.String(settings.SkillId),
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

func SetupBucket() (err error) {
	s3c := s3.NewFromConfig(cfg)

	bucket := aws.String(settings.Bucket)
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

func BuildCode() (err error) {

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

	var archive *os.File
	if archive, err = os.Create("temp/lexstream.zip"); err != nil {
		return
	}
	defer archive.Close()
	zipWriter := zip.NewWriter(archive)
	defer zipWriter.Close()

	var writer io.Writer
	writer, err = zipWriter.Create("lexstream")
	if err != nil {
		return
	}
	var reader *os.File
	reader, err = os.Open("temp/lexstream")
	defer reader.Close()
	_, err = io.Copy(writer, reader)

	return err
}

func UploadCode() (err error) {

	var archive *os.File
	if archive, err = os.Open("temp/lexstream.zip"); err != nil {
		return
	}
	defer archive.Close()

	s3c := s3.NewFromConfig(cfg)
	_, err = s3c.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(settings.Bucket),
		Key:    aws.String("lexstream.zip"),
		Body:   archive,
	})
	return
}

func Cleanup() (err error) {
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

	execute(
		Setup,
		SetupBucket,
		BuildCode,
		UploadCode,
		SetupRole,
		SetupLambda,
		Cleanup,
	)
}
