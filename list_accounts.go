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

import "context"

type ListAccountsRequest struct {
	Pagination        *PaginationParams `json:"pagination_params"`
	RetailPortfolioId string            `json:"retail_portfolio_id"`
}

type ListAccountsResponse struct {
	Accounts   []*Account           `json:"accounts"`
	Request    *ListAccountsRequest `json:"request"`
	Pagination *Pagination
}

func (c Client) ListAccounts(
	ctx context.Context,
	request *ListAccountsRequest,
) (*ListAccountsResponse, error) {

	path := "/brokerage/accounts"

	var queryParams string

	queryParams = appendPaginationParams(queryParams, request.Pagination)

	response := &ListAccountsResponse{Request: request}

	if err := getPrivate(ctx, c, path, queryParams, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
