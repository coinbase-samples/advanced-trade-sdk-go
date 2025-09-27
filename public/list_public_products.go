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
	"github.com/coinbase-samples/advanced-trade-sdk-go/utils"
	"github.com/coinbase-samples/core-go"
)

type ListPublicProductsRequest struct {
	ProductType            string                  `json:"product_type"`
	ProductIds             []string                `json:"product_ids"`
	ContractExpiryType     string                  `json:"contract_expiry_type"`
	ExpiringContractStatus string                  `json:"expiring_contract_status"`
	Pagination             *model.PaginationParams `json:"pagination_params"`
}

type ListPublicProductsResponse struct {
	Products []*model.Product           `json:"products"`
	Request  *ListPublicProductsRequest `json:"request"`
}

func (s publicServiceImpl) ListPublicProducts(
	ctx context.Context,
	request *ListPublicProductsRequest,
) (*ListPublicProductsResponse, error) {

	path := fmt.Sprintf("/brokerage/market/products")

	var queryParams string
	for _, p := range request.ProductIds {
		queryParams = core.AppendHttpQueryParam(queryParams, "product_ids", p)
	}

	if len(request.ProductType) > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "product_type", request.ProductType)
	}

	if len(request.ContractExpiryType) > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "contract_expiry_type", request.ContractExpiryType)
	}

	if len(request.ExpiringContractStatus) > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "expiring_contract_status", request.ExpiringContractStatus)
	}

	queryParams = utils.AppendPaginationParams(queryParams, request.Pagination)

	response := &ListPublicProductsResponse{Request: request}

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
