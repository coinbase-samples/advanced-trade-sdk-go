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

type EditOrderRequest struct {
	OrderId string `json:"order_id"`
	Price   string `json:"price"`
	Size    string `json:"size"`
}

type EditOrderResponse struct {
	Success    bool              `json:"success"`
	EditErrors []*EditError      `json:"errors,omitempty"`
	Request    *EditOrderRequest `json:"request"`
}

func (c Client) EditOrder(
	ctx context.Context,
	request *EditOrderRequest,
) (*EditOrderResponse, error) {

	path := fmt.Sprint("/brokerage/orders/edit")

	response := &EditOrderResponse{Request: request}

	if err := post(ctx, c, path, emptyQueryParams, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
