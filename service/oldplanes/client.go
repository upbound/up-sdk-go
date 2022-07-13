// Copyright 2021 Upbound Inc
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

// DEPRECATED(hasheddan): please use the controlplanes package.

package oldplanes

import (
	"context"
	"net/http"
	"path"

	"github.com/google/uuid"

	"github.com/upbound/up-sdk-go"
	"github.com/upbound/up-sdk-go/service/tokens"
)

const (
	basePath     = "v1/controlPlanes"
	tokensPath   = "tokens"
	viewOnlyPath = "viewOnly"
)

// Client is a control planes client.
type Client struct {
	*up.Config
}

// NewClient build a control planes client from the passed config.
func NewClient(cfg *up.Config) *Client {
	return &Client{
		cfg,
	}
}

// Create a control plane on Upbound.
func (c *Client) Create(ctx context.Context, params *ControlPlaneCreateParameters) (*ControlPlaneResponse, error) {
	req, err := c.Client.NewRequest(ctx, http.MethodPost, basePath, "", params)
	if err != nil {
		return nil, err
	}
	cp := &ControlPlaneResponse{}
	err = c.Client.Do(req, &cp)
	if err != nil {
		return nil, err
	}
	return cp, nil
}

// Get a control plane on Upbound.
func (c *Client) Get(ctx context.Context, id uuid.UUID) (*ControlPlaneResponse, error) { // nolint:interfacer
	req, err := c.Client.NewRequest(ctx, http.MethodGet, basePath, id.String(), nil)
	if err != nil {
		return nil, err
	}
	cp := &ControlPlaneResponse{}
	err = c.Client.Do(req, &cp)
	if err != nil {
		return nil, err
	}
	return cp, nil
}

// GetTokens a control plane on Upbound.
func (c *Client) GetTokens(ctx context.Context, id uuid.UUID) (*tokens.TokensResponse, error) { // nolint:interfacer
	req, err := c.Client.NewRequest(ctx, http.MethodGet, basePath, path.Join(id.String(), tokensPath), nil)
	if err != nil {
		return nil, err
	}
	t := &tokens.TokensResponse{}
	err = c.Client.Do(req, &t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

// Delete a control plane on Upbound.
func (c *Client) Delete(ctx context.Context, id uuid.UUID) error { // nolint:interfacer
	req, err := c.Client.NewRequest(ctx, http.MethodDelete, basePath, id.String(), nil)
	if err != nil {
		return err
	}
	return c.Client.Do(req, nil)
}

// SetViewOnly sets the view-only value of a control plane on Upbound.
func (c *Client) SetViewOnly(ctx context.Context, id uuid.UUID, viewOnly bool) error { // nolint:interfacer
	req, err := c.Client.NewRequest(ctx, http.MethodPut, basePath, path.Join(id.String(), viewOnlyPath), &controlPlaneViewOnlyParameters{
		IsViewOnly: viewOnly,
	})
	if err != nil {
		return err
	}
	return c.Client.Do(req, nil)
}
