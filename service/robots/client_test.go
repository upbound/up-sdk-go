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

package robots

import (
	"context"
	"net/http"
	"path"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/upbound/up-sdk-go"
	"github.com/upbound/up-sdk-go/fake"
	"github.com/upbound/up-sdk-go/service/tokens"
)

func TestCreate(t *testing.T) {
	errBoom := errors.New("boom")

	cases := map[string]struct {
		reason string
		cfg    *up.Config
		params *RobotCreateParameters
		want   *RobotResponse
		err    error
	}{
		"NewRequestFailed": {
			reason: "Failing to construct a request should return an error.",
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: func(_ context.Context, method, prefix, urlPath string, body interface{}) (*http.Request, error) {
						if method != http.MethodPost {
							t.Errorf("unexpected method: %s", method)
						}
						if prefix != basePath {
							t.Errorf("unexpected prefix: %s", prefix)
						}
						if urlPath != "" {
							t.Errorf("unexpected path: %s", urlPath)
						}
						if _, ok := body.(*robotCreateRequest); !ok {
							t.Errorf("unexpected body: %v", body)
						}
						return nil, errBoom
					},
				},
			},
			err: errBoom,
		},
		"DoFailed": {
			reason: "Failing to execute request should return an error.",
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: func(_ context.Context, method, prefix, urlPath string, body interface{}) (*http.Request, error) {
						if method != http.MethodPost {
							t.Errorf("unexpected method: %s", method)
						}
						if prefix != basePath {
							t.Errorf("unexpected prefix: %s", prefix)
						}
						if urlPath != "" {
							t.Errorf("unexpected path: %s", urlPath)
						}
						if _, ok := body.(*robotCreateRequest); !ok {
							t.Errorf("unexpected body: %v", body)
						}
						return nil, nil
					},
					MockDo: fake.NewMockDoFn(errBoom),
				},
			},
			err: errBoom,
		},
		"Successful": {
			reason: "A successful request should not return an error.",
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: func(_ context.Context, method, prefix, urlPath string, body interface{}) (*http.Request, error) {
						if method != http.MethodPost {
							t.Errorf("unexpected method: %s", method)
						}
						if prefix != basePath {
							t.Errorf("unexpected prefix: %s", prefix)
						}
						if urlPath != "" {
							t.Errorf("unexpected path: %s", urlPath)
						}
						if _, ok := body.(*robotCreateRequest); !ok {
							t.Errorf("unexpected body: %v", body)
						}
						return nil, nil
					},
					MockDo: fake.NewMockDoFn(nil),
				},
			},
			want: &RobotResponse{},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			c := NewClient(tc.cfg)
			res, err := c.Create(context.Background(), tc.params)
			if diff := cmp.Diff(tc.err, err, cmpopts.EquateErrors()); diff != "" {
				t.Errorf("\n%s\nCreate(...): -want error, +got error:\n%s", tc.reason, diff)
			}
			if diff := cmp.Diff(tc.want, res); diff != "" {
				t.Errorf("\n%s\nCreate(...): -want, +got:\n%s", tc.reason, diff)
			}
		})
	}
}

func TestGet(t *testing.T) {
	errBoom := errors.New("boom")
	uid := uuid.MustParse("4654b8b5-c01d-4fbe-8800-22c347c21383")

	cases := map[string]struct {
		reason string
		cfg    *up.Config
		id     uuid.UUID
		want   *RobotResponse
		err    error
	}{
		"NewRequestFailed": {
			reason: "Failing to construct a request should return an error.",
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: func(_ context.Context, method, prefix, urlPath string, body interface{}) (*http.Request, error) {
						if method != http.MethodGet {
							t.Errorf("unexpected method: %s", method)
						}
						if prefix != basePath {
							t.Errorf("unexpected prefix: %s", prefix)
						}
						if urlPath != uid.String() {
							t.Errorf("unexpected path: %s", urlPath)
						}
						if body != nil {
							t.Errorf("unexpected body: %v", body)
						}
						return nil, errBoom
					},
				},
			},
			id:  uid,
			err: errBoom,
		},
		"DoFailed": {
			reason: "Failing to execute request should return an error.",
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: func(_ context.Context, method, prefix, urlPath string, body interface{}) (*http.Request, error) {
						if method != http.MethodGet {
							t.Errorf("unexpected method: %s", method)
						}
						if prefix != basePath {
							t.Errorf("unexpected prefix: %s", prefix)
						}
						if urlPath != uid.String() {
							t.Errorf("unexpected path: %s", urlPath)
						}
						if body != nil {
							t.Errorf("unexpected body: %v", body)
						}
						return nil, nil
					},
					MockDo: fake.NewMockDoFn(errBoom),
				},
			},
			id:  uid,
			err: errBoom,
		},
		"Successful": {
			reason: "A successful request should not return an error.",
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: func(_ context.Context, method, prefix, urlPath string, body interface{}) (*http.Request, error) {
						if method != http.MethodGet {
							t.Errorf("unexpected method: %s", method)
						}
						if prefix != basePath {
							t.Errorf("unexpected prefix: %s", prefix)
						}
						if urlPath != uid.String() {
							t.Errorf("unexpected path: %s", urlPath)
						}
						if body != nil {
							t.Errorf("unexpected body: %v", body)
						}
						return nil, nil
					},
					MockDo: fake.NewMockDoFn(nil),
				},
			},
			id:   uid,
			want: &RobotResponse{},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			c := NewClient(tc.cfg)
			res, err := c.Get(context.Background(), tc.id)
			if diff := cmp.Diff(tc.err, err, cmpopts.EquateErrors()); diff != "" {
				t.Errorf("\n%s\nGet(...): -want error, +got error:\n%s", tc.reason, diff)
			}
			if diff := cmp.Diff(tc.want, res); diff != "" {
				t.Errorf("\n%s\nGet(...): -want, +got:\n%s", tc.reason, diff)
			}
		})
	}
}

func TestListTokens(t *testing.T) {
	errBoom := errors.New("boom")
	uid := uuid.MustParse("4654b8b5-c01d-4fbe-8800-22c347c21383")

	cases := map[string]struct {
		reason string
		cfg    *up.Config
		id     uuid.UUID
		want   *tokens.TokensResponse
		err    error
	}{
		"NewRequestFailed": {
			reason: "Failing to construct a request should return an error.",
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: func(_ context.Context, method, prefix, urlPath string, body interface{}) (*http.Request, error) {
						if method != http.MethodGet {
							t.Errorf("unexpected method: %s", method)
						}
						if prefix != basePath {
							t.Errorf("unexpected prefix: %s", prefix)
						}
						if urlPath != path.Join(uid.String(), tokensPath) {
							t.Errorf("unexpected path: %s", urlPath)
						}
						if body != nil {
							t.Errorf("unexpected body: %v", body)
						}
						return nil, errBoom
					},
				},
			},
			id:  uid,
			err: errBoom,
		},
		"DoFailed": {
			reason: "Failing to execute request should return an error.",
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: func(_ context.Context, method, prefix, urlPath string, body interface{}) (*http.Request, error) {
						if method != http.MethodGet {
							t.Errorf("unexpected method: %s", method)
						}
						if prefix != basePath {
							t.Errorf("unexpected prefix: %s", prefix)
						}
						if urlPath != path.Join(uid.String(), tokensPath) {
							t.Errorf("unexpected path: %s", urlPath)
						}
						if body != nil {
							t.Errorf("unexpected body: %v", body)
						}
						return nil, nil
					},
					MockDo: fake.NewMockDoFn(errBoom),
				},
			},
			id:  uid,
			err: errBoom,
		},
		"Successful": {
			reason: "A successful request should not return an error.",
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: fake.NewMockNewRequestFn(nil, nil),
					MockDo:         fake.NewMockDoFn(nil),
				},
			},
			id:   uid,
			want: &tokens.TokensResponse{},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			c := NewClient(tc.cfg)
			res, err := c.ListTokens(context.Background(), tc.id)
			if diff := cmp.Diff(tc.err, err, cmpopts.EquateErrors()); diff != "" {
				t.Errorf("\n%s\nListTokens(...): -want error, +got error:\n%s", tc.reason, diff)
			}
			if diff := cmp.Diff(tc.want, res); diff != "" {
				t.Errorf("\n%s\nListTokens(...): -want, +got:\n%s", tc.reason, diff)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	errBoom := errors.New("boom")
	uid := uuid.MustParse("4654b8b5-c01d-4fbe-8800-22c347c21383")

	cases := map[string]struct {
		reason string
		cfg    *up.Config
		err    error
	}{
		"NewRequestFailed": {
			reason: "Failing to construct a request should return an error.",
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: func(_ context.Context, method, prefix, urlPath string, body interface{}) (*http.Request, error) {
						if method != http.MethodDelete {
							t.Errorf("unexpected method: %s", method)
						}
						if prefix != basePath {
							t.Errorf("unexpected prefix: %s", prefix)
						}
						if urlPath != uid.String() {
							t.Errorf("unexpected path: %s", urlPath)
						}
						if body != nil {
							t.Errorf("unexpected body: %v", body)
						}
						return nil, errBoom
					},
				},
			},
			err: errBoom,
		},
		"DoFailed": {
			reason: "Failing to execute request should return an error.",
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: func(_ context.Context, method, prefix, urlPath string, body interface{}) (*http.Request, error) {
						if method != http.MethodDelete {
							t.Errorf("unexpected method: %s", method)
						}
						if prefix != basePath {
							t.Errorf("unexpected prefix: %s", prefix)
						}
						if urlPath != uid.String() {
							t.Errorf("unexpected path: %s", urlPath)
						}
						if body != nil {
							t.Errorf("unexpected body: %v", body)
						}
						return nil, nil
					},
					MockDo: fake.NewMockDoFn(errBoom),
				},
			},
			err: errBoom,
		},
		"Successful": {
			reason: "A successful request should not return an error.",
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: func(_ context.Context, method, prefix, urlPath string, body interface{}) (*http.Request, error) {
						if method != http.MethodDelete {
							t.Errorf("unexpected method: %s", method)
						}
						if prefix != basePath {
							t.Errorf("unexpected prefix: %s", prefix)
						}
						if urlPath != uid.String() {
							t.Errorf("unexpected path: %s", urlPath)
						}
						if body != nil {
							t.Errorf("unexpected body: %v", body)
						}
						return nil, nil
					},
					MockDo: fake.NewMockDoFn(nil),
				},
			},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			c := NewClient(tc.cfg)
			err := c.Delete(context.Background(), uid)
			if diff := cmp.Diff(tc.err, err, cmpopts.EquateErrors()); diff != "" {
				t.Errorf("\n%s\nDelete(...): -want error, +got error:\n%s", tc.reason, diff)
			}
		})
	}
}
