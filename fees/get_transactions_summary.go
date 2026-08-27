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

package fees

import (
	"context"
	"fmt"

	"github.com/coinbase-samples/advanced-trade-sdk-go/client"
	"github.com/coinbase-samples/advanced-trade-sdk-go/model"
	"github.com/coinbase-samples/core-go"
)

type GetTransactionsSummaryRequest struct {
	ProductType        string `json:"product_type"`
	ContractExpiryTime string `json:"contract_expiry_type"`
}

type GetTransactionsSummaryResponse struct {
	TotalVolume             float64                        `json:"total_volume"`
	TotalFees               float64                        `json:"total_fees"`
	FeeTier                 model.FeeTier                  `json:"fee_tier"`
	MarginRate              model.Rate                     `json:"margin_rate"`
	GoodsAndServicesTax     model.Gst                      `json:"goods_and_services_tax"`
	AdvancedTradeOnlyVolume float64                        `json:"advanced_trade_only_volume"`
	CoinbaseProVolume       float64                        `json:"coinbase_pro_volume"`
	CoinbaseProFees         float64                        `json:"coinbase_pro_fees"`
	Request                 *GetTransactionsSummaryRequest `json:"request"`
}

func (s feesServiceImpl) GetTransactionsSummary(
	ctx context.Context,
	request *GetTransactionsSummaryRequest,
) (*GetTransactionsSummaryResponse, error) {

	path := fmt.Sprint("/brokerage/transaction_summary")

	response := &GetTransactionsSummaryResponse{Request: request}

	if err := core.HttpGet(
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
