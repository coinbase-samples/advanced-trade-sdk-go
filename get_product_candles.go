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

type GetProductCandlesRequest struct {
	ProductId   string `json:"product_id"`
	Start       string `json:"start"`
	End         string `json:"end"`
	Granularity string `json:"granularity"`
}

type GetProductCandlesResponse struct {
	Candles *[]Candle                 `json:"candles"`
	Request *GetProductCandlesRequest `json:"request"`
}

func (c Client) GetProductCandles(
	ctx context.Context,
	request *GetProductCandlesRequest,
) (*GetProductCandlesResponse, error) {

	path := fmt.Sprintf("/brokerage/products/%s/candles", request.ProductId)

	response := &GetProductCandlesResponse{Request: request}

	var queryParams string

	queryParams = appendQueryParam(queryParams, "product_id", request.ProductId)

	queryParams = appendQueryParam(queryParams, "granularity", request.Granularity)

	queryParams = appendQueryParam(queryParams, "start", request.Start)

	queryParams = appendQueryParam(queryParams, "end", request.End)

	if err := get(ctx, c, path, queryParams, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
