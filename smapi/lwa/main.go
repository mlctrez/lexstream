package main

import (
	"fmt"
	"github.com/mlctrez/lexstream/smapi"
)

func main() {

	err := smapi.GenerateLwaTokens()
	if err != nil {
		fmt.Println(err)
	}

}
