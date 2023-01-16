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

package organizations

import (
	"context"
	"fmt"
	"net/http"
	"path"
	"strconv"

	"github.com/upbound/up-sdk-go"
)

const (
	basePath = "v1/organizations"
)

// Client is an organizations client.
type Client struct {
	*up.Config
}

// NewClient builds an organizations client from the passed config.
func NewClient(cfg *up.Config) *Client {
	return &Client{
		cfg,
	}
}

// Create an organization on Upbound.
func (c *Client) Create(ctx context.Context, params *OrganizationCreateParameters) error {
	req, err := c.Client.NewRequest(ctx, http.MethodPost, basePath, "", params)
	if err != nil {
		return err
	}
	return c.Client.Do(req, nil)
}

// Get an organization on Upbound.
func (c *Client) Get(ctx context.Context, id uint) (*Organization, error) {
	req, err := c.Client.NewRequest(ctx, http.MethodGet, basePath, strconv.FormatUint(uint64(id), 10), nil)
	if err != nil {
		return nil, err
	}
	org := &Organization{}
	err = c.Client.Do(req, &org)
	if err != nil {
		return nil, err
	}
	return org, nil
}

// List all organizations for the authenticated user on Upbound.
func (c *Client) List(ctx context.Context) ([]Organization, error) {
	req, err := c.Client.NewRequest(ctx, http.MethodGet, basePath, "", nil)
	if err != nil {
		return nil, err
	}
	orgs := []Organization{}
	err = c.Client.Do(req, &orgs)
	if err != nil {
		return nil, err
	}
	return orgs, nil
}

// ListRobots list all robots the user can access in the organization on
// Upbound.
// TODO(hasheddan): move this to robots client when API is updated.
func (c *Client) ListRobots(ctx context.Context, id uint) ([]Robot, error) {
	req, err := c.Client.NewRequest(ctx, http.MethodGet, basePath, path.Join(strconv.FormatUint(uint64(id), 10), "robots"), nil)
	if err != nil {
		return nil, err
	}
	rs := []Robot{}
	err = c.Client.Do(req, &rs)
	if err != nil {
		return nil, err
	}
	return rs, nil
}

// Delete an organization on Upbound.
func (c *Client) Delete(ctx context.Context, id uint) error {
	req, err := c.Client.NewRequest(ctx, http.MethodDelete, basePath, strconv.FormatUint(uint64(id), 10), nil)
	if err != nil {
		return err
	}
	return c.Client.Do(req, nil)
}

// ListInvites list all invites for the organization on Upbound.
func (c *Client) ListInvites(ctx context.Context, orgID uint) ([]Invite, error) {
	req, err := c.Client.NewRequest(ctx, http.MethodGet, basePath, fmt.Sprintf("%d/invites", orgID), nil)
	if err != nil {
		return nil, err
	}
	invites := []Invite{}
	err = c.Client.Do(req, &invites)
	if err != nil {
		return nil, err
	}
	return invites, nil
}

// DeleteInvite deletes an invite for the organization on Upbound.
func (c *Client) DeleteInvite(ctx context.Context, orgID uint, inviteID uint) error {
	req, err := c.Client.NewRequest(ctx, http.MethodDelete, basePath, fmt.Sprintf("%d/invites/%d", orgID, inviteID), nil)
	if err != nil {
		return err
	}
	return c.Client.Do(req, nil)
}

// CreateInvite creates an invite for the organization on Upbound.
func (c *Client) CreateInvite(ctx context.Context, orgID uint, params *OrganizationInviteCreateParameters) error {
	req, err := c.Client.NewRequest(ctx, http.MethodPost, basePath, fmt.Sprintf("%d/invites", orgID), params)
	if err != nil {
		return err
	}

	return c.Client.Do(req, nil)
}

// ListMembers list all members for the organization on Upbound.
func (c *Client) ListMembers(ctx context.Context, orgID uint) ([]Member, error) {
	req, err := c.Client.NewRequest(ctx, http.MethodGet, basePath, fmt.Sprintf("%d/members", orgID), nil)
	if err != nil {
		return nil, err
	}
	members := []Member{}
	err = c.Client.Do(req, &members)
	if err != nil {
		return nil, err
	}
	return members, nil
}

// RemoveMember removes a member for the organization on Upbound.
func (c *Client) RemoveMember(ctx context.Context, orgID uint, userID uint) error {
	req, err := c.Client.NewRequest(ctx, http.MethodDelete, basePath, fmt.Sprintf("%d/members/%d", orgID, userID), nil)
	if err != nil {
		return err
	}
	return c.Client.Do(req, nil)
}
