// Copyright 2022 Upbound Inc
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

package repositories

import (
	"context"
	"net/http"
	"path"

	"github.com/upbound/up-sdk-go"
	"github.com/upbound/up-sdk-go/service/common"
)

const (
	basePath = "v1/repositories"
)

// Client is a repositories client.
type Client struct {
	*up.Config
}

// NewClient build a repositories client from the passed config.
func NewClient(cfg *up.Config) *Client {
	return &Client{
		cfg,
	}
}

// CreateOrUpdate a repository on Upbound.
func (c *Client) CreateOrUpdate(ctx context.Context, account string, name string) error {
	// TODO(hasheddan): allow passing parameters in body when supported by API.
	// For now, a body is expected, but it may be empty.
	req, err := c.Client.NewRequest(ctx, http.MethodPut, basePath, path.Join(account, name), struct{}{})
	if err != nil {
		return err
	}
	return c.Client.Do(req, nil)
}

// Get a repository on Upbound.
func (c *Client) Get(ctx context.Context, account, name string) (*RepositoryResponse, error) { // nolint:interfacer
	req, err := c.Client.NewRequest(ctx, http.MethodGet, basePath, path.Join(account, name), nil)
	if err != nil {
		return nil, err
	}
	r := &RepositoryResponse{}
	if err := c.Client.Do(req, r); err != nil {
		return nil, err
	}
	return r, nil
}

// List all repositories in the given account on Upbound.
func (c *Client) List(ctx context.Context, account string, opts ...common.ListOption) (*RepositoryListResponse, error) {
	req, err := c.Client.NewRequest(ctx, http.MethodGet, basePath, account, nil)
	if err != nil {
		return nil, err
	}
	for _, o := range opts {
		o(req)
	}
	r := &RepositoryListResponse{}
	if err := c.Client.Do(req, r); err != nil {
		return nil, err
	}
	return r, nil
}

// Delete a repository on Upbound.
func (c *Client) Delete(ctx context.Context, account, name string) error { // nolint:interfacer
	req, err := c.Client.NewRequest(ctx, http.MethodDelete, basePath, path.Join(account, name), nil)
	if err != nil {
		return err
	}
	return c.Client.Do(req, nil)
}
