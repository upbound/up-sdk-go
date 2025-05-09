// Copyright 2025 Upbound Inc
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

package users

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"

	"github.com/crossplane/crossplane-runtime/pkg/errors"

	"github.com/upbound/up-sdk-go"
	"github.com/upbound/up-sdk-go/fake"
	"github.com/upbound/up-sdk-go/service/common"
	"github.com/upbound/up-sdk-go/service/tokens"
)

func TestListTokens(t *testing.T) {
	errBoom := errors.New("boom")
	userID := uint(42)

	cases := map[string]struct {
		reason string
		cfg    *up.Config
		want   *tokens.TokensResponse
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
					MockNewRequest: fake.NewMockNewRequestFn(&http.Request{}, nil),
					MockDo:         fake.NewMockDoFn(errBoom),
				},
			},
			err: errBoom,
		},
		"Successful": {
			reason: "A successful request should return the token list as dataset-compatible values.",
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: fake.NewMockNewRequestFn(&http.Request{}, nil),
					MockDo: func(_ *http.Request, v interface{}) error {
						if tr, ok := v.(*tokens.TokensResponse); ok {
							*tr = tokens.TokensResponse{
								DataSet: []common.DataSet{
									{
										ID: uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"),
										AttributeSet: map[string]any{
											"name": "test-token",
										},
										Meta: map[string]any{
											"createdAt": time.Now().Add(-1 * time.Hour).Format(time.RFC3339),
										},
									},
								},
							}
						}
						return nil
					},
				},
			},
			want: &tokens.TokensResponse{
				DataSet: []common.DataSet{
					{
						ID: uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"),
						AttributeSet: map[string]any{
							"name": "test-token",
						},
						Meta: map[string]any{
							"createdAt": time.Now().Add(-1 * time.Hour).Format(time.RFC3339),
						},
					},
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			c := NewClient(tc.cfg)
			got, err := c.ListTokens(context.Background(), userID)

			if diff := cmp.Diff(tc.err, err, cmpopts.EquateErrors()); diff != "" {
				t.Errorf("\n%s\nListTokens(...): -want error, +got error:\n%s", tc.reason, diff)
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("\n%s\nListTokens(...): -want, +got:\n%s", tc.reason, diff)
			}
		})
	}
}
