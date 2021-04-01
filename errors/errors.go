package errors

import (
	"fmt"
)

var _ error = &Error{}

// ErrorType indicates the type of error.
type ErrorType string

// Upbound SDK error types.
const (
	ErrorTypeNotFound ErrorType = "NotFound"
	ErrorTypeUnknown  ErrorType = "Unknown"
)

// Error is an Upbound SDK error.
type Error struct {
	Err     error
	Type    ErrorType
	Message string
}

// Error returns the error message with the underlying error.
func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Message, e.Err.Error())
}

// IsNotFound returns true if the error type is ErrorTypeNotFound, and false
// otherwise.
func (e *Error) IsNotFound() bool {
	return e.Type == ErrorTypeNotFound
}

// New constructs a new Upbound SDK error.
func New(err error, msg string, errType ErrorType) *Error {
	return &Error{
		Err:     err,
		Type:    errType,
		Message: msg,
	}
}

// IsNotFound returns true if the error is an Upbound SDK NotFound error, and
// false otherwise.
func IsNotFound(err error) bool {
	e, ok := err.(interface {
		IsNotFound() bool
	})
	if !ok {
		return false
	}
	return e.IsNotFound()
}
