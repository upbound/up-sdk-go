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

package controlplanes

import (
	"context"
	"net/http"
	"path"

	"github.com/upbound/up-sdk-go"
)

const (
	basePath = "v1/controlPlanes"
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
func (c *Client) Create(ctx context.Context, account string, params *ControlPlaneCreateParameters) (*ControlPlaneResponse, error) {
	req, err := c.Client.NewRequest(ctx, http.MethodPost, basePath, account, params)
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
func (c *Client) Get(ctx context.Context, account, name string) (*ControlPlaneResponse, error) { // nolint:interfacer
	req, err := c.Client.NewRequest(ctx, http.MethodGet, basePath, path.Join(account, name), nil)
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

// List all control planes in the given account on Upbound.
func (c *Client) List(ctx context.Context, account string) ([]ControlPlaneResponse, error) {
	req, err := c.Client.NewRequest(ctx, http.MethodGet, basePath, account, nil)
	if err != nil {
		return nil, err
	}
	cp := []ControlPlaneResponse{}
	err = c.Client.Do(req, &cp)
	if err != nil {
		return nil, err
	}
	return cp, nil
}

// Delete a control plane on Upbound.
func (c *Client) Delete(ctx context.Context, account, name string) error { // nolint:interfacer
	req, err := c.Client.NewRequest(ctx, http.MethodDelete, basePath, path.Join(account, name), nil)
	if err != nil {
		return err
	}
	return c.Client.Do(req, nil)
}
