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

type Config struct {
	c client.Client
}

func main() {
	cfg := Config{
		c: *client.NewClient(5 * time.Second),
	}

	url := getUrlArg()

	pages := make(map[string]int)
	fmt.Println(cfg.crawlPage(url, url, pages))
	fmt.Println(pages)
}
