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

package portfolios

import (
	"context"

	"github.com/coinbase-samples/advanced-trade-sdk-go/client"
	"github.com/coinbase-samples/advanced-trade-sdk-go/model"
	"github.com/coinbase-samples/core-go"
)

type ListPortfoliosRequest struct {
	PortfolioType string `json:"portfolio_type,omitempty"`
}

type ListPortfoliosResponse struct {
	Portfolios []*model.Portfolio     `json:"portfolios"`
	Request    *ListPortfoliosRequest `json:"request"`
}

func (s portfoliosServiceImpl) ListPortfolios(
	ctx context.Context,
	request *ListPortfoliosRequest,
) (*ListPortfoliosResponse, error) {

	path := "/brokerage/portfolios"

	var queryParams string
	if len(request.PortfolioType) > 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "portfolio_type", request.PortfolioType)
	}

	response := &ListPortfoliosResponse{Request: request}

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
