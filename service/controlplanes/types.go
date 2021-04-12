package controlplanes

import (
	"time"

	"github.com/google/uuid"
)

// ControlPlaneStatus is the status of a control plane on Upbound Cloud.
type ControlPlaneStatus string

// A control plane will always be in one of the following phases.
const (
	ControlPlaneStatusProvisioning ControlPlaneStatus = "provisioning"
	ControlPlaneStatusUpdating     ControlPlaneStatus = "updating"
	ControlPlaneStatusReady        ControlPlaneStatus = "ready"
	ControlPlaneStatusDeleting     ControlPlaneStatus = "deleting"
)

// ControlPlane describes a control plane.
type ControlPlane struct {
	ID          uuid.UUID  `json:"id,omitempty"`
	Name        string     `json:"name,omitempty"`
	Description string     `json:"description,omitempty"`
	CreatorID   uint       `json:"creatorId,omitempty"`
	Reserved    bool       `json:"reserved"`
	SelfHosted  bool       `json:"selfHosted"`
	CreatedAt   *time.Time `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
	ExpiresAt   time.Time  `json:"expiresAt"`
	DeletedAt   *time.Time `json:"deletedAt,omitempty"`
}

// ControlPlanePermissionGroup describes control plane permissions for the
// authenticated user.
type ControlPlanePermissionGroup string

const (
	// ControlPlanePermissionMember has the ability to read the basic
	// environment of the team
	ControlPlanePermissionMember ControlPlanePermissionGroup = "member"
	// ControlPlanePermissionOwner has the ability to modify any object in a
	// linked control plane, including deleting the control plane.
	ControlPlanePermissionOwner ControlPlanePermissionGroup = "owner"
	// ControlPlanePermissionNone has no permissions on the control plane.
	ControlPlanePermissionNone ControlPlanePermissionGroup = "none"
)

// ControlPlaneResponse is the HTTP body returned by the Upbound API when
// fetching control planes.
type ControlPlaneResponse struct {
	ControlPlane           ControlPlane                `json:"controlPlane"`
	ControlPlaneStatus     ControlPlaneStatus          `json:"controlPlaneStatus,omitempty"`
	ControlPlanePermission ControlPlanePermissionGroup `json:"controlPlanePermission,omitempty"`
}

// ControlPlaneCreateParameters are the parameters for creating a control plane.
type ControlPlaneCreateParameters struct {
	Namespace     string `json:"namespace"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	SelfHosted    bool   `json:"selfHosted,omitempty"`
	KubeClusterID string `json:"kubeClusterID,omitempty"`
}
