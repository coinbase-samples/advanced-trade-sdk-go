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
	"regexp"
	"strings"
)

// HTML detection patterns
var (
	htmlDocumentPattern = regexp.MustCompile(`(?i)^\s*<\s*(!DOCTYPE\s+html|html)`)
	htmlTagPattern      = regexp.MustCompile(`<[^>]*>`)
	titlePattern        = regexp.MustCompile(`(?i)<title[^>]*>(.*?)</title>`)
)

// IsHTMLResponse checks if the given text appears to be an HTML response
func IsHTMLResponse(text string) bool {
	if text == "" {
		return false
	}

	// Check for common HTML indicators
	trimmed := strings.TrimSpace(text)

	// Check for HTML document declaration or opening tag
	if htmlDocumentPattern.MatchString(trimmed) {
		return true
	}

	// Check for multiple distinct HTML tags which likely indicates HTML content
	// We need at least 3 tags to consider it HTML (to avoid single element cases)
	matches := htmlTagPattern.FindAllString(trimmed, -1)
	if len(matches) >= 3 {
		return true
	}

	// Special case: if we see html, head, body, or title tags, it's definitely HTML
	if strings.Contains(strings.ToLower(trimmed), "<head>") ||
		strings.Contains(strings.ToLower(trimmed), "<body>") ||
		strings.Contains(strings.ToLower(trimmed), "<title>") {
		return true
	}

	return false
}

// ExtractErrorMessage extracts a meaningful error message from HTML content
func ExtractErrorMessage(htmlContent string) string {
	if !IsHTMLResponse(htmlContent) {
		return htmlContent
	}

	// Try to extract title from HTML
	if matches := titlePattern.FindStringSubmatch(htmlContent); len(matches) > 1 {
		title := strings.TrimSpace(matches[1])
		if title != "" && title != "Coinbase" {
			return title
		}
	}

	// If we can't extract meaningful info from HTML, return a generic message
	return "request failed"
}

// SanitizeRateLimitError creates a clean error message for rate limiting scenarios
func SanitizeRateLimitError(originalError string, statusCode int) string {
	if statusCode == 429 {
		// For HTTP 429, always return a concise rate limit message
		return "rate limit exceeded"
	}

	if IsHTMLResponse(originalError) {
		extracted := ExtractErrorMessage(originalError)
		if extracted == "request failed" {
			// If we couldn't extract anything meaningful, provide context based on status code
			switch statusCode {
			case 400:
				return "bad request"
			case 401:
				return "unauthorized"
			case 403:
				return "forbidden"
			case 404:
				return "not found"
			case 500:
				return "internal server error"
			case 502:
				return "bad gateway"
			case 503:
				return "service unavailable"
			default:
				return "request failed"
			}
		}
		return extracted
	}

	return originalError
}

// SanitizeErrorMessage provides a general-purpose error message sanitizer
func SanitizeErrorMessage(errorMessage string) string {
	if IsHTMLResponse(errorMessage) {
		return ExtractErrorMessage(errorMessage)
	}
	return errorMessage
}
