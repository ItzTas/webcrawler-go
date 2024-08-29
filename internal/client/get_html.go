package client

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (c *Client) GetHTML(rawURL string) (string, error) {
	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", err
	}

	res, err := c.c.Do(req)
	if err != nil {
		return "", err
	}

	if res.StatusCode >= 400 {
		return "", fmt.Errorf("got error status code: %v", res.StatusCode)
	}

	if contType := res.Header.Get("Content-type"); !strings.Contains(contType, "text/html") {
		return "", fmt.Errorf("expected html content type but got: %v", contType)
	}

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(dat), nil
}
