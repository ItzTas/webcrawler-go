package main

import (
	"net/url"
	"sync"
	"time"

	"github.com/ItzTas/webcrawler-go/internal/client"
)

type Config struct {
	c                  client.Client
	pages              map[string]int
	maxPages           int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func newConfig(timeout time.Duration, rawBaseURL string, maxConcurrency, maxPages int) (*Config, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, err
	}

	cfg := Config{
		c:                  *client.NewClient(timeout),
		pages:              make(map[string]int),
		maxPages:           maxPages,
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
	}

	return &cfg, nil
}
