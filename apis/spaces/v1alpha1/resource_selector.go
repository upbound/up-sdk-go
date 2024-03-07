// Copyright 2024 Upbound Inc.
// All rights reserved

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// +kubebuilder:object:generate=false
type SharedObject interface {
	client.Object
	ControlPlaneSelector() func(obj client.Object) (bool, error)
}

// +kubebuilder:object:generate=false
type Object interface {
	client.Object
	ControlPlaneSelector() func(obj client.Object) (bool, error)
}

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
	// no names in the list is a match
	m := len(r.Names) == 0
	for _, n := range r.Names {
		if o.GetName() == n {
			m = true
			break
		}
	}
	// if there is no match on names, return early,
	// no point to check labelSelectors at all.
	if !m {
		return m, nil
	}
	// check if any label selector matches
	for i := range r.LabelSelectors {
		labelSelector, err := metav1.LabelSelectorAsSelector(&r.LabelSelectors[i])
		if err != nil {
			return false, err
		}
		if m = m && labelSelector.Matches(o); m {
			// match, return early
			break
		}
	}
	return m, nil
}
