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

type GetProductCandlesRequest struct {
	ProductId   string `json:"product_id"`
	Start       string `json:"start"`
	End         string `json:"end"`
	Limit       string `json:"limit,omitempty"`
	Granularity string `json:"granularity"`
}

type GetProductCandlesResponse struct {
	Candles *[]model.Candle           `json:"candles"`
	Request *GetProductCandlesRequest `json:"request"`
}

func (s productsServiceImpl) GetProductCandles(
	ctx context.Context,
	request *GetProductCandlesRequest,
) (*GetProductCandlesResponse, error) {

	path := fmt.Sprintf("/brokerage/products/%s/candles", request.ProductId)

	response := &GetProductCandlesResponse{Request: request}

	var queryParams string
	queryParams = core.AppendHttpQueryParam(queryParams, "product_id", request.ProductId)
	queryParams = core.AppendHttpQueryParam(queryParams, "granularity", request.Granularity)
	queryParams = core.AppendHttpQueryParam(queryParams, "start", request.Start)
	queryParams = core.AppendHttpQueryParam(queryParams, "end", request.End)
	queryParams = core.AppendHttpQueryParam(queryParams, "limit", request.Limit)

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
