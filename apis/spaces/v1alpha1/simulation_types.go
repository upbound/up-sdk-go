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
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="SOURCE",type="string",JSONPath=".spec.controlPlaneName"
// +kubebuilder:printcolumn:name="SIMULATED",type="string",JSONPath=".status.simulatedControlPlaneName"
// +kubebuilder:printcolumn:name="ACCEPTING-CHANGES",type="string",JSONPath=".status.conditions[?(@.type=='AcceptingChanges')].status"
// +kubebuilder:printcolumn:name="STATE",type="string",JSONPath=".status.conditions[?(@.type=='AcceptingChanges')].reason"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,categories=spaces

// A Simulation creates a simulation of a source ControlPlane. You can apply a
// change set to the simulated control plane. When the Simulation is complete it
// will detect the changes and report the difference compared to the source
// control plane.
type Simulation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SimulationSpec   `json:"spec"`
	Status SimulationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SimulationList contains a list of Simulations.
type SimulationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Simulation `json:"items"`
}

// SimulationSpec specifies how to run the simulation.
type SimulationSpec struct {
	// ControlPlaneName is the name of the ControlPlane to simulate a change to.
	// This control plane is known as the Simulation's 'source' control plane.
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="The source controlplane can't be changed"
	ControlPlaneName string `json:"controlPlaneName"`

	// DesiredState of the simulation.
	// +kubebuilder:default=AcceptingChanges
	// +kubebuilder:validation:Enum=AcceptingChanges;Complete;Terminated
	// +kubebuilder:validation:XValidation:rule="oldSelf != 'Complete' || self == 'Complete' || self == 'Terminated'",message="A complete Simulation can only be terminated"
	// +kubebuilder:validation:XValidation:rule="oldSelf != 'Terminated' || self == oldSelf",message="A Simulation can't be un-terminated"
	DesiredState SimulationState `json:"desiredState"`

	// CompletionCriteria specify how Spaces should determine when the
	// simulation is complete. If any of the criteria are met, Spaces will set
	// the Simulation's desired state to complete. Omit the criteria if you want
	// to manually mark the Simulation complete.
	// +optional
	CompletionCriteria []CompletionCriterion `json:"completionCriteria,omitempty"`
}

// SimulationState represents the lifecyle state of a simulation.
type SimulationState string

const (
	// SimulationStateUnknown indicates the simulation's state is unknown.
	SimulationStateUnknown SimulationState = "Unknown"

	// SimulationStateStarting indicates the simulation is starting. It's not
	// yet ready to accept changes.
	SimulationStateStarting SimulationState = "Starting"

	// SimulationStateAcceptingChanges indicates the simulation is accepting
	// changes.
	SimulationStateAcceptingChanges SimulationState = "AcceptingChanges"

	// SimulationStateComplete indicates the simulation is complete. Changes
	// made once a simulation is complete won't appear in the change summary.
	// It's still possible to connect to and explore a complete simulation.
	SimulationStateComplete SimulationState = "Complete"

	// SimulationStateTerminated indicates the simulation has terminated. You
	// can explore the change summary, but the simulation is no longer
	// accessible.
	SimulationStateTerminated SimulationState = "Terminated"
)

// A CompletionCriterion specifies when a simulation is complete.
type CompletionCriterion struct {
	// Type of criterion.
	// +kubebuilder:validation:Enum=Duration
	Type CompletionCriterionType `json:"type"`

	// TODO(negz): Make Duration an optional field once we support other types
	// of completion criteria.

	// Duration after which the simulation is complete.
	Duration metav1.Duration `json:"duration"`
}

// A CompletionCriterionType is a type of criterion.
type CompletionCriterionType string

const (
	// CompletionCriterionTypeDuration specifies that a simulation is complete
	// after a specified duration. The duration is relative to when the
	// simulation entered the AcceptingChanges state.
	CompletionCriterionTypeDuration CompletionCriterionType = "Duration"
)

// SimulationStatus represents the observed state of a Simulation.
type SimulationStatus struct {
	xpv1.ResourceStatus `json:",inline"`

	// SimulatedControlPlaneName is the name of the control plane used to run
	// the simulation.
	// +kubebuilder:validation:MinLength=1
	// +optional
	SimulatedControlPlaneName *string `json:"simulatedControlPlaneName,omitempty"`

	// ControlPlaneData exported from the source control plane and imported to
	// the simulated control plane.
	// +optional
	ControlPlaneData *ControlPlaneData `json:"controlPlaneData,omitempty"`

	// Changes detected by the simulation. Only changes that happen while the
	// simulation is in the AcceptingChanges state are included.
	// +optional
	Changes []SimulationChange `json:"changes,omitempty"`
}

// ControlPlaneData exported from the source control plane and imported to the
// simulated control plane.
type ControlPlaneData struct {
	// ExportTimestamp is the time at which the source control plane's resources
	// were exported. Resources are exported to temporary storage before they're
	// imported to the simulated control plane.
	ExportTimestamp *metav1.Time `json:"exportTimestamp,omitempty"`

	// ImportTiemstamp is the time at which the source control plane's resources
	// were imported to the simulated control plane.
	ImportTimestamp *metav1.Time `json:"importTimestamp,omitempty"`
}

// A SimulationChange represents an object that changed while the simulation was
// in the AcceptingChanges state.
type SimulationChange struct {
	// Change type.
	// +kubebuilder:validation:Enum=Unknown;Create;Update;Delete
	Change SimulationChangeType `json:"change"`

	// ObjectReference to the changed object.
	ObjectReference ChangedObjectReference `json:"objectRef"`
}

// A SimulationChangeType represents the type of a change.
type SimulationChangeType string

// Simulation change types.
const (
	SimulationChangeTypeUnknown SimulationChangeType = "Unknown"
	SimulationChangeTypeCreate  SimulationChangeType = "Create"
	SimulationChangeTypeUpdate  SimulationChangeType = "Update"
	SimulationChangeTypeDelete  SimulationChangeType = "Delete"
)

// A ChangedObjectReference represents a changed object.
type ChangedObjectReference struct {
	// APIVersion of the changed resource.
	APIVersion string `json:"apiVersion"`

	// Kind of the changed resource.
	Kind string `json:"kind"`

	// Name of the changed resource.
	Name string `json:"name"`

	// Namespace of the changed resource.
	// +optional
	Namespace *string `json:"namespace,omitempty"`
}

// Simulation type metadata.
var (
	SimulationKind             = reflect.TypeOf(Simulation{}).Name()
	SimulationGroupKind        = schema.GroupKind{Group: Group, Kind: SimulationKind}.String()
	SimulationKindAPIVersion   = SimulationKind + "." + SchemeGroupVersion.String()
	SimulationGroupVersionKind = SchemeGroupVersion.WithKind(SimulationKind)
)

func init() {
	SchemeBuilder.Register(&Simulation{}, &SimulationList{})
}
