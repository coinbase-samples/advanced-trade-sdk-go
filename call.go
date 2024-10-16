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
	"bytes"
	"context"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const emptyQueryParams = ""

type apiRequest struct {
	path                   string
	query                  string
	httpMethod             string
	body                   []byte
	expectedHttpStatusCode int
	client                 Client
}

type apiResponse struct {
	request        *apiRequest
	body           []byte
	httpStatusCode int
	httpStatusMsg  string
	err            error
	errorMessage   *ErrorMessage
}

func generateJwt(method, path, host, keyName, privateKeyPEM string) (string, error) {
	keyBytes := []byte(privateKeyPEM)
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return "", fmt.Errorf("failed to parse PEM block containing the key")
	}

	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("failed to parse EC private key: %w", err)
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
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return signedToken, nil
}

func post(
	ctx context.Context,
	client Client,
	path,
	query string,
	request,
	response interface{},
) error {
	return callPrivate(ctx, client, path, query, http.MethodPost, http.StatusOK, request, response)
}

func getPrivate(
	ctx context.Context,
	client Client,
	path,
	query string,
	request,
	response interface{},
) error {
	return callPrivate(ctx, client, path, query, http.MethodGet, http.StatusOK, request, response)
}

func getPublic(
	ctx context.Context,
	client Client,
	path,
	query string,
	request,
	response interface{},
) error {
	return callPublic(ctx, client, path, query, http.MethodGet, http.StatusOK, request, response)
}

func put(
	ctx context.Context,
	client Client,
	path,
	query string,
	request,
	response interface{},
) error {
	return callPrivate(ctx, client, path, query, http.MethodPut, http.StatusOK, request, response)
}

func del(
	ctx context.Context,
	client Client,
	path,
	query string,
	request,
	response interface{},
) error {
	return callPrivate(ctx, client, path, query, http.MethodDelete, http.StatusOK, request, response)
}

func call(
	ctx context.Context,
	client Client,
	path,
	query,
	httpMethod string,
	expectedHttpStatusCode int,
	request,
	response interface{},
	isPublic bool,
) error {

	if client.Credentials == nil {
		return errors.New("credentials not set")
	}

	body, err := json.Marshal(request)
	if err != nil {
		return err
	}

	resp := makeCall(
		ctx,
		&apiRequest{
			path:                   path,
			query:                  query,
			httpMethod:             httpMethod,
			body:                   body,
			expectedHttpStatusCode: expectedHttpStatusCode,
			client:                 client,
		},
		isPublic,
	)

	if resp.err != nil {
		return resp.err
	}

	if err := json.Unmarshal(resp.body, response); err != nil {
		return err
	}

	return nil
}

func callPublic(
	ctx context.Context,
	client Client,
	path,
	query,
	httpMethod string,
	expectedHttpStatusCode int,
	request,
	response interface{},
) error {
	return call(ctx, client, path, query, httpMethod, expectedHttpStatusCode, request, response, true)
}

func callPrivate(
	ctx context.Context,
	client Client,
	path,
	query,
	httpMethod string,
	expectedHttpStatusCode int,
	request,
	response interface{},
) error {
	return call(ctx, client, path, query, httpMethod, expectedHttpStatusCode, request, response, false)
}

func makeCall(ctx context.Context, request *apiRequest, isPublic bool) *apiResponse {
	response := &apiResponse{
		request: request,
	}

	var jwtToken string

	callUrl := fmt.Sprintf("%s%s%s", request.client.HttpBaseUrl, request.path, request.query)
	parsedUrl, err := url.Parse(callUrl)
	if err != nil {
		response.err = fmt.Errorf("invalid URL: %s - %w", callUrl, err)
		return response
	}

	if !isPublic {
		var err error
		jwtToken, err = generateJwt(request.httpMethod, parsedUrl.Path, parsedUrl.Host, request.client.Credentials.AccessKey, request.client.Credentials.PrivatePemKey)
		if err != nil {
			response.err = fmt.Errorf("failed to generate JWT: %w", err)
			return response
		}
	}

	req, err := http.NewRequestWithContext(ctx, request.httpMethod, callUrl, bytes.NewReader(request.body))
	if err != nil {
		response.err = err
		return response
	}

	req.Header.Add("Accept", "application/json")
	if !isPublic {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", jwtToken))
	}
	res, err := request.client.HttpClient.Do(req)
	if err != nil {
		response.err = err
		return response
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		response.err = err
		return response
	}

	response.body = body
	response.httpStatusCode = res.StatusCode
	response.httpStatusMsg = res.Status

	if request.expectedHttpStatusCode > 0 && res.StatusCode != request.expectedHttpStatusCode {

		var errMsg ErrorMessage
		if strings.Contains(string(response.body), "message") {
			_ = json.Unmarshal(response.body, &errMsg)
			response.errorMessage = &errMsg
		}

		responseMsg := string(body)

		if response.errorMessage != nil && len(response.errorMessage.Value) > 0 {
			responseMsg = response.errorMessage.Value
		}

		response.err = fmt.Errorf(
			"expected status code: %d - received: %d - status msg: %s - url %s - msg: %s",
			request.expectedHttpStatusCode,
			res.StatusCode,
			res.Status,
			callUrl,
			responseMsg,
		)
	}

	return response
}
