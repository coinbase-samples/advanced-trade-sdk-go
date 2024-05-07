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

type CreateOrderPreviewRequest struct {
	ProductId          string             `json:"product_id"`
	Side               string             `json:"side"`
	OrderConfiguration OrderConfiguration `json:"order_configuration"`
	CommissionRate     *Rate              `json:"commission_rate,omitempty"`
	IsMax              *bool              `json:"is_max,omitempty"`
	TradableBalance    string             `json:"tradable_balance,omitempty"`
	SkipFcmRiskCheck   *bool              `json:"skip_fcm_risk_check,omitempty"`
	Leverage           string             `json:"leverage,omitempty"`
	MarginType         string             `json:"margin_type,omitempty"`
	RetailPortfolioId  string             `json:"retail_portfolio_id,omitempty"`
}

type CreateOrderPreviewResponse struct {
	OrderTotal         string                     `json:"order_total"`
	CommissionTotal    string                     `json:"commission_total"`
	Errs               []string                   `json:"errs"`
	Warning            []string                   `json:"warning"`
	QuoteSize          string                     `json:"quote_size"`
	BaseSize           string                     `json:"base_size"`
	BestBid            string                     `json:"best_bid"`
	BestAsk            string                     `json:"best_ask"`
	IsMax              bool                       `json:"is_max"`
	OrderMarginTotal   string                     `json:"order_margin_total"`
	Leverage           string                     `json:"leverage"`
	LongLeverage       string                     `json:"long_leverage"`
	ShortLeverage      string                     `json:"short_leverage"`
	Slippage           string                     `json:"slippage"`
	AverageFilledPrice string                     `json:"average_filled_price"`
	Request            *CreateOrderPreviewRequest `json:"request"`
}

func (c Client) CreateOrderPreview(
	ctx context.Context,
	request *CreateOrderPreviewRequest,
) (*CreateOrderPreviewResponse, error) {

	path := fmt.Sprint("/brokerage/orders/preview")

	response := &CreateOrderPreviewResponse{Request: request}

	if err := post(ctx, c, path, emptyQueryParams, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
