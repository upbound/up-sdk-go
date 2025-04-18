// Copyright 2023 Upbound Inc.
// All rights reserved

// Package request offers functions for working with requestIDs.
package request

import (
	"context"

	"github.com/google/uuid"
)

// WithID puts the request ID into the current context.
func WithID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, contextRequestIDKey, id)
}

// IDFromContext returns the request ID from the context.
// A zero ID is returned if there are no idenfiers in the
// current context.
func IDFromContext(ctx context.Context) string {
	v := ctx.Value(contextRequestIDKey)
	if v == nil {
		return ""
	}
	id, ok := v.(string)
	if !ok {
		panic("requestID is not of type string")
	}
	return id
}

// NewID generates a new UUID string.
func NewID() string {
	return uuid.NewString()
}

type contextRequestIDType struct{}

var contextRequestIDKey = &contextRequestIDType{} //nolint:gochecknoglobals //This is intended to be global.
