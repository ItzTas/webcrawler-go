package main

import (
	"fmt"
	"net/url"
)

func (c *Config) crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) error {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return err
	}
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		return err
	}

	if baseURL.Hostname() != currentURL.Hostname() {
		return nil
	}

	normCurr, err := normalizeURL(rawCurrentURL)
	if err != nil {
		return err
	}

	if _, ok := pages[normCurr]; ok {
		pages[normCurr] += 1
		return nil
	}

	pages[normCurr] = 1

	html, err := c.c.GetHTML(rawCurrentURL)
	if err != nil {
		return nil
	}

	fmt.Printf("\ngot html of url: %v\n", currentURL)

	urls, err := getURLsFromHTML(html, rawBaseURL)
	if err != nil {
		return err
	}

	for _, u := range urls {
		err := c.crawlPage(rawBaseURL, u, pages)
		if err != nil {
			return err
		}
	}
	return nil
}
