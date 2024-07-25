// Copyright 2024 Upbound Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package auth

const (
	// ScopeOrganizationsPrefix is the required prefix for an upbound
	// organization scope.
	ScopeOrganizationsPrefix = "upbound:org:"

	// ContentTypeFormURLEncoded is the form URL content type accepted by the
	// auth service.
	ContentTypeFormURLEncoded = "application/x-www-form-urlencoded"

	// APIGroupAuth is the resource group for auth requests.
	APIGroupAuth = "tokenexchange.upbound.io"

	// APIGroupAuthVersion is the resource version for auth requests.
	APIGroupAuthVersion = "v1alpha1"
)

const (
	// AudienceSpacesAPI is the scope for accessing a space itself.
	AudienceSpacesAPI = "upbound:spaces:api"

	// AudienceSpacesControlPlanes is the scope for accessing control planes
	// within the space.
	AudienceSpacesControlPlanes = "upbound:spaces:controlplanes"
)

const (
	// ParamGrantType specifies which type of authorization type is being
	// granted.
	ParamGrantType = "grant_type"

	// ParamAudience specifies which (possibly multiple) audiences should be
	// granted in an exchange.
	ParamAudience = "audience"

	// ParamScope specifies the scope of the grant for the given token, such as
	// an organization.
	ParamScope = "scope"

	// ParamSubjectToken is the subject up for exchange, such as a session
	// token.
	ParamSubjectToken = "subject_token"

	// ParamSubjectTokenType is the type of the subject being used for the
	// exchange.
	ParamSubjectTokenType = "subject_token_type"
)

const (
	// GrantTypeTokenExchange means that the grant_type specifies a token exchange flow
	// ref: https://datatracker.ietf.org/doc/html/rfc8693#section-2.1
	GrantTypeTokenExchange = "urn:ietf:params:oauth:grant-type:token-exchange"

	// TokenTypeIDToken means a token type that is OIDC-compliant
	// ref: https://datatracker.ietf.org/doc/html/rfc8693#TokenTypeIdentifiers
	TokenTypeIDToken = "urn:ietf:params:oauth:token-type:id_token"
)

// TokenExchangeResponse defines the response from the server when completing a
// successful token exchange request
type TokenExchangeResponse struct {
	AccessToken     string `json:"access_token"`
	IssuedTokenType string `json:"issued_token_type"`
	TokenType       string `json:"token_type"`
	ExpiresIn       int    `json:"expires_in"`
}
