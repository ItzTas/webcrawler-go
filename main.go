package main

import (
	"fmt"
	"os"
)

func getArgs() []string {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	return args
}

func main() {
	args := getArgs()

	fmt.Println(args)
}
