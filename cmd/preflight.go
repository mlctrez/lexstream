package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"golang.org/x/mod/semver"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

// Check for pre-requisites
func main() {
	execute(
		checkNodeVersion,
		checkAlexaSkillsKit,
		checkAwsConfig,
		checkAwsCli,
	)
}

func execute(functions ...func()) {
	for _, f := range functions {
		vop := reflect.ValueOf(f).Pointer()
		name := runtime.FuncForPC(vop).Name()
		fmt.Println(strings.TrimPrefix(name, "main."))
		f()
	}
}

func checkAlexaSkillsKit() {

	workspace := ".ask"
	ask := filepath.Join(workspace, "node_modules", ".bin", "ask")

	if _, err := os.Stat(ask); os.IsNotExist(err) {
		exitOn("", os.MkdirAll(workspace, 0755))

		fmt.Println("installing alexa skills, should take less than a minute")
		command := exec.Command("npm", "install", "--save-dev", "ask-cli")
		command.Dir = workspace

		output, installError := command.CombinedOutput()
		if installError != nil {
			fmt.Println(string(output))
			exitOn("", installError)
		}

		// fix bad @src in this file to allow ask util to work
		//  .ask/node_modules/ask-cli/lib/commands/util/util-commander.js
		path := ".ask/node_modules/ask-cli/lib/commands/util/util-commander.js"
		fileData := exitOn(os.ReadFile(path))
		replaced := strings.ReplaceAll(string(fileData), "@src/commands/util", "./")
		exitOn("", os.WriteFile(path, []byte(replaced), 0664))

	}

	askConfig := filepath.Join(userHome(), ".ask", "cli_config")
	if _, err := os.Stat(askConfig); os.IsNotExist(err) {
		configDoc := "https://developer.amazon.com/en-US/docs/alexa/smapi/ask-cli-command-reference.html#configure-command"
		exitOn("", fmt.Errorf("ask configuration file missing\n  %s\n  %s", configDoc, ask+" configure"))
	}

	output, err := exec.Command(ask, "smapi", "get-vendor-list").CombinedOutput()
	if err != nil {
		exitOn(output, err)
	}

	output, err = exec.Command(ask, "util").CombinedOutput()
	if err != nil {
		exitOn(output, err)
	}

}

func checkNodeVersion() {

	minVersion := "v8.3"
	command := exec.Command("node", "--version")
	nodeVersion := strings.TrimSpace(string(exitOn(command.CombinedOutput())))

	if semver.Compare(nodeVersion, minVersion) < 0 {
		exitOn("", fmt.Errorf("minimum node version %s - current version %s", minVersion, nodeVersion))
	}
}

func checkAwsCli() {
	exitOn(exec.Command("aws", "--version").CombinedOutput())
}

func checkAwsConfig() {
	ctx := context.TODO()

	cfg := exitOn(config.LoadDefaultConfig(ctx))

	exitOn(s3.NewFromConfig(cfg).ListBuckets(ctx, nil))

	exitOn(lambda.NewFromConfig(cfg).ListFunctions(ctx, nil))

	exitOn(iam.NewFromConfig(cfg).GetAccountSummary(ctx, nil))
}

func userHome() string {
	return exitOn(os.UserHomeDir())
}

func exitOn[K any](a K, err error) K {
	if err != nil {
		pc, file, no, ok := runtime.Caller(1)
		dir, wdErr := os.Getwd()
		if ok && wdErr == nil {
			file = strings.TrimPrefix(strings.TrimPrefix(file, dir), string(os.PathSeparator))
			funcName := runtime.FuncForPC(pc).Name()
			fmt.Printf("%s:%d %s : %+v \n", file, no, funcName, err)
			os.Exit(-1)
		} else {
			panic(err)
		}
	}
	return a
}
