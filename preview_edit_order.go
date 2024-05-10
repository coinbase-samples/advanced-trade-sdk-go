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

package adv

import (
	"context"
	"fmt"
)

type PreviewEditOrderRequest struct {
	OrderId string `json:"order_id"`
	Price   string `json:"price"`
	Size    string `json:"size"`
}

type PreviewEditOrderResponse struct {
	EditErrors      []*EditError             `json:"errors,omitempty"`
	Slippage        string                   `json:"slippage"`
	OrderTotal      string                   `json:"order_total"`
	CommissionTotal string                   `json:"commission_total"`
	QuoteSize       string                   `json:"quote_size"`
	BaseSize        string                   `json:"base_size"`
	BestBid         string                   `json:"best_bid"`
	BestAsk         string                   `json:"best_ask"`
	Leverage        string                   `json:"leverage"`
	LongLeverage    string                   `json:"long_leverage"`
	ShortLeverage   string                   `json:"short_leverage"`
	Request         *PreviewEditOrderRequest `json:"request"`
}

func (c Client) PreviewEditOrder(
	ctx context.Context,
	request *PreviewEditOrderRequest,
) (*PreviewEditOrderResponse, error) {
	path := fmt.Sprint("/brokerage/orders/edit_preview")

	response := &PreviewEditOrderResponse{Request: request}

	if err := post(ctx, c, path, emptyQueryParams, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
