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

package teams

import (
	"github.com/upbound/up-sdk-go/service/common"
)

// TeamCreateParameters are the parameters for creating a team.
type TeamCreateParameters struct {
	Name           string `json:"name"`
	OrganizationID uint   `json:"organizationId"`
}

// TeamsResponse is the response returned from team operations.
type TeamsResponse struct { //nolint:golint
	DataSet []common.DataSet `json:"data"`
}

// TeamResponse is the response returned from team operations.
type TeamResponse struct {
	common.DataSet `json:"data"`
}

// TeamAttributes are the attributes of a team.
type TeamAttributes struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
