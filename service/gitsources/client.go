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

package gitsources

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/upbound/up-sdk-go"
)

const (
	basePath  = "v1/gitSources"
	loginPath = "github/client/login"
	portParam = "cli"
)

// Client is a gitsources client.
type Client struct {
	*up.Config
}

// NewClient builds a configurations client from the given config.
func NewClient(cfg *up.Config) *Client {
	return &Client{cfg}
}

// Login does a gitsources login.
func (c *Client) Login(ctx context.Context, port int) (LoginResponse, error) {
	var loginResponse LoginResponse

	// We have to use the HTTPClient because unlike other Upbound APIs,
	// the body will be HTML. We need to extract the data we need from
	// the status code and the headers
	hc, ok := c.Client.(*up.HTTPClient)
	if !ok {
		return loginResponse, errors.New("Unexpected error: client isn't an HTTP client")
	}

	// The HTTPClient follows redirects, but we don't want that -- we
	// need to know what the redirect is. The CLI will take different
	// actions if the redirect goes to github.com (which indicates we
	// need to login) or Upbound (which indicates we already logged in, or there
	// is a problem).
	oldRedirect := hc.HTTP.CheckRedirect
	hc.HTTP.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	path := loginPath
	if port != 0 {
		path += fmt.Sprintf("?%s=%d", portParam, port)
	}
	req, err := hc.NewRequest(ctx, http.MethodGet, basePath, path, nil)
	if err != nil {
		return loginResponse, err
	}

	response, err := hc.HTTP.Do(req)
	hc.HTTP.CheckRedirect = oldRedirect
	if err != nil {
		return loginResponse, err
	}
	defer response.Body.Close() // nolint:errcheck
	loginResponse.StatusCode = response.StatusCode
	loginResponse.RedirectURL, err = url.Parse(response.Header.Get("location"))
	return loginResponse, err
}
