// Copyright 2023 Upbound Inc.
// All rights reserved

package v1alpha1

import (
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestResourceSelector(t *testing.T) {
	type object struct {
		name   string
		labels map[string]string
	}
	tests := map[string]struct {
		reason   string
		obj      object
		selector ResourceSelector
		matched  bool
		wantErr  bool
	}{
		"ObjectNameMatched": {
			reason: "object is matched if its name is declared in the list of names",
			obj: object{
				name: "foo",
			},
			selector: ResourceSelector{
				Names: []string{"foo", "bar"},
			},
			matched: true,
		},
		"ObjectNameNotMatched": {
			reason: "object is not matched if its name is not declared in the list of names",
			obj: object{
				name: "foo",
			},
			selector: ResourceSelector{
				Names: []string{"foo2", "bar"},
			},
			matched: false,
		},
		"ObjectLabelsMatched": {
			reason: "object is matched if it has labels declared in the selector",
			obj: object{
				labels: map[string]string{
					"l1": "v1",
					"l2": "v2",
				},
			},
			selector: ResourceSelector{
				LabelSelectors: []metav1.LabelSelector{
					{
						MatchLabels: map[string]string{
							"l1": "v1",
							"l2": "v2",
						},
					},
					// or
					{
						MatchLabels: map[string]string{
							"l3": "v3",
							"l4": "v4",
						},
					},
				},
			},
			matched: true,
		},
		"ObjectLabelsExist": {
			reason: "object is matched if it has labels declared in the selector",
			obj: object{
				labels: map[string]string{
					"l1": "v1",
					"l2": "v2",
				},
			},
			selector: ResourceSelector{
				LabelSelectors: []metav1.LabelSelector{
					{
						MatchExpressions: []metav1.LabelSelectorRequirement{
							{
								Key:      "l1",
								Operator: metav1.LabelSelectorOpExists,
							},
						},
					},
				},
			},
			matched: true,
		},
		"ExpressionErr": {
			reason: "return error if mathing expression is invalid",
			obj: object{
				labels: map[string]string{
					"l1": "v1",
					"l2": "v2",
				},
			},
			selector: ResourceSelector{
				LabelSelectors: []metav1.LabelSelector{
					{
						MatchExpressions: []metav1.LabelSelectorRequirement{
							{
								Key:      "l1",
								Operator: "foo",
							},
						},
					},
				},
			},
			wantErr: true,
		},
		"ObjectLabelsNotMatched": {
			reason: "object is not matched if it does not have labels declared in the selector",
			obj: object{
				labels: map[string]string{
					"l5": "v5",
					"l6": "v6",
				},
			},
			selector: ResourceSelector{
				LabelSelectors: []metav1.LabelSelector{
					{
						MatchLabels: map[string]string{
							"l1": "v1",
							"l2": "v2",
						},
					},
					// or
					{
						MatchLabels: map[string]string{
							"l3": "v3",
							"l4": "v4",
						},
					},
				},
			},
			matched: false,
		},
		"ObjectLabelsOrNameNotMatched": {
			reason: "object is not matched if neither its name nor labels matche the declared selector",
			obj: object{
				name: "foo",
				labels: map[string]string{
					"l5": "v5",
					"l6": "v6",
				},
			},
			selector: ResourceSelector{
				Names: []string{"bar"},
				// AND
				LabelSelectors: []metav1.LabelSelector{
					{
						MatchLabels: map[string]string{
							"l1": "v1",
							"l2": "v2",
						},
					},
					// or
					{
						MatchLabels: map[string]string{
							"l3": "v3",
							"l4": "v4",
						},
					},
				},
			},
			matched: false,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			obj := &unstructured.Unstructured{}
			obj.SetName(tc.obj.name)
			obj.SetLabels(tc.obj.labels)

			m, err := tc.selector.Matches(obj)
			if (err != nil) != tc.wantErr {
				t.Errorf("Matches() returns error unexpectedly: %v", tc.reason)
				return
			}
			if m != tc.matched {
				t.Errorf("Matches() got = %v, want %v: %v", m, tc.matched, tc.reason)
			}
		})
	}
}
