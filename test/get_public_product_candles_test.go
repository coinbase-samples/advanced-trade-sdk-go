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
	"net/http"
	"testing"
	"time"
)

func TestGetPublicProductCandles(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := adv.NewClient(&adv.Credentials{}, http.Client{})

	candlesResponse, err := client.GetPublicProductCandles(ctx, &adv.GetPublicProductCandlesRequest{
		ProductId:   "BTC-USD",
		Granularity: "ONE_MINUTE",
	})

	if err != nil {
		t.Fatal("failed to get public product candles:", err)
	}

	if candlesResponse == nil || len(*candlesResponse.Candles) == 0 {
		t.Fatal("no candles found or nil response")
	}
}
