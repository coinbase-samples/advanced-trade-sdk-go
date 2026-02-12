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

type GetProductBookRequest struct {
	ProductId string `json:"product_id"`
	Limit     string `json:"limit,omitempty"`
}

type GetProductBookResponse struct {
	PriceBook *model.PriceBook       `json:"pricebook"`
	Request   *GetProductBookRequest `json:"request"`
}

func (s productsServiceImpl) GetProductBook(
	ctx context.Context,
	request *GetProductBookRequest,
) (*GetProductBookResponse, error) {

	path := fmt.Sprintf("/brokerage/product_book?product_id=%s", request.ProductId)

	response := &GetProductBookResponse{Request: request}

	if err := client.HttpGet(
		ctx,
		s.client,
		path,
		core.EmptyQueryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		response,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return response, nil
}
