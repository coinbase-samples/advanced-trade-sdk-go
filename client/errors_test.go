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
	"errors"
	"testing"
)

func TestRateLimitError(t *testing.T) {
	tests := []struct {
		name          string
		statusCode    int
		message       string
		url           string
		expectedError string
	}{
		{
			name:          "Rate limit error with URL",
			statusCode:    429,
			message:       "rate limit exceeded",
			url:           "https://api.coinbase.com/api/v3/brokerage/orders",
			expectedError: "rate limit exceeded (HTTP 429) for https://api.coinbase.com/api/v3/brokerage/orders",
		},
		{
			name:          "Rate limit error without URL",
			statusCode:    429,
			message:       "rate limit exceeded",
			url:           "",
			expectedError: "rate limit exceeded (HTTP 429)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &RateLimitError{
				StatusCode: tt.statusCode,
				Message:    tt.message,
				URL:        tt.url,
			}

			result := err.Error()
			if result != tt.expectedError {
				t.Errorf("RateLimitError.Error() = %q, expected %q", result, tt.expectedError)
			}
		})
	}
}

func TestSanitizedAPIError(t *testing.T) {
	originalErr := errors.New("original error")

	tests := []struct {
		name          string
		statusCode    int
		message       string
		url           string
		expectedError string
	}{
		{
			name:          "Sanitized error with URL",
			statusCode:    400,
			message:       "bad request",
			url:           "https://api.coinbase.com/api/v3/brokerage/orders",
			expectedError: "bad request (HTTP 400) for https://api.coinbase.com/api/v3/brokerage/orders",
		},
		{
			name:          "Sanitized error without URL",
			statusCode:    500,
			message:       "internal server error",
			url:           "",
			expectedError: "internal server error (HTTP 500)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &SanitizedAPIError{
				StatusCode: tt.statusCode,
				Message:    tt.message,
				URL:        tt.url,
				Original:   originalErr,
			}

			result := err.Error()
			if result != tt.expectedError {
				t.Errorf("SanitizedAPIError.Error() = %q, expected %q", result, tt.expectedError)
			}

			// Test unwrapping
			if err.Unwrap() != originalErr {
				t.Errorf("SanitizedAPIError.Unwrap() did not return original error")
			}
		})
	}
}

func TestSanitizeError(t *testing.T) {
	tests := []struct {
		name         string
		inputError   error
		expectedType string
		expectedMsg  string
		isRateLimit  bool
	}{
		{
			name:         "Nil error",
			inputError:   nil,
			expectedType: "nil",
		},
		{
			name:         "Core-go rate limit error (429)",
			inputError:   errors.New("Unexpected response: <html><head><title>Coinbase</title></head><body>Rate limit exceeded</body></html> Expected Status Codes: [200], Received Status Code: 429, URL: https://api.coinbase.com/api/v3/brokerage/orders"),
			expectedType: "*client.RateLimitError",
			expectedMsg:  "rate limit exceeded (HTTP 429) for https://api.coinbase.com/api/v3/brokerage/orders",
			isRateLimit:  true,
		},
		{
			name:         "Core-go HTML error (400)",
			inputError:   errors.New("Unexpected response: <html><head><title>Bad Request</title></head><body>Invalid request</body></html> Expected Status Codes: [200], Received Status Code: 400, URL: https://api.coinbase.com/api/v3/brokerage/orders"),
			expectedType: "*client.SanitizedAPIError",
			expectedMsg:  "Bad Request (HTTP 400) for https://api.coinbase.com/api/v3/brokerage/orders",
			isRateLimit:  false,
		},
		{
			name:         "Core-go generic HTML error (500)",
			inputError:   errors.New("Unexpected response: <html><head><title>Coinbase</title></head><body>Error</body></html> Expected Status Codes: [200], Received Status Code: 500, URL: https://api.coinbase.com/api/v3/brokerage/orders"),
			expectedType: "*client.SanitizedAPIError",
			expectedMsg:  "internal server error (HTTP 500) for https://api.coinbase.com/api/v3/brokerage/orders",
			isRateLimit:  false,
		},
		{
			name:         "HTML content not matching core-go pattern",
			inputError:   errors.New("<html><head><title>Error Page</title></head><body>Something went wrong</body></html>"),
			expectedType: "*client.SanitizedAPIError",
			expectedMsg:  "Error Page (HTTP 0)",
			isRateLimit:  false,
		},
		{
			name:         "Regular non-HTML error",
			inputError:   errors.New("connection timeout"),
			expectedType: "error",
			expectedMsg:  "connection timeout",
			isRateLimit:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SanitizeError(tt.inputError)

			if result == nil {
				if tt.expectedType != "nil" {
					t.Errorf("SanitizeError() = nil, expected %s", tt.expectedType)
				}
				return
			}

			// Check error message
			if result.Error() != tt.expectedMsg {
				t.Errorf("SanitizeError().Error() = %q, expected %q", result.Error(), tt.expectedMsg)
			}

			// Check if it's identified as rate limit error
			if IsRateLimitError(result) != tt.isRateLimit {
				t.Errorf("IsRateLimitError() = %v, expected %v", IsRateLimitError(result), tt.isRateLimit)
			}
		})
	}
}

func TestIsRateLimitError(t *testing.T) {
	tests := []struct {
		name     string
		error    error
		expected bool
	}{
		{
			name:     "Nil error",
			error:    nil,
			expected: false,
		},
		{
			name:     "RateLimitError",
			error:    &RateLimitError{StatusCode: 429},
			expected: true,
		},
		{
			name:     "SanitizedAPIError with 429",
			error:    &SanitizedAPIError{StatusCode: 429},
			expected: true,
		},
		{
			name:     "SanitizedAPIError with other status",
			error:    &SanitizedAPIError{StatusCode: 400},
			expected: false,
		},
		{
			name:     "Error with rate limit in message",
			error:    errors.New("rate limit exceeded"),
			expected: true,
		},
		{
			name:     "Error with too many requests in message",
			error:    errors.New("too many requests"),
			expected: true,
		},
		{
			name:     "Error with 429 in message",
			error:    errors.New("HTTP 429 error"),
			expected: true,
		},
		{
			name:     "Regular error",
			error:    errors.New("connection timeout"),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsRateLimitError(tt.error)
			if result != tt.expected {
				t.Errorf("IsRateLimitError() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestGetStatusCode(t *testing.T) {
	tests := []struct {
		name     string
		error    error
		expected int
	}{
		{
			name:     "RateLimitError",
			error:    &RateLimitError{StatusCode: 429},
			expected: 429,
		},
		{
			name:     "SanitizedAPIError",
			error:    &SanitizedAPIError{StatusCode: 400},
			expected: 400,
		},
		{
			name:     "Regular error",
			error:    errors.New("some error"),
			expected: 0,
		},
		{
			name:     "Nil error",
			error:    nil,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetStatusCode(tt.error)
			if result != tt.expected {
				t.Errorf("GetStatusCode() = %d, expected %d", result, tt.expected)
			}
		})
	}
}
