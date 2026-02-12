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

package converts

import (
	"context"
	"fmt"

	"github.com/coinbase-samples/advanced-trade-sdk-go/client"
	"github.com/coinbase-samples/advanced-trade-sdk-go/model"
	"github.com/coinbase-samples/core-go"
)

type CreateConvertQuoteRequest struct {
	FromAccount            string                        `json:"from_account"`
	ToAccount              string                        `json:"to_account"`
	Amount                 string                        `json:"amount"`
	TradeIncentiveMetadata *model.TradeIncentiveMetadata `json:"trade_incentive_metadata,omitempty"`
}

type CreateConvertQuoteResponse struct {
	Convert *model.Convert             `json:"trade"`
	Request *CreateConvertQuoteRequest `json:"request"`
}

func (s convertsServiceImpl) CreateConvertQuote(
	ctx context.Context,
	request *CreateConvertQuoteRequest,
) (*CreateConvertQuoteResponse, error) {

	path := fmt.Sprint("/brokerage/convert/quote")

	response := &CreateConvertQuoteResponse{Request: request}

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
