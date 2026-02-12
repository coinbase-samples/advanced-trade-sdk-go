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

package public

import (
	"context"
	"fmt"

	"github.com/coinbase-samples/advanced-trade-sdk-go/client"
	"github.com/coinbase-samples/advanced-trade-sdk-go/model"
	"github.com/coinbase-samples/core-go"
)

type GetPublicMarketTradesRequest struct {
	ProductId string `json:"product_id"`
	Limit     string `json:"limit"`
	Start     string `json:"start,omitempty"`
	End       string `json:"end,omitempty"`
}

type GetPublicMarketTradesResponse struct {
	Trades  []*model.Trade                `json:"trades"`
	BestBid string                        `json:"best_bid"`
	BestAsk string                        `json:"best_ask"`
	Request *GetPublicMarketTradesRequest `json:"request"`
}

func (s publicServiceImpl) GetPublicMarketTrades(
	ctx context.Context,
	request *GetPublicMarketTradesRequest,
) (*GetPublicMarketTradesResponse, error) {

	path := fmt.Sprintf("/brokerage/market/products/%s/ticker", request.ProductId)

	response := &GetPublicMarketTradesResponse{Request: request}

	var queryParams string

	queryParams = core.AppendHttpQueryParam(queryParams, "limit", request.Limit)

	if len(request.Start) > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "start", request.Start)
	}

	if len(request.End) > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "end", request.End)
	}

	if err := client.HttpGet(
		ctx,
		s.client,
		path,
		queryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		response,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return response, nil
}
