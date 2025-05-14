// Copyright 2025 Upbound Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package users contains client and functions for user endpoint.
package users

import (
	"context"
	"net/http"
	"path"
	"strconv"

	"github.com/upbound/up-sdk-go"
	"github.com/upbound/up-sdk-go/service/tokens"
)

const (
	basePath   = "/v1/users"
	tokensPath = "tokens"
)

// Client is an users client.
type Client struct {
	*up.Config
}

// NewClient builds an user client from the passed config.
func NewClient(cfg *up.Config) *Client {
	return &Client{
		cfg,
	}
}

// ListTokens lists tokens for a user on Upbound.
func (c *Client) ListTokens(ctx context.Context, userID uint) (*tokens.TokensResponse, error) {
	user := strconv.FormatUint(uint64(userID), 10)
	req, err := c.Client.NewRequest(ctx, http.MethodGet, basePath, path.Join(user, tokensPath), nil)
	if err != nil {
		return nil, err
	}
	r := &tokens.TokensResponse{}
	if err := c.Client.Do(req, r); err != nil {
		return nil, err
	}
	return r, nil
}
