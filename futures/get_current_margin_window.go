/**
 * Copyright 2025-present Coinbase Global, Inc.
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

package futures

import (
	"context"

	"github.com/coinbase-samples/advanced-trade-sdk-go/client"
	"github.com/coinbase-samples/advanced-trade-sdk-go/model"
	"github.com/coinbase-samples/core-go"
)

type GetCurrentMarginWindowRequest struct {
	MarginProfileType string `json:"margin_profile_type"`
}

type GetCurrentMarginWindowResponse struct {
	MarginSettings *model.MarginSettings          `json:"margin_settings"`
	Request        *GetCurrentMarginWindowRequest `json:"request"`
}

func (s futuresServiceImpl) GetCurrentMarginWindow(
	ctx context.Context,
	request *GetCurrentMarginWindowRequest,
) (*GetCurrentMarginWindowResponse, error) {

	path := "/brokerage/cfm/intraday/current_margin_window"

	response := &GetCurrentMarginWindowResponse{Request: request}

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
