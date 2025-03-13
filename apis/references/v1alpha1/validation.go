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
	"strings"

	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/validation/field"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

var validGrants = sets.New(xpv1.ManagementActionObserve, xpv1.ManagementActionCreate, xpv1.ManagementActionUpdate, xpv1.ManagementActionDelete, xpv1.ManagementActionAll)

// ValidateObjectReference validates an ObjectReference.
func ValidateObjectReference(pth *field.Path, r *ObjectReference) []error {
	var errs []error

	if r.APIVersion == "" {
		errs = append(errs, field.Required(pth.Child("apiVersion"), ""))
	} else if n := len(strings.Split(r.APIVersion, "/")); n != 1 && n != 2 {
		errs = append(errs, field.Invalid(pth.Child("apiVersion"), r.APIVersion, "must be in the format 'group/version'"))
	}
	if r.Kind == "" {
		errs = append(errs, field.Required(pth.Child("kind"), ""))
	} else if strings.ToUpper(r.Kind[:1]) != r.Kind[:1] {
		errs = append(errs, field.Invalid(pth.Child("kind"), r.Kind, "must start with an uppercase letter"))
	}

	if r.Name == "" {
		errs = append(errs, field.Required(pth.Child("name"), ""))
	}

	for i, g := range r.Grants {
		if !validGrants.Has(g) {
			errs = append(errs, field.NotSupported(pth.Child("grants").Index(i), g, sets.List(validGrants)))
		}
	}

	return errs
}

// ValidateReferenceSchema validates a ReferenceSchema.
func ValidateReferenceSchema(pth *field.Path, s *ReferenceSchema) []error {
	var errs []error

	for i := range s.References {
		if err := ValidateReferencePath(pth.Child("references").Index(i), &s.References[i]); err != nil {
			errs = append(errs, err...)
		}
	}

	return errs
}

// ValidateReferencePath validates a ReferencePath.
func ValidateReferencePath(pth *field.Path, rp *ReferencePath) []error {
	var errs []error

	if rp.JSONPath == "" {
		errs = append(errs, field.Required(pth.Child("jsonPath"), ""))
	} else if err := ValidateStructuralPath(pth.Child("jsonPath"), rp.JSONPath); err != nil {
		errs = append(errs, err)
	}

	for i := range rp.Kinds {
		if err := ValidateReferenceKind(pth.Child("kinds").Index(i), &rp.Kinds[i]); err != nil {
			errs = append(errs, err...)
		}
	}

	return errs
}

// ValidateReferenceKind validates a ReferencableKind.
func ValidateReferenceKind(pth *field.Path, rk *ReferencableKind) []error {
	var errs []error

	if rk.APIVersion == "" {
		errs = append(errs, field.Required(pth.Child("apiVersion"), ""))
	} else if cs := strings.Split(rk.APIVersion, "/"); len(cs) != 2 && len(cs) != 1 {
		errs = append(errs, field.Invalid(pth.Child("apiVersion"), rk.APIVersion, "must be in the format 'version' or 'group/version'"))
	}
	if rk.Kind == "" {
		errs = append(errs, field.Required(pth.Child("kind"), ""))
	} else if strings.ToUpper(rk.Kind[:1]) != rk.Kind[:1] {
		errs = append(errs, field.Invalid(pth.Child("kind"), rk.Kind, "must start with an uppercase letter"))
	}

	return errs
}
