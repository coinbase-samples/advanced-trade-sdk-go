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

type GetPublicProductBookRequest struct {
	ProductId string `json:"product_id"`
	Limit     string `json:"limit,omitempty"`
}

type GetPublicProductBookResponse struct {
	PriceBook *PriceBook                   `json:"pricebook"`
	Request   *GetPublicProductBookRequest `json:"request"`
}

func (c Client) GetPublicProductBook(
	ctx context.Context,
	request *GetPublicProductBookRequest,
) (*GetPublicProductBookResponse, error) {

	path := fmt.Sprintf("/brokerage/market/product_book?product_id=%s", request.ProductId)

	response := &GetPublicProductBookResponse{Request: request}

	if err := getPublic(ctx, c, path, emptyQueryParams, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
