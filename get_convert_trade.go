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

type GetConvertTradeRequest struct {
	TradeId     string `json:"trade_id"`
	FromAccount string `json:"from_account"`
	ToAccount   string `json:"to_account"`
}

type GetConvertTradeResponse struct {
	Convert *Convert                `json:"trade"`
	Request *GetConvertTradeRequest `json:"request"`
}

func (c Client) GetConvertTrade(
	ctx context.Context,
	request *GetConvertTradeRequest,
) (*GetConvertTradeResponse, error) {

	path := fmt.Sprintf("/brokerage/convert/trade/%s", request.TradeId)

	response := &GetConvertTradeResponse{Request: request}

	if err := get(ctx, c, path, emptyQueryParams, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
