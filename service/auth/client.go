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

package auth

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/upbound/up-sdk-go"
	"github.com/upbound/up-sdk-go/apis/auth/v1alpha1"
)

var (
	basePath = fmt.Sprintf("/apis/%s/%s", v1alpha1.APIGroupAuth, v1alpha1.APIGroupAuthVersion)
)

// Client is a tokenexchange client.
type Client struct {
	*up.Config
}

// NewClient builds a tokenexchange client from the passed config.
func NewClient(cfg *up.Config) *Client {
	return &Client{
		cfg,
	}
}

// GetOrgScopedToken a token on Upbound.
func (c *Client) GetOrgScopedToken(ctx context.Context, org, token string) (*v1alpha1.TokenExchangeResponse, error) { // nolint:interfacer
	body := url.Values{
		v1alpha1.ParamAudience:         []string{v1alpha1.AudienceSpacesAPI, v1alpha1.AudienceSpacesControlPlanes},
		v1alpha1.ParamGrantType:        []string{v1alpha1.GrantTypeTokenExchange},
		v1alpha1.ParamSubjectTokenType: []string{v1alpha1.TokenTypeIDToken},
		v1alpha1.ParamSubjectToken:     []string{token},
		v1alpha1.ParamScope:            []string{fmt.Sprintf("%s%s", v1alpha1.ScopeOrganizationsPrefix, org)},
	}

	req, err := c.Client.NewRequest(ctx, http.MethodPost, basePath, "orgscopedtokens", nil)
	if err != nil {
		return nil, err
	}

	reader := strings.NewReader(body.Encode())

	req.ContentLength = int64(reader.Len())
	req.Body = io.NopCloser(reader)
	req.GetBody = func() (io.ReadCloser, error) {
		return io.NopCloser(reader), nil
	}
	req.Header.Set("Content-Type", v1alpha1.ContentTypeFormURLEncoded)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	t := &v1alpha1.TokenExchangeResponse{}
	err = c.Client.Do(req, &t)
	if err != nil {
		return nil, err
	}
	return t, nil
}
