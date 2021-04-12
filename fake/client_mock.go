package fake

import (
	"context"
	"net/http"

	"github.com/upbound/up-sdk-go"
)

var _ up.Client = &MockClient{}

// MockClient is a mock of an Upbound SDK Client.
type MockClient struct {
	MockNewRequest func(ctx context.Context, method, prefix, urlPath string, body interface{}) (*http.Request, error)
	MockDo         func(req *http.Request, obj interface{}) error
}

// NewMockNewRequestFn creates a new MockNewRequest function.
func NewMockNewRequestFn(req *http.Request, err error) func(ctx context.Context, method, prefix, urlPath string, body interface{}) (*http.Request, error) {
	return func(ctx context.Context, method, prefix, urlPath string, body interface{}) (*http.Request, error) {
		return req, err
	}
}

// NewMockDoFn creates a new MockDo function.
func NewMockDoFn(err error) func(req *http.Request, obj interface{}) error {
	return func(req *http.Request, obj interface{}) error {
		return err
	}
}

// NewRequest calls the underlying MockNewRequest function.
func (m *MockClient) NewRequest(ctx context.Context, method, prefix, urlPath string, body interface{}) (*http.Request, error) {
	return m.MockNewRequest(ctx, method, prefix, urlPath, body)
}

// Do calls the underlying MockDo function.
func (m *MockClient) Do(req *http.Request, obj interface{}) error {
	return m.MockDo(req, obj)
}
