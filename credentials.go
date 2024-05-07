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
	"encoding/json"
	"fmt"
	"os"
)

type Credentials struct {
	AccessKey     string `json:"accessKey"`
	PrivatePemKey string `json:"privatePemKey"`
	PortfolioId   string `json:"portfolioId"`
}

func UnmarshalCredentials(b []byte) (*Credentials, error) {

	c := &Credentials{}
	if err := json.Unmarshal(b, c); err != nil {
		return nil, err
	}

	return c, nil
}

func ReadEnvCredentials(variableName string) (*Credentials, error) {

	v := os.Getenv(variableName)
	if len(v) == 0 {
		return nil, fmt.Errorf("%s not set as environment variable", variableName)
	}

	return UnmarshalCredentials([]byte(v))
}
