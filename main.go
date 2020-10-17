package main

import (
	"os"

	"github.com/mizzy/para/cmd"
)

func main() {
	err := cmd.Run(os.Args[1:])
	if err != nil {
		os.Exit(1)
	}
}
