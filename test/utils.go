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
	"encoding/json"
	adv "github.com/coinbase-samples/advanced-trade-sdk-go"
	"net/http"
	"os"
)

func setupClient() (*adv.Client, error) {
	credentials := &adv.Credentials{}
	if err := json.Unmarshal([]byte(os.Getenv("ADV_CREDENTIALS")), credentials); err != nil {
		return nil, err
	}
	client := adv.NewClient(credentials, http.Client{})
	return client, nil
}
