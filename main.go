package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func getArgs() (string, int, int) {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(args) < 4 {
		fmt.Println("not sufficient args want: url, maxConcurrency, maxPages")
		os.Exit(1)
	}
	maxConcorrency, err := strconv.Atoi(args[2])
	if err != nil {
		log.Fatal(err)
	}
	maxPages, err := strconv.Atoi(args[3])
	if err != nil {
		log.Fatal(err)
	}
	return args[1], maxConcorrency, maxPages
}

func main() {
	url, maxConcurrency, maxPages := getArgs()

	cfg, err := newConfig(5*time.Second, url, maxConcurrency, maxPages)
	if err != nil {
		log.Fatal(err)
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(url)
	cfg.wg.Wait()

	for k, v := range cfg.pages {
		fmt.Printf("%s, %d\n", k, v)
	}
}
