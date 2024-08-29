package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="/path/one">
					<span>Boot.dev</span>
				</a>
				<a href="https://other.com/path/one">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "absolute and relative urls for test",
			inputURL: "https://t.test",
			inputBody: `
		<html>
			<body>
				<a href="/test">
					<span>Boot.dev</span>
				</a>
				<a href="/">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://t.test/test", "https://t.test/"},
		},
		{
			name:     "absolute and relative urls for test with dots",
			inputURL: "https://t.test",
			inputBody: `
		<html>
			<body>
				<a href="../../test">
					<span>Boot.dev</span>
				</a>
				<a href="../../">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://t.test/test", "https://t.test/"},
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			actual, err := getURLsFromHTML(tt.inputBody, tt.inputURL)

			assert.Nil(err, "Test: %v '%s' FAIL: unexpected error: %v", tt.name, err)

			assert.Equal(tt.expected, actual, fmt.Sprintf(
				"Test %v - '%s' FAIL: expected result: %v, actual: %v", i, tt.name,
				"["+strings.Join(tt.expected, ", ")+"]",
				"["+strings.Join(actual, ", ")+"]",
			))
		})
	}
}

func TestGetURLsFromHTMLErrors(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		inputBody     string
		errorContains string
	}{
		{
			name: "handle invalid input url",
			inputBody: `
<html>
	<body>
		<a href="/path">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			inputURL:      "://invalid url",
			errorContains: "could not parse base url",
		},
		{
			name: "handle invalid href",
			inputBody: `
<html>
	<body>
		<a href="://invalid url">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			inputURL:      "https://t.test",
			errorContains: "could not parse href",
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			_, err := getURLsFromHTML(tt.inputBody, tt.inputURL)

			assert.NotNil(err, "Test: %v fail expected error to not be nil", i)

			assert.Contains(err.Error(), tt.errorContains, "Test: %v '%v' expected error to contain: %v but got: %v", i, tt.name, tt.errorContains, err.Error())
		})
	}
}
