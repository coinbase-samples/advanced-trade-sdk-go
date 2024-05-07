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

package test

import (
	"context"
	adv "github.com/coinbase-samples/advanced-trade-sdk-go"
	"testing"
)

func TestListPortfolios(t *testing.T) {
	client, err := setupClient()
	if err != nil {
		t.Fatalf("Error setting up client: %v", err)
	}

	ctx := context.Background()
	response, err := client.ListPortfolios(ctx, &adv.ListPortfoliosRequest{})

	if err != nil {
		t.Errorf("Failed to list portfolios: %v", err)
	}

	if response == nil {
		t.Error("Expected non-nil response, got nil")
	}

	if len(response.Portfolios) == 0 {
		t.Error("Expected at least one portfolio in the response")
	}
}
