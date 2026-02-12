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

package client

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/coinbase-samples/advanced-trade-sdk-go/utils"
)

// RateLimitError represents a rate limit exceeded error
type RateLimitError struct {
	StatusCode int
	Message    string
	URL        string
}

func (e *RateLimitError) Error() string {
	if e.URL != "" {
		return fmt.Sprintf("rate limit exceeded (HTTP %d) for %s", e.StatusCode, e.URL)
	}
	return fmt.Sprintf("rate limit exceeded (HTTP %d)", e.StatusCode)
}

// SanitizedAPIError represents a sanitized version of API errors
type SanitizedAPIError struct {
	StatusCode int
	Message    string
	URL        string
	Original   error
}

func (e *SanitizedAPIError) Error() string {
	if e.URL != "" {
		return fmt.Sprintf("%s (HTTP %d) for %s", e.Message, e.StatusCode, e.URL)
	}
	return fmt.Sprintf("%s (HTTP %d)", e.Message, e.StatusCode)
}

// Unwrap allows error unwrapping to access the original error
func (e *SanitizedAPIError) Unwrap() error {
	return e.Original
}

// Pattern to extract status code and URL from core-go error messages
var coreGoErrorPattern = regexp.MustCompile(`Unexpected response: (.+?) Expected Status Codes: \[[^\]]+\], Received Status Code: (\d+), URL: (.+)`)

// SanitizeError processes errors returned by core-go HTTP calls and sanitizes HTML responses
func SanitizeError(err error) error {
	if err == nil {
		return nil
	}

	errorText := err.Error()

	// Try to parse core-go error format
	matches := coreGoErrorPattern.FindStringSubmatch(errorText)
	if len(matches) == 4 {
		responseBody := matches[1]
		statusCodeStr := matches[2]
		url := matches[3]

		statusCode, parseErr := strconv.Atoi(statusCodeStr)
		if parseErr != nil {
			// If we can't parse status code, fall back to original error
			return err
		}

		// Handle rate limiting specifically
		if statusCode == http.StatusTooManyRequests {
			return &RateLimitError{
				StatusCode: statusCode,
				Message:    "rate limit exceeded",
				URL:        url,
			}
		}

		// Sanitize other HTML responses
		sanitizedMessage := utils.SanitizeRateLimitError(responseBody, statusCode)

		return &SanitizedAPIError{
			StatusCode: statusCode,
			Message:    sanitizedMessage,
			URL:        url,
			Original:   err,
		}
	}

	// If the error doesn't match the expected pattern, check if it contains HTML
	if utils.IsHTMLResponse(errorText) {
		sanitizedMessage := utils.SanitizeErrorMessage(errorText)
		return &SanitizedAPIError{
			StatusCode: 0, // Unknown status code
			Message:    sanitizedMessage,
			URL:        "",
			Original:   err,
		}
	}

	// Return original error if no sanitization is needed
	return err
}

// IsRateLimitError checks if the error is a rate limit error
func IsRateLimitError(err error) bool {
	_, ok := err.(*RateLimitError)
	if ok {
		return true
	}

	// Also check wrapped errors
	if sanitized, ok := err.(*SanitizedAPIError); ok {
		return sanitized.StatusCode == http.StatusTooManyRequests
	}

	// Check if error message indicates rate limiting
	if err != nil {
		errorText := strings.ToLower(err.Error())
		return strings.Contains(errorText, "rate limit") ||
			strings.Contains(errorText, "too many requests") ||
			strings.Contains(errorText, "429")
	}

	return false
}

// GetStatusCode extracts the HTTP status code from sanitized errors
func GetStatusCode(err error) int {
	if rateLimitErr, ok := err.(*RateLimitError); ok {
		return rateLimitErr.StatusCode
	}

	if sanitizedErr, ok := err.(*SanitizedAPIError); ok {
		return sanitizedErr.StatusCode
	}

	return 0
}
