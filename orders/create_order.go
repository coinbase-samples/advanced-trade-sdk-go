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

package orders

import (
	"context"
	"fmt"

	"github.com/coinbase-samples/advanced-trade-sdk-go/client"
	"github.com/coinbase-samples/advanced-trade-sdk-go/model"
	"github.com/coinbase-samples/core-go"
)

type CreateOrderRequest struct {
	ProductId          string                   `json:"product_id"`
	Side               string                   `json:"side"`
	ClientOrderId      string                   `json:"client_order_id"`
	OrderConfiguration model.OrderConfiguration `json:"order_configuration"`
	CommissionRate     *model.Rate              `json:"commission_rate,omitempty"`
	IsMax              *bool                    `json:"is_max,omitempty"`
	TradableBalance    string                   `json:"tradable_balance,omitempty"`
	SkipFcmRiskCheck   *bool                    `json:"skip_fcm_risk_check,omitempty"`
	Leverage           string                   `json:"leverage,omitempty"`
	MarginType         string                   `json:"margin_type,omitempty"`
	RetailPortfolioId  string                   `json:"retail_portfolio_id,omitempty"`
}

type CreateOrderResponse struct {
	Success            bool                     `json:"success"`
	FailureReason      string                   `json:"failure_reason,omitempty"`
	OrderId            string                   `json:"order_id,omitempty"`
	SuccessResponse    *model.SuccessResponse   `json:"success_response,omitempty"`
	ErrorResponse      *model.ErrorResponse     `json:"error_response,omitempty"`
	OrderConfiguration model.OrderConfiguration `json:"order_configuration"`
	Request            *CreateOrderRequest      `json:"request"`
}

func (s ordersServiceImpl) CreateOrder(
	ctx context.Context,
	request *CreateOrderRequest,
) (*CreateOrderResponse, error) {

	path := fmt.Sprint("/brokerage/orders")

	response := &CreateOrderResponse{Request: request}

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
