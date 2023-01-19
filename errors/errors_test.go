// Copyright 2021 Upbound Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package errors

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
)

func TestIsNotFound(t *testing.T) {
	cases := map[string]struct {
		reason string
		err    error
		want   bool
	}{
		"TestNotErrIsNotFound": {
			reason: "Error with type other than ErrorTypeNotFound should return false.",
			err:    &Error{Status: http.StatusInternalServerError, Title: http.StatusText(http.StatusInternalServerError)},
			want:   false,
		},
		"TestIsErrIsNotFound": {
			reason: "Error with type ErrorTypeNotFound should return true.",
			err: &Error{
				Status: http.StatusNotFound,
			},
			want: true,
		},
		"TestNotUpError": {
			reason: "Error that does not implement upError should return false.",
			err:    errors.New("other error"),
			want:   false,
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if diff := cmp.Diff(tc.want, IsNotFound(tc.err)); diff != "" {
				t.Errorf("\n%s\nIsNotFound(...): -want, +got:\n%s", tc.reason, diff)
			}
		})
	}
}
