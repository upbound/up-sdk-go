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
	"fmt"

	"github.com/theory/jsonpath"
	"github.com/theory/jsonpath/spec"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

// ValidateStructuralPath validates that the given JSONPath is a valid structural
// path, that is an RFC 9535 JSONPath expression with only name,
// non-negative index, and wildcard selector segments. Examples: .spec.foo,
// .spec.foo[2], .spec.foo[*], .spec.foo[*].bar.
func ValidateStructuralPath(pth *field.Path, s string) error {
	if s == "" {
		return field.Invalid(pth, s, "must not be empty")
	}
	if s == "." {
		return nil
	}
	if s[0] == '$' {
		return field.Invalid(pth, s, "must not start with '$'")
	}
	jp, err := jsonpath.Parse("$" + s)
	if err != nil {
		return field.Invalid(pth, jp, fmt.Sprintf("must be a valid JSONPath expression: %s", err))
	}
	for _, seg := range jp.Query().Segments() {
		if seg.IsDescendant() {
			return field.Invalid(pth, jp.String(), fmt.Sprintf("must not contain descendant selectors, found: %s", seg.String()))
		}
		for _, sel := range seg.Selectors() {
			switch sel := sel.(type) {
			case spec.Name, spec.WildcardSelector: // good
			case spec.Index:
				if sel < 0 {
					return field.Invalid(pth, jp.String(), fmt.Sprintf("must have non-negative index, found: %d", sel))
				}
			default:
				return field.Invalid(pth, jp.String(), fmt.Sprintf("must be a name, non-negative index, or wildcard selector, found unexpected: %s", sel.String()))
			}
		}
	}
	return nil
}

// ValidateNormalPath validates that the given JSONPath is a valid normal
// path according to RFC 9535. Examples: .spec.foo, .spec.foo[2].
func ValidateNormalPath(pth *field.Path, s string) error {
	if s == "" {
		return field.Invalid(pth, s, "must not be empty")
	}
	if s == "." {
		return nil
	}
	if s[0] == '$' {
		return field.Invalid(pth, s, "must not start with '$'")
	}
	jp, err := jsonpath.Parse("$" + s)
	if err != nil {
		return field.Invalid(pth, jp, fmt.Sprintf("must be a valid normal JSONPath expression: %s", err))
	}
	for _, seg := range jp.Query().Segments() {
		if seg.IsDescendant() {
			return field.Invalid(pth, jp.String(), fmt.Sprintf("must not contain descendant selectors, found: %s", seg.String()))
		}
		for _, sel := range seg.Selectors() {
			switch sel := sel.(type) {
			case spec.Name:
			case spec.Index:
				if sel < 0 {
					return field.Invalid(pth, jp.String(), fmt.Sprintf("must have non-negative index, found: %d", sel))
				}
			default:
				return field.Invalid(pth, jp.String(), fmt.Sprintf("must be a name or non-negative index, found unexpected: %s", sel.String()))
			}
		}
	}
	return nil
}

// PathLastNameSelector returns the last name selector in the given
// JSONPath if the last selector is of type name.
func PathLastNameSelector(s string) string {
	if s == "" || s == "." {
		return ""
	}
	jp, err := jsonpath.Parse("$" + s)
	if err != nil {
		return ""
	}
	segs := jp.Query().Segments()
	if len(segs) == 0 {
		return ""
	}
	sels := segs[len(segs)-1].Selectors()
	if len(sels) == 0 {
		return ""
	}
	switch sel := sels[len(sels)-1].(type) {
	case spec.Name:
		return string(sel)
	default:
		return ""
	}
}
