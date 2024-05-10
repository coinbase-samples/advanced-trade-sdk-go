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

type GetBestBidAskRequest struct {
	ProductIds []string `json:"product_ids,omitempty"`
}

type GetBestBidAskResponse struct {
	PriceBooks *[]PriceBook          `json:"pricebooks"`
	Request    *GetBestBidAskRequest `json:"request"`
}

func (c Client) GetBestBidAsk(
	ctx context.Context,
	request *GetBestBidAskRequest,
) (*GetBestBidAskResponse, error) {

	path := fmt.Sprintf("/brokerage/best_bid_ask")

	var queryParams string

	for _, p := range request.ProductIds {
		queryParams = appendQueryParam(queryParams, "product_ids", p)
	}

	response := &GetBestBidAskResponse{Request: request}

	if err := get(ctx, c, path, queryParams, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
