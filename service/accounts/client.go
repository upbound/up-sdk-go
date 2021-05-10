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

package accounts

import (
	"context"
	"net/http"
	"path"

	"github.com/upbound/up-sdk-go"
	"github.com/upbound/up-sdk-go/service/controlplanes"
)

const (
	basePath          = "v1/accounts"
	controlPlanesPath = "controlPlanes"
)

// Client is a accounts client.
type Client struct {
	*up.Config
}

// NewClient builds a accounts client from the passed config.
func NewClient(cfg *up.Config) *Client {
	return &Client{
		cfg,
	}
}

// Get a account on Upbound Cloud.
func (c *Client) Get(ctx context.Context, name string) (*AccountResponse, error) {
	req, err := c.Client.NewRequest(ctx, http.MethodGet, basePath, name, nil)
	if err != nil {
		return nil, err
	}
	ns := &AccountResponse{}
	err = c.Client.Do(req, &ns)
	if err != nil {
		return nil, err
	}
	return ns, nil
}

// List all accounts for the authenticated user on Upbound Cloud.
func (c *Client) List(ctx context.Context) ([]AccountResponse, error) {
	req, err := c.Client.NewRequest(ctx, http.MethodGet, basePath, "", nil)
	if err != nil {
		return nil, err
	}
	ns := []AccountResponse{}
	err = c.Client.Do(req, &ns)
	if err != nil {
		return nil, err
	}
	return ns, nil
}

// ListControlPlanes lists all control planes in the given account on Upbound Cloud.
func (c *Client) ListControlPlanes(ctx context.Context, name string) ([]controlplanes.ControlPlaneResponse, error) {
	req, err := c.Client.NewRequest(ctx, http.MethodGet, basePath, path.Join(name, controlPlanesPath), nil)
	if err != nil {
		return nil, err
	}
	cp := []controlplanes.ControlPlaneResponse{}
	err = c.Client.Do(req, &cp)
	if err != nil {
		return nil, err
	}
	return cp, nil
}
