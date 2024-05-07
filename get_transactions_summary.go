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

type GetTransactionsSummaryRequest struct {
	ProductType        string `json:"product_type"`
	ContractExpiryTime string `json:"contract_expiry_type"`
}

type GetTransactionsSummaryResponse struct {
	TotalVolume             int                            `json:"total_volume"`
	TotalFees               float64                        `json:"total_fees"`
	FeeTier                 FeeTier                        `json:"fee_tier"`
	MarginRate              Rate                           `json:"margin_rate"`
	GoodsAndServicesTax     Gst                            `json:"goods_and_services_tax"`
	AdvancedTradeOnlyVolume int                            `json:"advanced_trade_only_volume"`
	AdvancedTradeOnlyFees   float64                        `json:"advanced_trade_only_fees"`
	CoinbaseProVolume       int                            `json:"coinbase_pro_volume"`
	CoinbaseProFees         float64                        `json:"coinbase_pro_fees"`
	Request                 *GetTransactionsSummaryRequest `json:"request"`
}

func (c Client) GetTransactionsSummary(
	ctx context.Context,
	request *GetTransactionsSummaryRequest,
) (*GetTransactionsSummaryResponse, error) {

	path := fmt.Sprint("/brokerage/transaction_summary")

	response := &GetTransactionsSummaryResponse{Request: request}

	if err := get(ctx, c, path, emptyQueryParams, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
