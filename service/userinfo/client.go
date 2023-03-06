// Copyright 2023 Upbound Inc
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

package userinfo

import (
	"context"
	"net/http"

	"github.com/upbound/up-sdk-go"
)

const (
	basePath = "v1/self"
)

// Client is a user info client.
type Client struct {
	*up.Config
}

// NewClient builds a user info client from the passed config.
func NewClient(cfg *up.Config) *Client {
	return &Client{
		cfg,
	}
}

// Get gets user information on Upbound.
func (c *Client) Get(ctx context.Context) (*GetResponse, error) { // nolint:interfacer
	req, err := c.Client.NewRequest(ctx, http.MethodGet, basePath, "", nil)
	if err != nil {
		return nil, err
	}
	r := &GetResponse{}
	if err := c.Client.Do(req, r); err != nil {
		return nil, err
	}
	return r, nil
}
