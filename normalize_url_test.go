package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "remove scheme",
			input:    "https://t.test/path",
			expected: "t.test/path",
		},
		{
			name:     "remove path trailing slash",
			input:    "https://t.test/path/",
			expected: "t.test/path",
		},
		{
			name:     "remove name trailing slash ",
			input:    "https://t.test/",
			expected: "t.test",
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := normalizeURL(tt.input)

			assert.Nil(t, err, fmt.Sprintf("Test: %v '%s' FAIL: unexpected error: %v", i, tt.name, err))

			assert.Equal(t, tt.expected, actual, fmt.Sprintf("Test %v - '%s' FAIL: expected URL: %v, actual: %v", i, tt.name, tt.expected, actual))
		})
	}
}
