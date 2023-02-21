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
	"time"

	"github.com/google/uuid"
)

// ConfigurationResponse represents the concept of a Configuration on Upbound.
// It is used to configure a Managed Control Plane with a set of API types
// and optional extensions.
type ConfigurationResponse struct {
	ID            uuid.UUID  `json:"id"`
	Name          *string    `json:"name"`
	LatestVersion *string    `json:"latestVersion,omitempty"`
	TemplateID    string     `json:"templateID"`
	Provider      string     `json:"provider"`
	Context       string     `json:"context"`
	Repo          string     `json:"repo"`
	Branch        string     `json:"branch"`
	CreatorID     uint       `json:"creatorId"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     *time.Time `json:"updatedAt,omitempty"`
	SyncedAt      *time.Time `json:"syncedAt,omitempty"`
}

// ConfigurationListResponse is a list of all configurations belonging to the account.
type ConfigurationListResponse struct {
	Configurations []ConfigurationResponse `json:"configurations"`
}
