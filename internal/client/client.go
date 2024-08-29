package client

import (
	"net/http"
	"time"
)

type Client struct {
	c http.Client
}

func NewClient(timeout time.Duration) *Client {
	c := Client{
		c: http.Client{
			Timeout: timeout,
		},
	}

	return &c
}
