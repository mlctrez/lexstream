package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mlctrez/lexstream/skill"
)

func main() {
	lambda.Start(skill.Handle)
}
