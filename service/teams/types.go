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
	"time"

	"github.com/google/uuid"
)

// TeamCreateParameters are the parameters for creating a team.
type TeamCreateParameters struct {
	Name           string `json:"name"`
	OrganizationID uint   `json:"organizationId"`
}

// TeamResponse is the response returned from team operations.
type TeamResponse struct {
	ID             uuid.UUID  `json:"id"`
	OrganizationID uint       `json:"organizationId"`
	AccountID      uint       `json:"accountId"`
	Name           string     `json:"name"`
	CreatorID      uint       `json:"creatorId"`
	CreatedAt      *time.Time `json:"createdAt,omitempty"`
}

// TeamAttributes are the attributes of a team.
type TeamAttributes struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
