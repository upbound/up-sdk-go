// Copyright 2025 Upbound Inc
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
	"encoding/json"
	"fmt"

	"github.com/crossplane/crossplane-runtime/v2/pkg/resource/unstructured/composite"
)

const (
	// ControlPlaneDimensionLabelPrefix is the prefix for dimension labels, i.e.
	// of the shape dimension.scheduling.upbound.io/region: us-west-1. The part
	// after the prefix is the dimension name that is used during scheduling,
	// when selecting a control plane in an environment.
	ControlPlaneDimensionLabelPrefix = "dimension.scheduling.upbound.io/"

	// NamespaceResourceGroupLabel is the label applied to a control plane
	// group (namespace) to indicate that it serves the given resource group.
	// The value is an API group. Only control planes in groups (namespaces)
	// with this label will be considered for scheduling. This is a security
	// measure to prevent control planes from receiving claims just because the
	// attacker has access to the configuration OCI image and can create control
	// planes.
	// TODO(sttts): actually enforce this in the scheduler.
	NamespaceResourceGroupLabel = "scheduling.upbound.io/resource-group"

	// ControlPlanePortalLabel is the label applied to a control plane to show
	// up as portal in the console. The value must be "true".
	// TODO(tnthornton): assert this is only true via admission.
	ControlPlanePortalLabel = "upbound.io/portal"

	// CustomResourceDefinitionRemoteClaimLabel is the label applied to CRDs to
	// indicate that they are remote claims. This label is set when a RemoteConfiguration
	// creates a CRD.
	CustomResourceDefinitionRemoteClaimLabel = "internal.scheduling.upbound.io/remote"

	// CustomResourceDefinitionXRKindAnnotationKey is the annotation key used to
	// store the XR kind on the remote claim CRD. It is set by the RemoteConfigurationRevision
	// controller, and cunsumed by the bind controller to know the object type
	// to create.
	CustomResourceDefinitionXRKindAnnotationKey = "internal.scheduling.upbound.io/xr-kind"

	// RemoteClaimBindFinalizer is the finalizer used on remote claim CRs to
	// synchronize the life-cycle between the claim and the XR. It is first set
	// by the bind controller, but later relied on and removed on deletion by the
	// composite controller.
	RemoteClaimBindFinalizer = "bind.scheduling.upbound.io"

	// RemoteClaimEnvironmentAnnotationKey is the annotation key used to store
	// the environment on a remote claim CR. It can be set by the user on create,
	// and defaults to "default" if not set.
	RemoteClaimEnvironmentAnnotationKey = "scheduling.upbound.io/environment"

	// RemoteClaimExternalNameAnnotationKey is the annotation key used to store
	// the "external name" coordinates of the target control plane on a remote
	// claim CR. It holds schedulingv1alpha1.ResourceCoordinates as serialized
	// JSON. It is set by the bind controller following the scheduler's decision
	// in the Environment. Alternatively, it can be set by the user on create of
	// the remote claim CR for one-off scheduling.
	// TODO(sttts): come up with a scheduling authorization model. Now the user
	//              can set arbitrary control plane coordinates as external names.
	RemoteClaimExternalNameAnnotationKey = "scheduling.upbound.io/external-name"

	// RemoteClaimExternalNameLabelKey is the label key used to store the hash
	// of the external name of the target control plane. It is used to watch
	// matching remote claims from the composite controller, i.e. to reduce
	// resource consumption on the receiving service control plane.
	RemoteClaimExternalNameLabelKey = "internal.scheduling.upbound.io/external-name-hash"

	// RemoteClaimBoundAnnotationKey is the annotation key used to mark a remote
	// claim as successfully bound. It cannot be changed by the user. It is set
	// AFTER the XR has been created in the target control plane. This
	// annotation triggers the ownership of the remote claim CR to go over from
	// the bind controller to the composite controller.
	RemoteClaimBoundAnnotationKey = "internal.scheduling.upbound.io/bound"

	// CompositeUpstreamExternalNameAnnotationKey is the annotation key used to
	// store the ustream coordinates of the remote claim control plane that is
	// bound to the composite resource. It has the same format as
	// scheduling.upbound.io/external-name.
	CompositeUpstreamExternalNameAnnotationKey = "scheduling.upbound.io/upstream-external-name"

	// CompositeRemoteClaimFinalizer is the finalizer used on composites with
	// remote claim CRs.
	CompositeRemoteClaimFinalizer = "composite.scheduling.upbound.io"
)

// GetCompositeClaimCoordinates extracts the external name from a composite, or
// returns nil if the composite has no external name annotation.
func GetCompositeClaimCoordinates(xr *composite.Unstructured) (ResourceCoordinates, bool, error) {
	en := xr.GetAnnotations()[CompositeUpstreamExternalNameAnnotationKey]
	if en == "" {
		return ResourceCoordinates{}, false, nil
	}

	var ec ResourceCoordinates
	if err := json.Unmarshal([]byte(en), &ec); err != nil {
		return ResourceCoordinates{}, false, err
	}
	return ec, true, nil
}

// ExternalName returns the external name of resource coordinates as a string.
func ExternalName(coords *ResourceCoordinates) (string, error) {
	if coords == nil {
		return "", nil
	}

	bs, err := json.Marshal(coords)
	if err != nil {
		return "", fmt.Errorf("failed to marshal external name: %w", err)
	}
	return string(bs), nil
}
