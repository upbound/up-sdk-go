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
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
)

const (
	// DefaultEnvironmentName is the default name of an Environment.
	DefaultEnvironmentName = "default"
)

// Environment specifies where remote claims are scheduled to.
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,categories={scheduling,envs},shortName=env
type Environment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EnvironmentSpec   `json:"spec,omitempty"`
	Status EnvironmentStatus `json:"status,omitempty"`
}

// EnvironmentList contains a list of Environment.
//
// +kubebuilder:object:root=true
type EnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Environment `json:"items"`
}

// Objects return the list of items.
func (s *EnvironmentList) Objects() []client.Object {
	objs := make([]client.Object, len(s.Items))
	for i := range s.Items {
		objs[i] = &s.Items[i]
	}
	return objs
}

// EnvironmentSpec defines the desired state of Environment.
type EnvironmentSpec struct {
	// Dimensions are the label keys and values to select control planes
	// implementing the desired API resource group. These dimensions apply to
	// all resource groups uniformly. Per resource group dimensions can
	// be specified in the resource group configuration.
	//
	// +optional
	Dimensions Dimensions `json:"dimensions,omitempty"`

	// ResourceGroups define label keys and values and potentially other
	// properties to select a control plane along dimensions for individual
	// API resource groups.
	//
	// A resource group not specified in this list will not be scheduled.
	//
	// +optional
	// +listType=map
	// +listMapKey=name
	ResourceGroups []ResourceGroup `json:"resourceGroups,omitempty"`
	// Schedule defines the scheduling result of the resource group. It is
	// set by the scheduler and cannot be mutated by the user. But the user
	// can unset it to allow the scheduler to reschedule the resource group.
	//
	// If unset, the resource group will be rescheduled by first setting the
	// proposed schedule in the status and then setting this field here.
	//
	// If no scheduling is possible, the Ready condition will be set to False
	// with the NoControlPlaneAvailable reason.
	//
	// A schedule will never change unless the user unsets it.
	//
	// +optional
	// +listType=map
	// +listMapKey=name
	Schedule []ResourceSchedule `json:"schedule,omitempty"`
}

// ResourceGroup defines the desired scheduling of a resource group.
type ResourceGroup struct {
	// Name is the name of the resource group.
	Name string `json:"name"`

	// Dimensions are the label keys and values to select control planes
	// implementing the desired API resource group. These are in addition to
	// the dimensions specified at the .spec level and override them.
	Dimensions Dimensions `json:"dimensions,omitempty"`
}

// Dimensions defines label keys and values to select a control plane along
// dimensions.
type Dimensions map[string]string

// ResourceSchedule indicates where instances of a resource group are scheduled to,
// i.e. by which control plane and in which group and space claims will be
// implemented.
type ResourceSchedule struct {
	// Name is the name of the resource group.
	// +required
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name"`

	ResourceCoordinates `json:",inline"`
}

// ResourceCoordinates identity a control plane serving a resource group.
type ResourceCoordinates struct {
	// Space is the name of the space in which the resource group is scheduled.
	// This is empty if the resource group cannot be scheduled.
	// +optional
	Space string `json:"space"`

	// Group is the name of the group in which the resource group is scheduled.
	// This is empty if the resource group cannot be scheduled.
	// +optional
	Group string `json:"group"`

	// ControlPlane is the name of the control plane in which the resource group is scheduled.
	// This is empty if the resource group cannot be scheduled.
	// +optional
	ControlPlane string `json:"controlPlane"`
}

// EnvironmentStatus defines the observed state of the Environment.
type EnvironmentStatus struct {
	xpv1.ResourceStatus `json:",inline"`

	// ProposedSchedule indicates where instances of a resource group according
	// to the current dimensions are proposed to be scheduled to.
	ResourceGroups []ResourceStatus `json:"resourceGroups,omitempty"`
}

// ResourceStatus represents the status of one individual resource group.
type ResourceStatus struct {
	// Name is the name of the resource group.
	// +required
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name"`

	// Reason is a machine-readable reason for the scheduling decision.
	// +optional
	Reason string `json:"reason"`

	// Message is a human-readable message for the scheduling decision.
	// +optional
	Message string `json:"message"`

	// Proposed is the proposed schedule for the resource group. This either
	// matches the schedule in the spec or is a new proposal by the scheduler.
	//
	// In the later case, the user can accept the proposal by removing the
	// existing schedule in the spec for the given resource group.
	//
	// +optional
	Proposed *ResourceCoordinates `json:"proposed,omitempty"`
}

const (
	// ScheduleUpToDateType is a condition type indicating that the schedule
	// in spec.schedule is up to date, or whether the scheduler has proposed
	// a new schedule in status.proposedSchedule. The condition is True if
	// both match.
	ScheduleUpToDateType = "ScheduleUpToDate"

	// ScheduledReason indicates that the spec.schedule matches the
	// proposed schedules in status for every resource.
	ScheduledReason = "Scheduled"

	// SchedulePendingReason indicates that the spec.schedule does not match
	// the proposed schedules in status.
	SchedulePendingReason = "Pending"

	// ReadySchedulingFailedReason indicates that no control plane is
	// available for the given dimensions for at least one resource group.
	ReadySchedulingFailedReason = "SchedulingFailed"
)

// GetCondition of this Environment.
func (mg *Environment) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// SetConditions of this Environment.
func (mg *Environment) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// EnvironmentKind is the kind of the Environment.
var EnvironmentKind = reflect.TypeOf(Environment{}).Name()

func init() {
	SchemeBuilder.Register(&Environment{}, &EnvironmentList{})
}
