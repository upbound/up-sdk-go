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

package robots

import (
	"context"
	"net/http"

	"github.com/google/uuid"

	"github.com/upbound/up-sdk-go"
)

const (
	basePath = "v2/robots"
)

// Client is an robots client.
type Client struct {
	*up.Config
}

// NewClient builds an robots client from the passed config.
func NewClient(cfg *up.Config) *Client {
	return &Client{
		cfg,
	}
}

// Create a robot on Upbound.
func (c *Client) Create(ctx context.Context, params *RobotCreateParameters) (*RobotResponse, error) {
	body := &robotCreateRequest{
		Data: robotCreateParameters{
			Type:                  robotBody,
			RobotCreateParameters: params,
		},
	}
	req, err := c.Client.NewRequest(ctx, http.MethodPost, basePath, "", body)
	if err != nil {
		return nil, err
	}
	t := &RobotResponse{}
	err = c.Client.Do(req, &t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

// Get gets a robot on Upbound.
func (c *Client) Get(ctx context.Context, id uuid.UUID) (*RobotResponse, error) { // nolint:interfacer
	req, err := c.Client.NewRequest(ctx, http.MethodDelete, basePath, id.String(), nil)
	if err != nil {
		return nil, err
	}
	r := &RobotResponse{}
	if err := c.Client.Do(req, r); err != nil {
		return nil, err
	}
	return r, nil
}

// Delete an robot on Upbound.
func (c *Client) Delete(ctx context.Context, id uuid.UUID) error { // nolint:interfacer
	req, err := c.Client.NewRequest(ctx, http.MethodDelete, basePath, id.String(), nil)
	if err != nil {
		return err
	}
	return c.Client.Do(req, nil)
}
