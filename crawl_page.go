package main

import (
	"net/url"
)

func (cfg *Config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	if _, isFirst = cfg.pages[normalizedURL]; isFirst {
		cfg.pages[normalizedURL] += 1
		return false
	}

	cfg.pages[normalizedURL] = 1
	return true
}

func (cfg *Config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	cfg.mu.Lock()
	if len(cfg.pages) >= cfg.maxPages {
		cfg.mu.Unlock()
		return
	}
	cfg.mu.Unlock()

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		return
	}

	if cfg.baseURL.Hostname() != currentURL.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		return
	}

	if !cfg.addPageVisit(normalizedURL) {
		return
	}

	html, err := cfg.c.GetHTML(rawCurrentURL)
	if err != nil {
		return
	}

	urls, err := getURLsFromHTML(html, cfg.baseURL.String())
	if err != nil {
		return
	}

	for _, u := range urls {
		cfg.wg.Add(1)
		go cfg.crawlPage(u)
	}
	return
}
