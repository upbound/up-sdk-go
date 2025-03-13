// Copyright 2025 Upbound Inc.
// All rights reserved

package v1alpha1

import (
	"fmt"
	"testing"

	gocmp "github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

func TestValidateStructuralPath(t *testing.T) {
	tests := map[string]struct {
		jsonPath string
		wantErr  string
	}{
		"empty": {
			jsonPath: "",
			wantErr:  `spec.jsonPath: Invalid value: "": must not be empty`,
		},
		"dot": {
			jsonPath: ".",
			wantErr:  "<nil>",
		},
		"dollar": {
			jsonPath: "$.spec.foo",
			wantErr:  `spec.jsonPath: Invalid value: "$.spec.foo": must not start with '$'`,
		},
		"simple": {
			jsonPath: ".spec.foo",
			wantErr:  "<nil>",
		},
		"index": {
			jsonPath: ".spec.foo[0]",
			wantErr:  "<nil>",
		},
		"wildcard": {
			jsonPath: ".spec.foo[*].bar",
			wantErr:  "<nil>",
		},
		"descendent": {
			jsonPath: ".spec..foo",
			wantErr:  `spec.jsonPath: Invalid value: "$[\"spec\"]..[\"foo\"]": must not contain descendant selectors, found: ..["foo"]`,
		},
		"filter": {
			jsonPath: ".spec.foo[?(@.bar)]",
			wantErr:  `spec.jsonPath: Invalid value: "$[\"spec\"][\"foo\"][?(@[\"bar\"])]": must be a name, non-negative index, or wildcard selector, found unexpected: ?(@["bar"])`,
		},
		"last": {
			jsonPath: ".spec.foo[-1]",
			wantErr:  `spec.jsonPath: Invalid value: "$[\"spec\"][\"foo\"][-1]": must have non-negative index, found: -1`,
		},
		"range": {
			jsonPath: ".spec.foo[1:2]",
			wantErr:  `spec.jsonPath: Invalid value: "$[\"spec\"][\"foo\"][1:2]": must be a name, non-negative index, or wildcard selector, found unexpected: 1:2`,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			err := ValidateStructuralPath(field.NewPath("spec", "jsonPath"), tt.jsonPath)
			if diff := gocmp.Diff(tt.wantErr, fmt.Sprintf("%v", err)); diff != "" {
				t.Errorf("ValidateStructuralPath() +got, -want\n%s", diff)
			}
		})
	}
}

func TestValidateNormalPath(t *testing.T) {
	tests := map[string]struct {
		jsonPath string
		wantErr  string
	}{
		"empty": {
			jsonPath: "",
			wantErr:  `spec.jsonPath: Invalid value: "": must not be empty`,
		},
		"dot": {
			jsonPath: ".",
			wantErr:  "<nil>",
		},
		"dollar": {
			jsonPath: "$.spec.foo",
			wantErr:  `spec.jsonPath: Invalid value: "$.spec.foo": must not start with '$'`,
		},
		"simple": {
			jsonPath: ".spec.foo",
			wantErr:  "<nil>",
		},
		"index": {
			jsonPath: ".spec.foo[0]",
			wantErr:  "<nil>",
		},
		"wildcard": {
			jsonPath: ".spec.foo[*].bar",
			wantErr:  `spec.jsonPath: Invalid value: "$[\"spec\"][\"foo\"][*][\"bar\"]": must be a name or non-negative index, found unexpected: *`,
		},
		"descendent": {
			jsonPath: ".spec..foo",
			wantErr:  `spec.jsonPath: Invalid value: "$[\"spec\"]..[\"foo\"]": must not contain descendant selectors, found: ..["foo"]`,
		},
		"filter": {
			jsonPath: ".spec.foo[?(@.bar)]",
			wantErr:  `spec.jsonPath: Invalid value: "$[\"spec\"][\"foo\"][?(@[\"bar\"])]": must be a name or non-negative index, found unexpected: ?(@["bar"])`,
		},
		"last": {
			jsonPath: ".spec.foo[-1]",
			wantErr:  `spec.jsonPath: Invalid value: "$[\"spec\"][\"foo\"][-1]": must have non-negative index, found: -1`,
		},
		"range": {
			jsonPath: ".spec.foo[1:2]",
			wantErr:  `spec.jsonPath: Invalid value: "$[\"spec\"][\"foo\"][1:2]": must be a name or non-negative index, found unexpected: 1:2`,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			err := ValidateNormalPath(field.NewPath("spec", "jsonPath"), tt.jsonPath)
			if diff := gocmp.Diff(tt.wantErr, fmt.Sprintf("%v", err)); diff != "" {
				t.Errorf("ValidateNormalPath() +got, -want\n%s", diff)
			}
		})
	}
}

func TestPathLastNameSelector(t *testing.T) {
	tests := map[string]struct {
		p    string
		want string
	}{
		"empty": {
			p: "",
		},
		"dot": {
			p: ".",
		},
		"simple": {
			p:    ".spec.foo",
			want: "foo",
		},
		"index": {
			p:    ".spec.foo[0].foo",
			want: "foo",
		},
		"wildcard": {
			p:    ".spec.foo[*].bar",
			want: "bar",
		},
		"descendent": {
			p:    ".spec..foo",
			want: "foo",
		},
		"filter": {
			p:    ".spec.foo[?(@.bar)].bar",
			want: "bar",
		},
		"last": {
			p:    ".spec.foo[-1].bar",
			want: "bar",
		},
		"range": {
			p:    ".spec.foo[1:2].bar",
			want: "bar",
		},
		"endingInIndex": {
			p: ".spec.foo.bar[0]",
		},
		"endingInWildcard": {
			p: ".spec.foo.bar[*]",
		},
		"endingInDescendent": {
			p:    ".spec.foo.bar..baz",
			want: "baz",
		},
		"endingInFilter": {
			p: ".spec.foo.bar[?(@.baz)]",
		},
		"endingInLast": {
			p: ".spec.foo.bar[-1]",
		},
		"endingInRange": {
			p: ".spec.foo.bar[1:2]",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := PathLastNameSelector(tt.p)
			if diff := gocmp.Diff(tt.want, got); diff != "" {
				t.Errorf("PathLastNameSelector() -want +got:\n%s", diff)
			}
		})
	}
}
