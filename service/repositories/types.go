// Copyright 2022 Upbound Inc
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

package repositories

import (
	"time"
)

// RepositoryType is the base type for repository types
type RepositoryType string

// PublishingPolicy determines whether packages will be indexed
type PublishPolicy string

const (
	// RepositoryTypeProvider indicates that the repository contains a provider.
	RepositoryTypeProvider RepositoryType = "provider"
	// RepositoryTypeConfiguration indicates that the repository contains a
	// configuration.
	RepositoryTypeConfiguration RepositoryType = "configuration"
	// RepositoryTypeFunction indicates that the repository contains a
	// function.
	RepositoryTypeFunction RepositoryType = "function"
)

// Repository describes a repository.
type Repository struct {
	RepositoryID   uint            `json:"repositoryId"`
	AccountID      uint            `json:"accountId"`
	Name           string          `json:"name"`
	Type           *RepositoryType `json:"type,omitempty"`
	Public         bool            `json:"public"`
	Official       bool            `json:"official"`
	CurrentVersion *string         `json:"currentVersion,omitempty"`
	CreatedAt      time.Time       `json:"createdAt"`
	UpdatedAt      *time.Time      `json:"updatedAt,omitempty"`
	Publish        *PublishPolicy  `json:"publishPolicy"`
}

// PackageStatusType indicates the status of a package.
type PackageStatusType string

const (
	// PackageStatusReceived indicates a package was received and no further
	// action has been taken.
	PackageStatusReceived PackageStatusType = "received"
	// PackageStatusAnalyzing indicates a package is in the process of being
	// validated.
	PackageStatusAnalyzing PackageStatusType = "analyzing"
	// PackageStatusRejected indicates the package failed validation.
	PackageStatusRejected PackageStatusType = "rejected"
	// PackageStatusAccepted indicates the package has been validated.
	PackageStatusAccepted PackageStatusType = "accepted"
	// PackageStatusPublished indicates the package has been validated and is
	// now published in Upbound Marketplace.
	PackageStatusPublished PackageStatusType = "published"
)

// Package describes a package in a repository.
type Package struct {
	PackageID    uint              `json:"packageId"`
	RepositoryID uint              `json:"repositoryId"`
	Version      string            `json:"version"`
	Status       PackageStatusType `json:"status"`
	Digest       string            `json:"digest"`
	Reason       *string           `json:"reason,omitempty"`
	CreatedAt    time.Time         `json:"createdAt"`
	UpdatedAt    *time.Time        `json:"updatedAt,omitempty"`
}

// RepositoryResponse is the HTTP body returned when fetching a repository.
type RepositoryResponse struct {
	Repository
	Versions []Package `json:"versions"`
}

// RepositoryListResponse is the HTTP body returned when listing repositories.
type RepositoryListResponse struct {
	Repositories []Repository `json:"repositories"`
	Size         int          `json:"size"`
	Page         int          `json:"page"`
	Count        int          `json:"count"`
}

// RepositoryCreateOrUpdateRequest is the HTTP body for creating or updating a
// repository.
type RepositoryCreateOrUpdateRequest struct {
	Public  bool `json:"public"`
	Publish bool `json:"publish"`
}
