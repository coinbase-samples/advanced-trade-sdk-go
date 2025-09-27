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

package client

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/coinbase-samples/advanced-trade-sdk-go/credentials"
	"github.com/coinbase-samples/core-go"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

var defaultV3ApiBaseUrl = "https://api.coinbase.com/api/v3"

var DefaultSuccessHttpStatusCodes = []int{http.StatusOK}

var defaultHeadersFunc = AddAdvancedHttpHeaders

type RestClient interface {
	SetHttpBaseUrl(u string) RestClient
	HttpBaseUrl() string
	HttpClient() *http.Client
	Credentials() *credentials.Credentials
	SetHeadersFunc(hf core.HttpHeaderFunc) RestClient
	HeadersFunc() core.HttpHeaderFunc
}

func NewRestClient(credentials *credentials.Credentials, httpClient http.Client) RestClient {
	return &restClientImpl{
		credentials: credentials,
		httpClient:  httpClient,
		baseUrl:     defaultV3ApiBaseUrl,
		headersFunc: defaultHeadersFunc,
	}
}

type restClientImpl struct {
	httpClient  http.Client
	credentials *credentials.Credentials
	headersFunc core.HttpHeaderFunc
	baseUrl     string
}

func (c *restClientImpl) HttpBaseUrl() string {
	return c.baseUrl
}

func (c *restClientImpl) SetHttpBaseUrl(u string) RestClient {
	c.baseUrl = u
	return c
}

func (c *restClientImpl) HttpClient() *http.Client {
	return &c.httpClient
}

func (c *restClientImpl) Credentials() *credentials.Credentials {
	return c.credentials
}

func (c *restClientImpl) SetHeadersFunc(hf core.HttpHeaderFunc) RestClient {
	c.headersFunc = hf
	return c
}

func (c *restClientImpl) HeadersFunc() core.HttpHeaderFunc {
	return c.headersFunc
}

func AddAdvancedHttpHeaders(req *http.Request, path string, body []byte, cl core.RestClient, t time.Time) {

	c := cl.(*restClientImpl)

	jwtToken := generateJwt(req.Method, path, req.Host, c.Credentials().AccessKey, c.Credentials().PrivatePemKey)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", jwtToken))

}

func generateJwt(method, path, host, keyName, privateKeyPEM string) string {
	keyBytes := []byte(privateKeyPEM)
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		log.Fatal("failed to parse PEM block containing the key")
	}

	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		log.Fatalf("failed to parse EC private key: %v", err)
	}

	now := time.Now()
	claims := jwt.MapClaims{
		"sub": keyName,
		"iss": "coinbase-cloud",
		"nbf": now.Unix(),
		"exp": now.Add(2 * time.Minute).Unix(),
		"uri": fmt.Sprintf("%s %s%s", method, host, path),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token.Header["kid"] = keyName
	token.Header["nonce"] = uuid.New().String()

	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		log.Fatalf("failed to sign token: %v", err)
	}

	return signedToken
}

func DefaultHttpClient() (http.Client, error) {
	return core.DefaultHttpClient()
}

// HttpPost wraps core.HttpPost with error sanitization to handle HTML responses
func HttpPost(
	ctx context.Context,
	client core.RestClient,
	path string,
	queryParams string,
	successHttpStatusCodes []int,
	request interface{},
	response interface{},
	headersFunc core.HttpHeaderFunc,
) error {
	err := core.HttpPost(
		ctx,
		client,
		path,
		queryParams,
		successHttpStatusCodes,
		request,
		response,
		headersFunc,
	)

	return SanitizeError(err)
}

// HttpGet wraps core.HttpGet with error sanitization to handle HTML responses
func HttpGet(
	ctx context.Context,
	client core.RestClient,
	path string,
	queryParams string,
	successHttpStatusCodes []int,
	request interface{},
	response interface{},
	headersFunc core.HttpHeaderFunc,
) error {
	err := core.HttpGet(
		ctx,
		client,
		path,
		queryParams,
		successHttpStatusCodes,
		request,
		response,
		headersFunc,
	)

	return SanitizeError(err)
}
