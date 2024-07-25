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

package repositorypermission

import (
	"time"

	"github.com/google/uuid"
)

// PermissionType represents the type of permission for a repository.
type PermissionType string

// PermissionTypes
const (
	PermissionAdmin PermissionType = "admin"
	PermissionRead  PermissionType = "read"
	PermissionWrite PermissionType = "write"
	PermissionView  PermissionType = "view"
)

// RepositoryPermission represents the permission to be set for a repository.
type RepositoryPermission struct {
	Permission PermissionType `json:"permission"`
}

// CreatePermission holds the parameters for creating a repository permission.
type CreatePermission struct {
	Permission RepositoryPermission `json:"permission"`
	Repository string               `json:"repository"`
}

// PermissionIdentifier holds the parameters for a repository permission.
type PermissionIdentifier struct {
	Repository string `json:"repository"`
}

// Permission represents a repository permission entry.
type Permission struct {
	TeamID         uuid.UUID      `json:"teamId"`
	RepositoryID   int            `json:"repositoryId"`
	AccountID      int            `json:"accountId"`
	Privilege      PermissionType `json:"privilege"`
	CreatorID      int            `json:"creatorId"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      *time.Time     `json:"updatedAt,omitempty"`
	RepositoryName string         `json:"repositoryName"`
}

// ListPermissionsResponse represents the response from listing repository permissions.
type ListPermissionsResponse struct {
	Permissions []Permission `json:"permissions"`
	Size        int          `json:"size"`
	Page        int          `json:"page"`
	Count       int          `json:"count"`
}
