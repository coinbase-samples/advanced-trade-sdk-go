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

package futures

import (
	"context"

	"github.com/coinbase-samples/advanced-trade-sdk-go/client"
)

type FuturesService interface {
	GetFuturesBalanceSummary(ctx context.Context, request *GetFuturesBalanceSummaryRequest) (*GetFuturesBalanceSummaryResponse, error)
	GetFuturesPosition(ctx context.Context, request *GetFuturesPositionRequest) (*GetFuturesPositionResponse, error)
	ListFuturesPositions(ctx context.Context, request *ListFuturesPositionsRequest) (*ListFuturesPositionsResponse, error)
	ListFuturesSweeps(ctx context.Context, request *ListFuturesSweepsRequest) (*ListFuturesSweepsResponse, error)
	ScheduleFuturesSweep(ctx context.Context, request *ScheduleFuturesSweepRequest) (*ScheduleFuturesSweepResponse, error)
	CancelPendingFuturesSweeps(ctx context.Context, request *CancelPendingFuturesSweepsRequest) (*CancelPendingFuturesSweepsResponse, error)
	GetCurrentMarginWindow(ctx context.Context, request *GetCurrentMarginWindowRequest) (*GetCurrentMarginWindowResponse, error)
	GetIntradayMarginSetting(ctx context.Context, request *GetIntradayMarginSettingRequest) (*GetIntradayMarginSettingResponse, error)
	SetIntradayMarginSetting(ctx context.Context, request *SetIntradayMarginSettingRequest) (*SetIntradayMarginSettingResponse, error)
}

func NewFuturesService(c client.RestClient) FuturesService {
	return &futuresServiceImpl{client: c}
}

type futuresServiceImpl struct {
	client client.RestClient
}
