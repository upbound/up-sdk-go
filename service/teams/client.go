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

package teams

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/upbound/up-sdk-go"
)

const (
	basePath = "v1/teams"
)

// Client is an teams client.
type Client struct {
	*up.Config
}

// NewClient builds an teams client from the passed config.
func NewClient(cfg *up.Config) *Client {
	return &Client{
		cfg,
	}
}

// Create creates a team on Upbound.
func (c *Client) Create(ctx context.Context, params *TeamCreateParameters) (*TeamResponse, error) { // nolint:interfacer
	req, err := c.Client.NewRequest(ctx, http.MethodPost, basePath, "", params)
	if err != nil {
		return nil, err
	}
	t := &TeamResponse{}
	err = c.Client.Do(req, &t)
	if err != nil {
		return nil, errors.Wrap(err, "cannot build request")
	}
	return t, nil
}

// Get gets a team on Upbound.
func (c *Client) Get(ctx context.Context, id uuid.UUID) (*TeamResponse, error) { // nolint:interfacer
	req, err := c.Client.NewRequest(ctx, http.MethodGet, basePath, id.String(), nil)
	if err != nil {
		return nil, err
	}
	t := &TeamResponse{}
	return t, c.Client.Do(req, t)
}

// Delete delete an team on Upbound.
func (c *Client) Delete(ctx context.Context, id uuid.UUID) error { // nolint:interfacer
	req, err := c.Client.NewRequest(ctx, http.MethodDelete, basePath, id.String(), nil)
	if err != nil {
		return err
	}
	return c.Client.Do(req, nil)
}
