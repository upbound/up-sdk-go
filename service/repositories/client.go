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

// CreateOrUpdate creates or updates a repository on Upbound. If a new
// repository is created, it will be public. If the repository exists, it will
// be made public.
//
// Deprecated: use CreateOrUpdateWithOpts instead.
func (c *Client) CreateOrUpdate(ctx context.Context, account string, name string) error {
	return c.CreateOrUpdateWithOptions(ctx, account, name, WithPublic())
}

// CreateOrUpdateWithOptions creates or updates a repository on
// Upbound. Repositories will be created as or made private by default.
func (c *Client) CreateOrUpdateWithOptions(ctx context.Context, account, name string, opts ...CreateOrUpdateOption) error {
	body := &RepositoryCreateOrUpdateRequest{}
	for _, opt := range opts {
		opt(body)
	}

	req, err := c.Client.NewRequest(ctx, http.MethodPut, basePath, path.Join(account, name), body)
	if err != nil {
		return err
	}
	return c.Client.Do(req, nil)
}

// CreateOrUpdateOption is an option for the CreateOrUpdateWithOptions method.
type CreateOrUpdateOption func(*RepositoryCreateOrUpdateRequest)

// WithPrivate makes a repository private.
func WithPrivate() CreateOrUpdateOption {
	return func(req *RepositoryCreateOrUpdateRequest) {
		req.Public = false
	}
}

// WithPublic makes a repository public.
func WithPublic() CreateOrUpdateOption {
	return func(req *RepositoryCreateOrUpdateRequest) {
		req.Public = true
	}
}

// WithDraft stops a repository from indexing package versions to the Upbound Marketplace.
func WithDraft() CreateOrUpdateOption {
	return func(req *RepositoryCreateOrUpdateRequest) {
		req.Publish = false
	}
}

// WithPublish makes a repository index package versions to the Upbound Marketplace.
func WithPublish() CreateOrUpdateOption {
	return func(req *RepositoryCreateOrUpdateRequest) {
		req.Publish = true
	}
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
