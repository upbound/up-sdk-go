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
	"fmt"
	"net/http"
)

var _ error = &Error{}

// Error is an Upbound SDK error response.
type Error struct {
	Status int     `json:"status"`
	Title  string  `json:"title"`
	Detail *string `json:"detail,omitempty"`
}

// Error returns the error message with the underlying error.
func (e *Error) Error() string {
	detail := ""
	if e.Detail != nil {
		detail = fmt.Sprintf(": %s", *e.Detail)
	}
	return fmt.Sprintf("%s%s", e.Title, detail)
}

// IsNotFound returns true if the error type is ErrorTypeNotFound, and false
// otherwise.
func (e *Error) IsNotFound() bool {
	return e.Status == http.StatusNotFound
}

// IsNotFound returns true if the error is an Upbound SDK NotFound error, and
// false otherwise.
func IsNotFound(err error) bool {
	e, ok := err.(interface { // nolint:errorlint
		IsNotFound() bool
	})
	if !ok {
		return false
	}
	return e.IsNotFound()
}
