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
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/coinbase-samples/advanced-trade-sdk-go/credentials"
	"github.com/coinbase-samples/core-go"
	"github.com/golang-jwt/jwt/v5"
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

	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", fmt.Sprintf("coinbase-advanced-go/%s", sdkVersion))

	jwtToken, err := generateJwt(req.Method, path, req.Host, c.Credentials().AccessKey, c.Credentials().PrivatePemKey)
	if err != nil {
		// core-go's header hook can't return an error, so we can't abort the
		// request from here. Log the cause and leave Authorization unset; the
		// API rejects it with 401, which the caller gets as a normal error
		// instead of the whole process being killed.
		log.Printf("failed to generate JWT: %v", err)
		return
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", jwtToken))

}

func generateJwt(method, path, host, keyName, privateKey string) (string, error) {
	key, signingMethod, err := parsePrivateKey(privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %w", err)
	}

	now := time.Now()
	claims := jwt.MapClaims{
		"sub": keyName,
		"iss": "coinbase-cloud",
		"nbf": now.Unix(),
		"exp": now.Add(2 * time.Minute).Unix(),
		"uri": fmt.Sprintf("%s %s%s", method, host, path),
	}

	token := jwt.NewWithClaims(signingMethod, claims)
	token.Header["kid"] = keyName
	token.Header["nonce"] = uuid.New().String()

	signedToken, err := token.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return signedToken, nil
}

// Two key types are supported:
//   - ECDSA (P-256), signed with ES256. Provided as a PEM-encoded key
//     ("-----BEGIN EC PRIVATE KEY-----" or a PKCS#8 "-----BEGIN PRIVATE KEY-----").
//   - Ed25519, signed with EdDSA. Provided either as a PKCS#8 PEM key or as a
//     base64-encoded raw key (32-byte seed or 64-byte seed+public key), which is
//     how the CDP portal hands out Ed25519 secrets.
func parsePrivateKey(privateKey string) (any, jwt.SigningMethod, error) {
	if strings.HasPrefix(strings.TrimSpace(privateKey), "-----BEGIN") {
		block, _ := pem.Decode([]byte(privateKey))
		if block == nil {
			return nil, nil, fmt.Errorf("failed to decode PEM block containing the private key")
		}

		if block.Type == "EC PRIVATE KEY" {
			key, err := x509.ParseECPrivateKey(block.Bytes)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to parse EC private key: %w", err)
			}
			return key, jwt.SigningMethodES256, nil
		}

		parsed, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse PKCS#8 private key: %w", err)
		}
		switch k := parsed.(type) {
		case *ecdsa.PrivateKey:
			return k, jwt.SigningMethodES256, nil
		case ed25519.PrivateKey:
			return k, jwt.SigningMethodEdDSA, nil
		default:
			return nil, nil, fmt.Errorf("unsupported private key type %T; expected ECDSA (P-256) or Ed25519", parsed)
		}
	}

	// Not PEM: treat the secret as a base64-encoded raw Ed25519 key. Strip any
	raw, err := base64.StdEncoding.DecodeString(strings.Join(strings.Fields(privateKey), ""))
	if err != nil {
		return nil, nil, fmt.Errorf("private key is neither PEM nor valid base64: %w", err)
	}
	switch len(raw) {
	case ed25519.SeedSize:
		return ed25519.NewKeyFromSeed(raw), jwt.SigningMethodEdDSA, nil
	case ed25519.PrivateKeySize:
		return ed25519.PrivateKey(raw), jwt.SigningMethodEdDSA, nil
	default:
		return nil, nil, fmt.Errorf("ed25519 raw key must decode to %d or %d bytes, got %d", ed25519.SeedSize, ed25519.PrivateKeySize, len(raw))
	}
}

func DefaultHttpClient() (http.Client, error) {
	return core.DefaultHttpClient()
}
