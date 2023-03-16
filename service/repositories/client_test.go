// Copyright 2022 Upbound Inc
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

package repositories

import (
	"context"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/pkg/errors"

	"github.com/upbound/up-sdk-go"
	"github.com/upbound/up-sdk-go/fake"
	"github.com/upbound/up-sdk-go/service/common"
)

func TestCreate(t *testing.T) {
	errBoom := errors.New("boom")
	testAccount := "testA"
	testRepo := "testR"

	type args struct {
		account string
		name    string
	}
	cases := map[string]struct {
		reason string
		args   args
		cfg    *up.Config
		err    error
	}{
		"NewRequestFailed": {
			reason: "Failing to construct a request should return an error.",
			args: args{
				account: testAccount,
				name:    testRepo,
			},
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: func(ctx context.Context, method, prefix, urlPath string, body interface{}) (*http.Request, error) {
						if method != http.MethodPut {
							t.Errorf("unexpected method: %s", method)
						}
						if prefix != basePath {
							t.Errorf("unexpected prefix: %s", method)
						}
						if urlPath != path.Join(testAccount, testRepo) {
							t.Errorf("unexpected account: %s", urlPath)
						}
						if body != struct{}{} {
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
			args: args{
				account: testAccount,
				name:    testRepo,
			},
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: func(ctx context.Context, method, prefix, urlPath string, body interface{}) (*http.Request, error) {
						if method != http.MethodPut {
							t.Errorf("unexpected method: %s", method)
						}
						if prefix != basePath {
							t.Errorf("unexpected prefix: %s", method)
						}
						if urlPath != path.Join(testAccount, testRepo) {
							t.Errorf("unexpected account: %s", urlPath)
						}
						if body != struct{}{} {
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
			args: args{
				account: testAccount,
				name:    testRepo,
			},
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: func(ctx context.Context, method, prefix, urlPath string, body interface{}) (*http.Request, error) {
						if method != http.MethodPut {
							t.Errorf("unexpected method: %s", method)
						}
						if prefix != basePath {
							t.Errorf("unexpected prefix: %s", method)
						}
						if urlPath != path.Join(testAccount, testRepo) {
							t.Errorf("unexpected account: %s", urlPath)
						}
						if body != struct{}{} {
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
			err := c.CreateOrUpdate(context.Background(), tc.args.account, tc.args.name)
			if diff := cmp.Diff(tc.err, err, cmpopts.EquateErrors()); diff != "" {
				t.Errorf("\n%s\nCreateOrUpdate(...): -want error, +got error:\n%s", tc.reason, diff)
			}
		})
	}
}

func TestGet(t *testing.T) {
	errBoom := errors.New("boom")
	testAccount := "testA"
	testRepo := "testR"

	type args struct {
		account string
		name    string
	}
	cases := map[string]struct {
		reason string
		args   args
		cfg    *up.Config
		want   *RepositoryResponse
		err    error
	}{
		"NewRequestFailed": {
			reason: "Failing to construct a request should return an error.",
			args: args{
				account: testAccount,
				name:    testRepo,
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
						if urlPath != path.Join(testAccount, testRepo) {
							t.Errorf("unexpected account: %s", urlPath)
						}
						return nil, errBoom
					},
				},
			},
			err: errBoom,
		},
		"DoFailed": {
			reason: "Failing to execute request should return an error.",
			args: args{
				account: testAccount,
				name:    testRepo,
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
						if urlPath != path.Join(testAccount, testRepo) {
							t.Errorf("unexpected account: %s", urlPath)
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
			args: args{
				account: testAccount,
				name:    testRepo,
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
						if urlPath != path.Join(testAccount, testRepo) {
							t.Errorf("unexpected account: %s", urlPath)
						}
						return nil, nil
					},
					MockDo: fake.NewMockDoFn(nil),
				},
			},
			want: &RepositoryResponse{},
		},
	}
	for name, tc := range cases {
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

func TestList(t *testing.T) {
	errBoom := errors.New("boom")
	testAccount := "upbound"
	testURL, _ := url.Parse("https://localhost:8080")

	type args struct {
		account string
		opts    []common.ListOption
	}
	cases := map[string]struct {
		reason string
		args   args
		cfg    *up.Config
		want   *RepositoryListResponse
		err    error
	}{
		"NewRequestFailed": {
			reason: "Failing to construct a request should return an error.",
			args: args{
				account: testAccount,
			},
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: func(_ context.Context, method, prefix, urlPath string, body interface{}) (*http.Request, error) {
						if method != http.MethodGet {
							t.Errorf("unexpected method: %s", method)
						}
						if prefix != basePath {
							t.Errorf("unexpected prefix: %s", prefix)
						}
						if urlPath != testAccount {
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
			args: args{
				account: testAccount,
				opts:    []common.ListOption{common.WithSize(50)},
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
						if urlPath != testAccount {
							t.Errorf("unexpected account: %s", urlPath)
						}
						if body != nil {
							t.Errorf("unexpected body: %v", body)
						}
						r, _ := http.NewRequestWithContext(ctx, http.MethodGet, testURL.String(), nil)
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
				account: testAccount,
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
						if urlPath != testAccount {
							t.Errorf("unexpected account: %s", urlPath)
						}
						if body != nil {
							t.Errorf("unexpected body: %v", body)
						}
						r, _ := http.NewRequestWithContext(ctx, http.MethodGet, testURL.String(), nil)
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
			want: &RepositoryListResponse{},
		},
		"SuccessfulWithSize": {
			reason: "A successful request with size should not return an error.",
			args: args{
				account: testAccount,
				opts:    []common.ListOption{common.WithSize(50)},
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
						if urlPath != testAccount {
							t.Errorf("unexpected account: %s", urlPath)
						}
						if body != nil {
							t.Errorf("unexpected body: %v", body)
						}
						r, _ := http.NewRequestWithContext(ctx, http.MethodGet, testURL.String(), nil)
						return r, nil
					},
					MockDo: func(req *http.Request, _ interface{}) error {
						if req.URL.Host != testURL.Host {
							t.Errorf("unexpected host: %s", req.URL.Host)
						}
						if req.URL.Query().Get(common.SizeParam) != strconv.FormatInt(50, 10) {
							t.Errorf("unexpected size: %s", req.URL.Query().Get(common.SizeParam))
						}
						return nil
					},
				},
			},
			want: &RepositoryListResponse{},
		},
		"SuccessfulWithPage": {
			reason: "A successful request with page should not return an error.",
			args: args{
				account: testAccount,
				opts:    []common.ListOption{common.WithPage(50)},
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
						if urlPath != testAccount {
							t.Errorf("unexpected account: %s", urlPath)
						}
						if body != nil {
							t.Errorf("unexpected body: %v", body)
						}
						r, _ := http.NewRequestWithContext(ctx, http.MethodGet, testURL.String(), nil)
						return r, nil
					},
					MockDo: func(req *http.Request, _ interface{}) error {
						if req.URL.Host != testURL.Host {
							t.Errorf("unexpected host: %s", req.URL.Host)
						}
						if req.URL.Query().Get(common.PageParam) != strconv.FormatInt(50, 10) {
							t.Errorf("unexpected page: %s", req.URL.Query().Get(common.PageParam))
						}
						return nil
					},
				},
			},
			want: &RepositoryListResponse{},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			c := NewClient(tc.cfg)
			res, err := c.List(context.Background(), tc.args.account, tc.args.opts...)
			if diff := cmp.Diff(tc.err, err, cmpopts.EquateErrors()); diff != "" {
				t.Errorf("\n%s\nList(...): -want error, +got error:\n%s", tc.reason, diff)
			}
			if diff := cmp.Diff(tc.want, res); diff != "" {
				t.Errorf("\n%s\nList(...): -want, +got:\n%s", tc.reason, diff)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	errBoom := errors.New("boom")
	testAccount := "testA"
	testRepo := "testR"

	type args struct {
		account string
		name    string
	}
	cases := map[string]struct {
		reason string
		args   args
		cfg    *up.Config
		err    error
	}{
		"NewRequestFailed": {
			reason: "Failing to construct a request should return an error.",
			args: args{
				account: testAccount,
				name:    testRepo,
			},
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: func(ctx context.Context, method, prefix, urlPath string, body interface{}) (*http.Request, error) {
						if method != http.MethodDelete {
							t.Errorf("unexpected method: %s", method)
						}
						if prefix != basePath {
							t.Errorf("unexpected prefix: %s", method)
						}
						if urlPath != path.Join(testAccount, testRepo) {
							t.Errorf("unexpected account: %s", urlPath)
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
			args: args{
				account: testAccount,
				name:    testRepo,
			},
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: func(ctx context.Context, method, prefix, urlPath string, body interface{}) (*http.Request, error) {
						if method != http.MethodDelete {
							t.Errorf("unexpected method: %s", method)
						}
						if prefix != basePath {
							t.Errorf("unexpected prefix: %s", method)
						}
						if urlPath != path.Join(testAccount, testRepo) {
							t.Errorf("unexpected account: %s", urlPath)
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
			args: args{
				account: testAccount,
				name:    testRepo,
			},
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: func(ctx context.Context, method, prefix, urlPath string, body interface{}) (*http.Request, error) {
						if method != http.MethodDelete {
							t.Errorf("unexpected method: %s", method)
						}
						if prefix != basePath {
							t.Errorf("unexpected prefix: %s", method)
						}
						if urlPath != path.Join(testAccount, testRepo) {
							t.Errorf("unexpected account: %s", urlPath)
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
			err := c.Delete(context.Background(), tc.args.account, tc.args.name)
			if diff := cmp.Diff(tc.err, err, cmpopts.EquateErrors()); diff != "" {
				t.Errorf("\n%s\nDelete(...): -want error, +got error:\n%s", tc.reason, diff)
			}
		})
	}
}
