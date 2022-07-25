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

package tokens

import "github.com/google/uuid"

// TokenOwnerType is the type of owner of the token.
type TokenOwnerType string

// Tokens can be owned by a user, control plane, or robot.
const (
	TokenOwnerUser         TokenOwnerType = "users"
	TokenOwnerControlPlane TokenOwnerType = "controlPlanes"
	TokenOwnerRobot        TokenOwnerType = "robots"
)

// bodyType is the type of request in the data body.
type bodyType string

const (
	// The only supported body type is tokens.
	tokenBody bodyType = "tokens"
)

// TokenResponse is the response returned from token operations.
// TODO(hasheddan): consider making token responses strongly typed.
type TokenResponse struct {
	DataSet `json:"data"`
}

// TokensResponse is the response returned from token operations.
// TODO(hasheddan): consider making token responses strongly typed.
type TokensResponse struct { //nolint:golint
	DataSet []DataSet `json:"data"`
}

// RelationshipSet represents set of relationships.
type RelationshipSet map[string]any

// AttributeSet represents set of attributes.
type AttributeSet map[string]any

// Meta represents metadata.
type Meta map[string]any

// DataSet represents a set of data in a token response body.
type DataSet struct {
	Type            string    `json:"type"`
	ID              uuid.UUID `json:"id"`
	AttributeSet    `json:"attributes"`
	RelationshipSet `json:"relationships"`
	Meta            `json:"meta"`
}

// TokenCreateParameters are the parameters for creating a token.
type TokenCreateParameters struct {
	Attributes    TokenAttributes    `json:"attributes"`
	Relationships TokenRelationships `json:"relationships,omitempty"`
}

// tokenCreateParameters is a wrapper around the public TokenCreateParameters
// struct which adds non-configurable fields.
type tokenCreateParameters struct {
	// Type must always be "tokens".
	Type                   bodyType `json:"type"`
	*TokenCreateParameters `json:",inline"`
}

// tokenCreateRequest wraps tokenCreateRequest with the proper request structure
// for the Upbound API.
type tokenCreateRequest struct {
	Data tokenCreateParameters `json:"data"`
}

// TokenUpdateParameters are the parameters for updating a token.
type TokenUpdateParameters struct {
	ID         uuid.UUID       `json:"id"`
	Attributes TokenAttributes `json:"attributes"`
}

// tokenUpdateParameters is a wrapper around the public TokenUpdateParameters
// struct which adds non-configurable fields.
type tokenUpdateParameters struct {
	// Type must always be "tokens".
	Type                   bodyType `json:"type"`
	*TokenUpdateParameters `json:",inline"`
}

// tokenCreateRequest wraps tokenCreateRequest with the proper request structure
// for the Upbound API.
type tokenUpdateRequest struct {
	Data tokenUpdateParameters `json:"data"`
}

// TokenRelationships represents relationships for a token.
type TokenRelationships struct {
	Owner TokenOwner
}

// TokenOwner represents owner of a token.
type TokenOwner struct {
	Data TokenOwnerData `json:"data"`
}

// TokenOwnerData describes a token owner.
type TokenOwnerData struct {
	Type TokenOwnerType `json:"type"`
	ID   uuid.UUID      `json:"id"`
}

// TokenAttributes represents attributes of a token.
type TokenAttributes struct {
	Name string `json:"name"`
}
