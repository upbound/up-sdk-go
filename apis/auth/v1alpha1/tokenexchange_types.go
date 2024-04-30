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

package v1alpha1

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	AudienceSpacesAPI           = "upbound:spaces:api"
	AudienceSpacesControlPlanes = "upbound:spaces:controlplanes"

	ScopeOrganizationsPrefix = "upbound:org:"

	ContentTypeFormURLEncoded = "application/x-www-form-urlencoded"

	APIGroupAuth        = "tokenexchange.upbound.io"
	APIGroupAuthVersion = "v1alpha1"

	organizationNameRegexString = "^(([a-zA-Z0-9]+-?)*[a-zA-Z0-9])$"

	ParamGrantType          = "grant_type"
	ParamAudience           = "audience"
	ParamScope              = "scope"
	ParamRequestedTokenType = "requested_token_type"
	ParamSubjectToken       = "subject_token"
	ParamSubjectTokenType   = "subject_token_type"
)

var organizationNameRegex = regexp.MustCompile(organizationNameRegexString)

type Scope interface {
	Scope() string
	Type() ScopeType
}

type ScopeParser interface {
	ParseScope(scopeStr string) (Scope, bool)
}

type ScopeType string

const ScopeTypeOrganization ScopeType = "organization"

// Organization describes a name (not numeric ID) of an Upbound organization.
// The name is guaranteed to be unique across other orgs and users, as org and user names are stored in the same "namespaces" table in Upbound API.
type Organization string

func (o Organization) Scope() string   { return ScopeOrganizationsPrefix + string(o) }
func (o Organization) Type() ScopeType { return ScopeTypeOrganization }
func (o Organization) OrganizationName() string {
	return string(o)
}

type OrgParser struct{}

func (OrgParser) ParseScope(str string) (Scope, bool) {
	// Validate that the organization name is between 2-100 characters, plus the prefix length
	if len(str) < 2+len(ScopeOrganizationsPrefix) || len(str) > 100+len(ScopeOrganizationsPrefix) {
		return nil, false
	}

	// Require the prefix to be there
	if !strings.HasPrefix(str, ScopeOrganizationsPrefix) {
		return nil, false
	}

	// Now that we know the scope prefix is there, remove it
	orgName := strings.TrimPrefix(str, ScopeOrganizationsPrefix)

	// The orgName must match the specified format
	if !organizationNameRegex.MatchString(orgName) {
		return nil, false
	}

	return Organization(orgName), true
}

// TODO: Where validate orgName according to prefix, also make that a different type?
func OrganizationScope(orgName string) Organization {
	return Organization(orgName)
}

const (
	// GrantTypeTokenExchange means that the grant_type specifies a token exchange flow
	// ref: https://datatracker.ietf.org/doc/html/rfc8693#section-2.1
	GrantTypeTokenExchange = "urn:ietf:params:oauth:grant-type:token-exchange"

	// TokenTypeIDToken means a token type that is OIDC-compliant
	// ref: https://datatracker.ietf.org/doc/html/rfc8693#TokenTypeIdentifiers
	TokenTypeIDToken = "urn:ietf:params:oauth:token-type:id_token"
)

type TokenExchangeResponse struct {
	AccessToken     string `json:"access_token"`
	IssuedTokenType string `json:"issued_token_type"`
	TokenType       string `json:"token_type"`
	ExpiresIn       int    `json:"expires_in"`
}

type OAuth2ErrorType string

const (
	OAuth2ErrorTypeInvalidRequest       = "invalid_request"
	OAuth2ErrorTypeInvalidScope         = "invalid_scope"
	OAuth2ErrorTypeUnsupportedGrantType = "unsupported_grant_type"
	OAuth2ErrorTypeInvalidClient        = "invalid_client"
	OAuth2ErrorTypeInvalidTarget        = "invalid_target"
)

func NewOAuth2Error(errType OAuth2ErrorType) *OAuth2Error {
	return &OAuth2Error{Type: errType}
}

// OAuth2Error implements https://www.rfc-editor.org/rfc/rfc6749#section-5.2
// The authorization server responds with an HTTP 400 (Bad Request)
// status code (unless specified otherwise) and includes the following
// parameters with the response:
type OAuth2Error struct {
	// REQUIRED.  A single ASCII [USASCII] error code from the OAuth2ErrorType constants
	Type OAuth2ErrorType `json:"error"`

	// OPTIONAL.  Human-readable ASCII [USASCII] text providing
	// additional information, used to assist the client developer in
	// understanding the error that occurred.
	// Values for the "error_description" parameter MUST NOT include
	// characters outside the set %x20-21 / %x23-5B / %x5D-7E.
	Description string `json:"error_description"`
}

func (e *OAuth2Error) Error() string {
	return fmt.Sprintf("type: %s, description: %s", e.Type, e.Description)
}

func (e *OAuth2Error) WithDescription(desc string) *OAuth2Error {
	e.Description = desc
	return e
}
