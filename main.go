package main

import (
	"fmt"
	"os"

	"github.com/mizzy/para/cmd"
)

const version = "0.2.0"

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: para COMMAND ARG0 [arg1 arg2 ...]")
		os.Exit(1)
	}

	err := cmd.Run(os.Args[1:])
	if err != nil {
		os.Exit(1)
	}
}
