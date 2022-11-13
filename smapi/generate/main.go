package main

import (
	"github.com/mlctrez/swaggerlt"
	"regexp"
)

func main() {
	ops := &swaggerlt.Options{
		PathRegex:   regexp.MustCompile("/(v0|v1)"),
		SpecFile:    "smapi/spec.json",
		ServiceName: "smapi",
		ModuleName:  "github.com/mlctrez/lexstream",
	}
	generator, err := swaggerlt.New(ops)
	if err != nil {
		panic(err)
	}
	err = generator.Execute()
	if err != nil {
		panic(err)
	}
}
