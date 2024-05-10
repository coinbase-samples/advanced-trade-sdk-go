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

type GetOrderRequest struct {
	OrderId            string `json:"order_id"`
	ClientOrderId      string `json:"client_order_id,omitempty"`
	UserNativeCurrency string `json:"user_native_currency,omitempty"`
}

type GetOrderResponse struct {
	Order   *Order           `json:"order"`
	Request *GetOrderRequest `json:"request"`
}

func (c Client) GetOrder(
	ctx context.Context,
	request *GetOrderRequest,
) (*GetOrderResponse, error) {

	path := fmt.Sprintf("/brokerage/orders/historical/%s", request.OrderId)

	response := &GetOrderResponse{Request: request}

	if err := get(ctx, c, path, emptyQueryParams, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
