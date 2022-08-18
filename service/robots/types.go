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
	"github.com/upbound/up-sdk-go/service/common"
)

// RobotOwnerType is the type of owner of the robot.
type RobotOwnerType string

// Robots can be owned by an organization.
const (
	RobotOwnerOrganization RobotOwnerType = "organization"
)

// bodyType is the type of request in the data body.
type bodyType string

const (
	// The only supported body type is robots.
	robotBody bodyType = "robots"
)

// RobotsResponse is the response returned from robot operations.
// TODO(hasheddan): consider making robot responses strongly typed.
type RobotsResponse struct { //nolint:golint
	DataSet []common.DataSet `json:"data"`
}

// RobotCreateParameters are the parameters for creating a robot.
type RobotCreateParameters struct {
	Attributes    RobotAttributes    `json:"attributes"`
	Relationships RobotRelationships `json:"relationships,omitempty"`
}

// robotCreateParameters is a wrapper around the public RobotCreateParameters
// struct which adds non-configurable fields.
type robotCreateParameters struct {
	// Type must always be "robots".
	Type                   bodyType `json:"type"`
	*RobotCreateParameters `json:",inline"`
}

// robotCreateRequest wraps robotCreateRequest with the proper request structure
// for the Upbound API.
type robotCreateRequest struct {
	Data robotCreateParameters `json:"data"`
}

// RobotRelationships represents relationships for a robot.
type RobotRelationships struct {
	Owner RobotOwner `json:"organization"`
}

// RobotOwner represents owner of a robot.
type RobotOwner struct {
	Data RobotOwnerData `json:"data"`
}

// RobotOwnerData describes a robot owner.
type RobotOwnerData struct {
	Type RobotOwnerType `json:"type"`
	ID   string         `json:"id"`
}

// RobotResponse is the response returned from robot operations.
// TODO(hasheddan): consider making robot responses strongly typed.
type RobotResponse struct {
	common.DataSet `json:"data"`
}

// RobotAttributes are the attributes of a robot.
type RobotAttributes struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
