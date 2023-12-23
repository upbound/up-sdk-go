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

package configurations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/upbound/up-sdk-go"
)

const (
	basePath          = "v1/configurations"
	templatesBasePath = "v1/configurationTemplates"
)

// Client is a configurations client.
type Client struct {
	*up.Config
}

// NewClient builds a configurations client from the given config.
func NewClient(cfg *up.Config) *Client {
	return &Client{cfg}
}

// List all configurations for an account on Upbound.
func (c *Client) List(ctx context.Context, account string) (*ConfigurationListResponse, error) {
	// Call listOnePage repeatedly to get all configurations.
	configurations := &ConfigurationListResponse{}
	page := 0
	for {
		configs, err := c.listOnePage(ctx, account, page)
		if err != nil {
			return nil, err
		}
		configurations.Configurations = append(configurations.Configurations, configs.Configurations...)
		configurations.Count += configs.Count
		if configs.Count == 0 || configs.Count < configs.Size {
			break
		}
		page++
	}
	return configurations, nil
}

// listOnePage gets one page of configuration.
func (c *Client) listOnePage(ctx context.Context, account string, page int) (*ConfigurationListResponse, error) {
	req, err := c.Client.NewRequest(ctx, http.MethodGet, basePath, fmt.Sprintf("%s?page=%d", account, page), nil)
	if err != nil {
		return nil, err
	}
	configurations := &ConfigurationListResponse{}
	err = c.Client.Do(req, &configurations)
	if err != nil {
		return nil, err
	}
	return configurations, nil
}

// Get a configuration for an account by name on Upbound.
func (c *Client) Get(ctx context.Context, account, name string) (*ConfigurationResponse, error) {
	req, err := c.Client.NewRequest(ctx, http.MethodGet, basePath, fmt.Sprintf("%s/%s", account, name), nil)
	if err != nil {
		return nil, err
	}
	configuration := &ConfigurationResponse{}
	err = c.Client.Do(req, &configuration)
	if err != nil {
		return nil, err
	}
	return configuration, nil
}

// Create a configuration on Upbound
// Note that when using GitHub, this will make a new repo, and that requires
// the GitHub app to be authorized and installed. That's a separate API, which
// begins with the gitsources login. The full login is implemented in the CLI (https://github.com/upbound/up).
func (c *Client) Create(ctx context.Context, account string, params *ConfigurationCreateParameters) (*ConfigurationResponse, error) {
	req, err := c.Client.NewRequest(ctx, http.MethodPost, basePath, account, params)
	if err != nil {
		return nil, err
	}
	configuration := &ConfigurationResponse{}
	err = c.Client.Do(req, &configuration)
	if err != nil {
		return nil, err
	}
	return configuration, nil
}

// Delete a configuration on an Upbound account by name.
// This operation can potentially orphan Managed Control Planes that have deployed
// the configuration, as they can no longer update them. The API will return an
// HTTP 403 status code in this case indicating that the configuration is still in use.
func (c *Client) Delete(ctx context.Context, account, name string) error {
	req, err := c.Client.NewRequest(ctx, http.MethodDelete, basePath, fmt.Sprintf("%s/%s", account, name), nil)
	if err != nil {
		return err
	}
	return c.Client.Do(req, nil)
}

// ListTemplates all configuration Templates that can be used to make a configuration
func (c *Client) ListTemplates(ctx context.Context) (*ConfigurationTemplateListResponse, error) {
	req, err := c.Client.NewRequest(ctx, http.MethodGet, templatesBasePath, "", nil)
	if err != nil {
		return nil, err
	}
	templates := &ConfigurationTemplateListResponse{}
	err = c.Client.Do(req, &templates)
	if err != nil {
		return nil, err
	}
	return templates, nil
}
