// Copyright 2024 Upbound Inc
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

package repositorypermission

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"github.com/upbound/up-sdk-go"
)

const (
	basePath = "v1/repoPermissions/%s/teams/%s"
)

// Client is a repositories permission client.
type Client struct {
	*up.Config
}

// NewClient build a repositories permission client from the passed config.
func NewClient(cfg *up.Config) *Client {
	return &Client{
		cfg,
	}
}

// Create assigns a specified permission to a team for a repository on Upbound.
func (c *Client) Create(ctx context.Context, organization string, teamID uuid.UUID, params CreatePermission) error {
	req, err := c.Client.NewRequest(ctx, http.MethodPut, fmt.Sprintf(basePath, organization, teamID), params.Repository, params.Permission)
	if err != nil {
		return err
	}
	return c.Client.Do(req, nil)
}

// Delete removes a specified permission from a team for a repository on Upbound.
func (c *Client) Delete(ctx context.Context, organization string, teamID uuid.UUID, params PermissionIdentifier) error { // nolint:interfacer
	req, err := c.Client.NewRequest(ctx, http.MethodDelete, fmt.Sprintf(basePath, organization, teamID), params.Repository, nil)
	if err != nil {
		return err
	}
	return c.Client.Do(req, nil)
}

// List retrieves all repository permissions assigned to a team on Upbound.
func (c *Client) List(ctx context.Context, organization string, teamID uuid.UUID) (*ListPermissionsResponse, error) {
	req, err := c.Client.NewRequest(ctx, http.MethodGet, fmt.Sprintf(basePath, organization, teamID), "", nil)
	if err != nil {
		return nil, err
	}
	r := &ListPermissionsResponse{}
	if err := c.Client.Do(req, r); err != nil {
		return nil, err
	}
	return r, nil
}
