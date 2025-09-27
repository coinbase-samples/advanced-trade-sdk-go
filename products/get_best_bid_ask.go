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

package products

import (
	"context"
	"fmt"

	"github.com/coinbase-samples/advanced-trade-sdk-go/client"
	"github.com/coinbase-samples/advanced-trade-sdk-go/model"
	"github.com/coinbase-samples/core-go"
)

type GetBestBidAskRequest struct {
	ProductIds []string `json:"product_ids,omitempty"`
}

type GetBestBidAskResponse struct {
	PriceBooks *[]model.PriceBook    `json:"pricebooks"`
	Request    *GetBestBidAskRequest `json:"request"`
}

func (s productsServiceImpl) GetBestBidAsk(
	ctx context.Context,
	request *GetBestBidAskRequest,
) (*GetBestBidAskResponse, error) {

	path := fmt.Sprintf("/brokerage/best_bid_ask")

	var queryParams string
	for _, p := range request.ProductIds {
		queryParams = core.AppendHttpQueryParam(queryParams, "product_ids", p)
	}

	response := &GetBestBidAskResponse{Request: request}

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
