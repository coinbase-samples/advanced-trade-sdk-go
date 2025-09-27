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

package perpetuals

import (
	"context"

	"github.com/coinbase-samples/advanced-trade-sdk-go/client"
	"github.com/coinbase-samples/advanced-trade-sdk-go/model"
	"github.com/coinbase-samples/core-go"
)

type AllocatePortfolioRequest struct {
	PortfolioUuid string `json:"portfolio_uuid"`
	Symbol        string `json:"string"`
	Amount        string `json:"amount"`
	Currency      string `json:"currency"`
}

type AllocatePortfolioResponse struct {
	Description string                    `json:"description"`
	Schema      *model.Schema             `json:"schema"`
	Request     *AllocatePortfolioRequest `json:"request"`
}

func (s productsServiceImpl) AllocatePortfolio(
	ctx context.Context,
	request *AllocatePortfolioRequest,
) (*AllocatePortfolioResponse, error) {

	path := "/brokerage/intx/allocate"

	response := &AllocatePortfolioResponse{Request: request}

	if err := client.HttpPost(
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
