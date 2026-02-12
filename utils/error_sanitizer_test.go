/**
 * Copyright 2024-present Coinbase Global, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package utils

import (
	"testing"
)

func TestIsHTMLResponse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: false,
		},
		{
			name:     "Plain text",
			input:    "Simple error message",
			expected: false,
		},
		{
			name:     "JSON response",
			input:    `{"error": "rate limit exceeded"}`,
			expected: false,
		},
		{
			name:     "HTML document with DOCTYPE",
			input:    `<!DOCTYPE html><html><head><title>Error</title></head></html>`,
			expected: true,
		},
		{
			name:     "HTML document without DOCTYPE",
			input:    `<html><head><title>Coinbase</title></head><body>Error page</body></html>`,
			expected: true,
		},
		{
			name:     "HTML fragment with multiple tags",
			input:    `<div><p>Error occurred</p><span>Details</span></div>`,
			expected: true,
		},
		{
			name:     "Single HTML tag",
			input:    `<p>Single tag</p>`,
			expected: false,
		},
		{
			name: "Coinbase rate limit HTML (truncated example)",
			input: `<html>
  <head>
    <title>Coinbase</title>
    <meta name="robots" content="noindex">
    <style>body { font-family: Arial; }</style>
  </head>
  <body>
    <h1>Rate Limit Exceeded</h1>
  </body>
</html>`,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsHTMLResponse(tt.input)
			if result != tt.expected {
				t.Errorf("IsHTMLResponse(%q) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestExtractErrorMessage(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Plain text - no change",
			input:    "Simple error message",
			expected: "Simple error message",
		},
		{
			name:     "HTML with meaningful title",
			input:    `<html><head><title>Rate Limit Exceeded</title></head><body>Content</body></html>`,
			expected: "Rate Limit Exceeded",
		},
		{
			name:     "HTML with Coinbase title (should fallback)",
			input:    `<html><head><title>Coinbase</title></head><body>Error content</body></html>`,
			expected: "request failed",
		},
		{
			name:     "HTML without title",
			input:    `<html><head></head><body><h1>Error</h1></body></html>`,
			expected: "request failed",
		},
		{
			name:     "HTML with empty title",
			input:    `<html><head><title></title></head><body>Error</body></html>`,
			expected: "request failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExtractErrorMessage(tt.input)
			if result != tt.expected {
				t.Errorf("ExtractErrorMessage(%q) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestSanitizeRateLimitError(t *testing.T) {
	tests := []struct {
		name         string
		errorMessage string
		statusCode   int
		expected     string
	}{
		{
			name:         "HTTP 429 - always returns rate limit message",
			errorMessage: "Any message",
			statusCode:   429,
			expected:     "rate limit exceeded",
		},
		{
			name:         "HTTP 429 with HTML - still returns rate limit message",
			errorMessage: `<html><head><title>Error</title></head><body>Content</body></html>`,
			statusCode:   429,
			expected:     "rate limit exceeded",
		},
		{
			name:         "HTTP 400 with HTML",
			errorMessage: `<html><head><title>Bad Request</title></head><body>Content</body></html>`,
			statusCode:   400,
			expected:     "Bad Request",
		},
		{
			name:         "HTTP 500 with generic HTML",
			errorMessage: `<html><head><title>Coinbase</title></head><body>Error</body></html>`,
			statusCode:   500,
			expected:     "internal server error",
		},
		{
			name:         "HTTP 401 with plain text",
			errorMessage: "Unauthorized access",
			statusCode:   401,
			expected:     "Unauthorized access",
		},
		{
			name:         "Unknown status code with HTML",
			errorMessage: `<html><head><title>Custom Error</title></head><body>Content</body></html>`,
			statusCode:   418,
			expected:     "Custom Error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SanitizeRateLimitError(tt.errorMessage, tt.statusCode)
			if result != tt.expected {
				t.Errorf("SanitizeRateLimitError(%q, %d) = %q, expected %q", tt.errorMessage, tt.statusCode, result, tt.expected)
			}
		})
	}
}

func TestSanitizeErrorMessage(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Plain text message",
			input:    "Connection timeout",
			expected: "Connection timeout",
		},
		{
			name:     "HTML response",
			input:    `<html><head><title>Service Unavailable</title></head><body>Try again later</body></html>`,
			expected: "Service Unavailable",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SanitizeErrorMessage(tt.input)
			if result != tt.expected {
				t.Errorf("SanitizeErrorMessage(%q) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}
