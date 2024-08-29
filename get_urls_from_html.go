package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	nodes, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return []string{}, fmt.Errorf("could not parse body: %v", err)
	}

	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return []string{}, fmt.Errorf("could not parse base url: %v", err)
	}
	var links []string
	var f func(n *html.Node) error
	f = func(n *html.Node) error {
		if n.Type == html.ElementNode && n.DataAtom == atom.A {
			for _, a := range n.Attr {
				if a.Key == "href" {
					toApp, err := url.Parse(a.Val)
					if err != nil {
						return fmt.Errorf("could not parse href: %v error: %v", a.Val, err)
					}
					toApp = baseURL.ResolveReference(toApp)
					links = append(links, toApp.String())
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if err := f(c); err != nil {
				return err
			}
		}
		return nil
	}
	err = f(nodes)
	if err != nil {
		return []string{}, err
	}

	return links, nil
}
