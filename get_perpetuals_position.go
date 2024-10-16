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

type GetPerpetualsPositionRequest struct {
	PortfolioUuid string `json:"portfolio_uuid"`
	Symbol        string `json:"symbol"`
}

type GetPerpetualsPositionResponse struct {
	Position *IntxPosition                 `json:"position"`
	Request  *GetPerpetualsPositionRequest `json:"request"`
}

func (c Client) GetPerpetualsPosition(
	ctx context.Context,
	request *GetPerpetualsPositionRequest,
) (*GetPerpetualsPositionResponse, error) {

	path := fmt.Sprintf("/brokerage/intx/positions/%s/%s", request.PortfolioUuid, request.Symbol)

	response := &GetPerpetualsPositionResponse{Request: request}

	if err := getPrivate(ctx, c, path, emptyQueryParams, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
