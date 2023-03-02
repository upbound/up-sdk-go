// Copyright 2023 Upbound Inc
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

package configurations

import (
	"context"
	"net/http"
	"net/url"
	"path"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/pkg/errors"

	"github.com/upbound/up-sdk-go"
	"github.com/upbound/up-sdk-go/fake"
)

func TestList(t *testing.T) {
	errBoom := errors.New("boom")
	testURL, _ := url.Parse("https://localhost:8080")
	var account string = "someaccount"
	type args struct {
		account string
	}
	testCases := map[string]struct {
		reason string
		args   args
		cfg    *up.Config
		want   *ConfigurationListResponse
		err    error
	}{
		"NewRequestFailed": {
			reason: "Failing to construct a request should return an error.",
			args: args{
				account: account,
			},
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: fake.NewMockNewRequestFn(nil, errBoom),
				},
			},
			err: errBoom,
		},
		"DoFailed": {
			reason: "Failing to execute request should return an error.",
			args: args{
				account: account,
			},
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: func(ctx context.Context, method, prefix, urlPath string, body interface{}) (*http.Request, error) {
						if method != http.MethodGet {
							t.Errorf("unexpected method: %s", method)
						}
						if prefix != basePath {
							t.Errorf("unexpected prefix: %s", method)
						}
						if urlPath != account {
							t.Errorf("unexpected path: %s", urlPath)
						}
						r, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, testURL.String(), nil)
						return r, nil
					},
					MockDo: func(req *http.Request, _ interface{}) error {
						if req.URL.Host != testURL.Host {
							t.Errorf("unexpected host: %s", req.URL.Host)
						}
						return errBoom
					},
				},
			},
			err: errBoom,
		},
		"Successful": {
			reason: "A successful request should not return an error.",
			args: args{
				account: account,
			},
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: func(ctx context.Context, method, prefix, urlPath string, body interface{}) (*http.Request, error) {
						if method != http.MethodGet {
							t.Errorf("unexpected method: %s", method)
						}
						if prefix != basePath {
							t.Errorf("unexpected prefix: %s", method)
						}
						if urlPath != account {
							t.Errorf("unexpected path: %s", urlPath)
						}
						r, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, testURL.String(), nil)
						return r, nil
					},
					MockDo: func(req *http.Request, _ interface{}) error {
						if req.URL.Host != testURL.Host {
							t.Errorf("unexpected host: %s", req.URL.Host)
						}
						return nil
					},
				},
			},
			want: &ConfigurationListResponse{},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			c := NewClient(tc.cfg)
			res, err := c.List(context.Background(), tc.args.account)
			if diff := cmp.Diff(tc.err, err, cmpopts.EquateErrors()); diff != "" {
				t.Errorf("\n%s\nList(...): -want error, +got error:\n%s", tc.reason, diff)
			}
			if diff := cmp.Diff(tc.want, res); diff != "" {
				t.Errorf("\n%s\nList(...): -want, +got:\n%s", tc.reason, diff)
			}
		})
	}
}

func TestGet(t *testing.T) {
	errBoom := errors.New("boom")
	testURL, _ := url.Parse("https://localhost:8080")
	var account string = "someaccount"
	var name string = "someconfiguration"
	type args struct {
		account string
		name    string
	}
	testCases := map[string]struct {
		reason string
		args   args
		cfg    *up.Config
		want   *ConfigurationResponse
		err    error
	}{
		"NewRequestFailed": {
			reason: "Failing to construct a request should return an error.",
			args: args{
				account: account,
				name:    name,
			},
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: fake.NewMockNewRequestFn(nil, errBoom),
				},
			},
			err: errBoom,
		},
		"DoFailed": {
			reason: "Failing to execute request should return an error.",
			args: args{
				account: account,
				name:    name,
			},
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: func(ctx context.Context, method, prefix, urlPath string, body interface{}) (*http.Request, error) {
						if method != http.MethodGet {
							t.Errorf("unexpected method: %s", method)
						}
						if prefix != basePath {
							t.Errorf("unexpected prefix: %s", method)
						}
						if urlPath != path.Join(account, name) {
							t.Errorf("unexpected path: %s", urlPath)
						}
						r, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, testURL.String(), nil)
						return r, nil
					},
					MockDo: func(req *http.Request, _ interface{}) error {
						if req.URL.Host != testURL.Host {
							t.Errorf("unexpected host: %s", req.URL.Host)
						}
						return errBoom
					},
				},
			},
			err: errBoom,
		},
		"Successful": {
			reason: "A successful request should not return an error.",
			args: args{
				account: account,
				name:    name,
			},
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: func(ctx context.Context, method, prefix, urlPath string, body interface{}) (*http.Request, error) {
						if method != http.MethodGet {
							t.Errorf("unexpected method: %s", method)
						}
						if prefix != basePath {
							t.Errorf("unexpected prefix: %s", method)
						}
						if urlPath != path.Join(account, name) {
							t.Errorf("unexpected path: %s", urlPath)
						}
						r, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, testURL.String(), nil)
						return r, nil
					},
					MockDo: func(req *http.Request, _ interface{}) error {
						if req.URL.Host != testURL.Host {
							t.Errorf("unexpected host: %s", req.URL.Host)
						}
						return nil
					},
				},
			},
			want: &ConfigurationResponse{},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			c := NewClient(tc.cfg)
			res, err := c.Get(context.Background(), tc.args.account, tc.args.name)
			if diff := cmp.Diff(tc.err, err, cmpopts.EquateErrors()); diff != "" {
				t.Errorf("\n%s\nGet(...): -want error, +got error:\n%s", tc.reason, diff)
			}
			if diff := cmp.Diff(tc.want, res); diff != "" {
				t.Errorf("\n%s\nGet(...): -want, +got:\n%s", tc.reason, diff)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	errBoom := errors.New("boom")
	account := "upbound"
	testCases := map[string]struct {
		reason string
		cfg    *up.Config
		params *ConfigurationCreateParameters
		want   *ConfigurationResponse
		err    error
	}{
		"NewRequestFailed": {
			reason: "Failing to construct a request should return an error.",
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: fake.NewMockNewRequestFn(nil, errBoom),
				},
			},
			err: errBoom,
		},
		"DoFailed": {
			reason: "Failing to execute request should return an error.",
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: fake.NewMockNewRequestFn(nil, nil),
					MockDo:         fake.NewMockDoFn(errBoom),
				},
			},
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
			want: &ConfigurationResponse{},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			c := NewClient(tc.cfg)
			res, err := c.Create(context.Background(), account, tc.params)
			if diff := cmp.Diff(tc.err, err, cmpopts.EquateErrors()); diff != "" {
				t.Errorf("\n%s\nCreate(...): -want error, +got error:\n%s", tc.reason, diff)
			}
			if diff := cmp.Diff(tc.want, res); diff != "" {
				t.Errorf("\n%s\nCreate(...): -want, +got:\n%s", tc.reason, diff)
			}
		})
	}
}
