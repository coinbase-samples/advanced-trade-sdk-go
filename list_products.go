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

type ListProductsRequest struct {
	ProductType            string            `json:"product_type,omitempty"`
	ProductIds             []string          `json:"product_ids,omitempty"`
	ContractExpiryType     string            `json:"contract_expiry_type,omitempty"`
	ExpiringContractStatus string            `json:"expiring_contract_status,omitempty"`
	Pagination             *PaginationParams `json:"pagination_params,omitempty"`
}

type ListProductsResponse struct {
	Products []*Product           `json:"products"`
	Request  *ListProductsRequest `json:"request"`
}

func (c Client) ListProducts(
	ctx context.Context,
	request *ListProductsRequest,
) (*ListProductsResponse, error) {

	path := fmt.Sprintf("/brokerage/products")

	var queryParams string

	for _, p := range request.ProductIds {
		queryParams = appendQueryParam(queryParams, "product_ids", p)
	}

	if len(request.ProductType) > 0 {
		queryParams = appendQueryParam(queryParams, "product_type", request.ProductType)
	}

	if len(request.ContractExpiryType) > 0 {
		queryParams = appendQueryParam(queryParams, "contract_expiry_type", request.ContractExpiryType)
	}

	if len(request.ExpiringContractStatus) > 0 {
		queryParams = appendQueryParam(queryParams, "expiring_contract_status", request.ExpiringContractStatus)
	}

	queryParams = appendPaginationParams(queryParams, request.Pagination)

	response := &ListProductsResponse{Request: request}

	if err := get(ctx, c, path, queryParams, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
