package main

import (
	"net/url"
	"strings"
)

func normalizeURL(urlstr string) (string, error) {
	u, err := url.Parse(urlstr)
	if err != nil {
		return "", err
	}

	norm := u.Host + u.Path

	norm = strings.TrimRight(norm, "/")

	return norm, nil
}
