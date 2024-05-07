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

package test

import (
	"context"
	adv "github.com/coinbase-samples/advanced-trade-sdk-go"
	"testing"
	"time"
)

func TestGetOrder(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := setupClient()
	if err != nil {
		t.Fatal(err)
	}

	ordersResponse, err := client.ListOrders(ctx, &adv.ListOrdersRequest{
		OrderStatus: []string{"FILLED"},
	})

	if err != nil {
		t.Fatal("failed to list orders:", err)
	}

	if ordersResponse == nil || len(ordersResponse.Orders) == 0 {
		t.Fatal("no orders found or nil response")
	}

	firstOrder := ordersResponse.Orders[0]

	testGetSpecificOrder(t, client, firstOrder.OrderId)
}

func testGetSpecificOrder(t *testing.T, client *adv.Client, orderId string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := client.GetOrder(ctx, &adv.GetOrderRequest{
		OrderId: orderId,
	})

	if err != nil {
		t.Fatal("Failed to get order details:", err)
	}

	if response == nil {
		t.Fatal("Expected non-nil response for GetOrder")
	}

	if response.Order == nil {
		t.Fatal("Expected non-nil Order in the response")
	}

	if response.Order.OrderId != orderId {
		t.Fatal("Order ID mismatch in the response")
	}
}
