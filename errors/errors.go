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
)

var _ error = &Error{}

// ErrorType indicates the type of error.
type ErrorType string

// Upbound SDK error types.
const (
	ErrorTypeForbidden    ErrorType = "Forbidden"
	ErrorTypeNotFound     ErrorType = "NotFound"
	ErrorTypeUnauthorized ErrorType = "Unauthorized"
	ErrorTypeUnknown      ErrorType = "Unknown"
)

// Error is an Upbound SDK error.
type Error struct {
	Response ErrorResponse
	Type     ErrorType
	Message  string
}

// ErrorResponse is an Upbound SDK error response.
type ErrorResponse struct {
	Status int     `json:"status"`
	Title  string  `json:"title"`
	Detail *string `json:"detail,omitempty"`
}

// Error returns the error message with the underlying error.
func (e *ErrorResponse) Error() string {
	detail := ""
	if e.Detail != nil {
		detail = fmt.Sprintf(": %s", *e.Detail)
	}
	return fmt.Sprintf("%s%s", e.Title, detail)
}

// Error returns the error message with the underlying error.
func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Message, e.Response.Error())
}

// IsNotFound returns true if the error type is ErrorTypeNotFound, and false
// otherwise.
func (e *Error) IsNotFound() bool {
	return e.Type == ErrorTypeNotFound
}

// New constructs a new Upbound SDK error.
func New(err ErrorResponse, msg string, errType ErrorType) *Error {
	return &Error{
		Response: err,
		Type:     errType,
		Message:  msg,
	}
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
