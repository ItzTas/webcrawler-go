package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ItzTas/webcrawler-go/internal/client"
)

func getUrlArg() string {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	return args[1]
}

func main() {
	c := client.NewClient(5 * time.Second)

	url := getUrlArg()

	fmt.Println(c.GetHTML(url))
}
