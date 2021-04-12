package fake

import (
	"context"
	"net/http"

	"github.com/upbound/up-sdk-go"
)

var _ up.Client = &MockClient{}

type MockClient struct {
	MockNewRequest func(ctx context.Context, method, prefix, urlPath string, body interface{}) (*http.Request, error)
	MockDo         func(req *http.Request, obj interface{}) error
}

func NewMockNewRequestFn(req *http.Request, err error) func(ctx context.Context, method, prefix, urlPath string, body interface{}) (*http.Request, error) {
	return func(ctx context.Context, method, prefix, urlPath string, body interface{}) (*http.Request, error) {
		return req, err
	}
}

func NewMockDoFn(err error) func(req *http.Request, obj interface{}) error {
	return func(req *http.Request, obj interface{}) error {
		return err
	}
}

func (m *MockClient) NewRequest(ctx context.Context, method, prefix, urlPath string, body interface{}) (*http.Request, error) {
	return m.MockNewRequest(ctx, method, prefix, urlPath, body)
}

func (m *MockClient) Do(req *http.Request, obj interface{}) error {
	return m.MockDo(req, obj)
}
