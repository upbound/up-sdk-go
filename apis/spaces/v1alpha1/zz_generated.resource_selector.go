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

// Generated from spaces/v1beta1/resource_selector.go by ../hack/duplicate_api_type.sh. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ResourceSelector defines the selector for resource matching.
// An object is going to be matched if any of the provided label selectors
// matches object's labels AND any of provided names are equal to the object name.
type ResourceSelector struct {

	// A resource is matched if any of the label selector matches.
	// In case when the list is empty, resource is matched too.
	// +optional
	LabelSelectors []metav1.LabelSelector `json:"labelSelectors,omitempty"`

	// A resource is selected if its metadata.name matches any of the provided names.
	// In case when the list is empty, resource is matched too.
	// +optional
	// +listType=set
	Names []string `json:"names,omitempty"`
}

// Matchable is a resource that is potentially matchable by a resource selector
// +kubebuilder:object:generate=false
type Matchable interface {
	labels.Labels
	// GetName return the resource name
	GetName() string
}

var _ Matchable = &matchableObject{}

type matchableObject struct {
	obj client.Object
}

func (s *matchableObject) Has(label string) (exists bool) {
	_, has := s.obj.GetLabels()[label]
	return has
}

func (s *matchableObject) Get(label string) (value string) {
	return s.obj.GetLabels()[label]
}

func (s *matchableObject) GetName() string {
	return s.obj.GetName()
}

// Matches returns true if the provided object is matched by the selector
func (r *ResourceSelector) Matches(obj client.Object) (bool, error) { //nolint:gocyclo
	o := &matchableObject{obj: obj}

	// check if any name matches
	namesMatch := len(r.Names) == 0 // no names in the list is a match
	for _, n := range r.Names {
		if o.GetName() == n {
			namesMatch = true
			break
		}
	}
	if !namesMatch {
		return false, nil
	}

	// check if any label selector matches
	if len(r.LabelSelectors) == 0 {
		return true, nil
	}
	for i := range r.LabelSelectors {
		labelSelector, err := metav1.LabelSelectorAsSelector(&r.LabelSelectors[i])
		if err != nil {
			return false, err
		}
		if labelSelector.Matches(o) {
			return true, nil
		}
	}

	return false, nil
}
