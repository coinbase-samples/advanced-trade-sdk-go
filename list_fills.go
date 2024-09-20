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
)

type ListFillsRequest struct {
	OrderId                string `json:"order_id,omitempty"`
	ProductId              string `json:"product_id,omitempty"`
	StartSequenceTimestamp string `json:"start_sequence_timestamp,omitempty"`
	EndSequenceTimestamp   string `json:"end_sequence_timestamp,omitempty"`
	Limit                  string `json:"limit,omitempty"`
	Cursor                 string `json:"cursor,omitempty"`
}

type ListFillsResponse struct {
	Fills   []*Fill           `json:"fills"`
	Cursor  string            `json:"cursor"`
	Request *ListFillsRequest `json:"request"`
}

func (c Client) ListFills(
	ctx context.Context,
	request *ListFillsRequest,
) (*ListFillsResponse, error) {

	path := "/brokerage/orders/historical/fills"

	var queryParams string

	if request.OrderId != "" {
		queryParams = appendQueryParam(queryParams, "order_id", request.OrderId)
	}

	if request.ProductId != "" {
		queryParams = appendQueryParam(queryParams, "product_id", request.ProductId)
	}
	if request.StartSequenceTimestamp != "" {
		queryParams = appendQueryParam(queryParams, "start_sequence_timestamp", request.StartSequenceTimestamp)
	}
	if request.EndSequenceTimestamp != "" {
		queryParams = appendQueryParam(queryParams, "end_sequence_timestamp", request.EndSequenceTimestamp)
	}
	if request.Limit != "" {
		queryParams = appendQueryParam(queryParams, "limit", request.Limit)
	}

	response := &ListFillsResponse{Request: request}

	if err := getPrivate(ctx, c, path, queryParams, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
