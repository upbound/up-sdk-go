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
	MockWith       func(modifiers ...up.ClientModifierFn) up.Client
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

// With implements up.Client.
func (m *MockClient) With(modifiers ...up.ClientModifierFn) up.Client {
	return m.MockWith(modifiers...)
}
