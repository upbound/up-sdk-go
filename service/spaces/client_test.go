// Copyright 2024 Upbound Inc
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

package spaces

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/crossplane/crossplane-runtime/pkg/test"

	"github.com/upbound/up-sdk-go"
)

func TestClient_Create(t *testing.T) {
	type args struct {
		namespace string
		space     *Space
		opts      *metav1.CreateOptions
	}
	type want struct {
		space *Space
		err   error
	}
	tests := map[string]struct {
		reason  string
		handler http.HandlerFunc
		args    args
		want    want
	}{
		"StatusError": {
			reason: "returns status error",
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
				writeObject(t, &metav1.Status{
					Status: metav1.StatusFailure,
					Code:   http.StatusInternalServerError,
					Reason: metav1.StatusReasonInternalError,
				}, w)
			},
			want: want{
				err: &errors.StatusError{
					ErrStatus: metav1.Status{
						TypeMeta: metav1.TypeMeta{
							Kind:       "Status",
							APIVersion: "v1",
						},
						Status: metav1.StatusFailure,
						Code:   http.StatusInternalServerError,
						Reason: metav1.StatusReasonInternalError,
					},
				},
			},
		},
		"Success": {
			reason: "returns the space",
			handler: func(w http.ResponseWriter, r *http.Request) {
				writeObject(t, &Space{
					ObjectMeta: metav1.ObjectMeta{
						Name: "space-aaaa",
					},
				}, w)
			},
			args: args{
				namespace: "test-org",
				space: &Space{
					ObjectMeta: metav1.ObjectMeta{
						GenerateName: "space-",
					},
				},
			},
			want: want{
				space: &Space{
					ObjectMeta: metav1.ObjectMeta{
						Name: "space-aaaa",
					},
				},
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			s := httptest.NewServer(tc.handler)
			defer s.Close()
			c := NewClient(up.NewConfig(func(cfg *up.Config) {
				cfg.Client = up.NewClient(func(u *up.HTTPClient) {
					u.BaseURL = parseURL(t, s.URL)
					u.HTTP = s.Client()
				})
			}))
			got, err := c.Create(ctx, tc.args.namespace, tc.args.space, tc.args.opts)
			if diff := cmp.Diff(tc.want.err, err, test.EquateErrors()); diff != "" {
				t.Errorf("\n%s\nCreate(...): -want error, +got error:\n%s", tc.reason, diff)
			}
			if diff := cmp.Diff(tc.want.space, got, cmpopts.IgnoreTypes(metav1.TypeMeta{})); diff != "" {
				t.Errorf("\n%s\nCreate(...): -want, +got:\n%s", tc.reason, diff)
			}
		})
	}
}

func TestClient_Delete(t *testing.T) {
	type args struct {
		namespace string
		name      string
		opts      *metav1.DeleteOptions
	}
	type want struct {
		err error
	}
	tests := map[string]struct {
		reason  string
		handler http.HandlerFunc
		args    args
		want    want
	}{
		"StatusError": {
			reason: "returns status error",
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
				writeObject(t, &metav1.Status{
					Status: metav1.StatusFailure,
					Code:   http.StatusInternalServerError,
					Reason: metav1.StatusReasonInternalError,
				}, w)
			},
			want: want{
				err: &errors.StatusError{
					ErrStatus: metav1.Status{
						TypeMeta: metav1.TypeMeta{
							Kind:       "Status",
							APIVersion: "v1",
						},
						Status: metav1.StatusFailure,
						Code:   http.StatusInternalServerError,
						Reason: metav1.StatusReasonInternalError,
					},
				},
			},
		},
		"Success": {
			reason: "returns no error",
			handler: func(w http.ResponseWriter, r *http.Request) {
				writeObject(t, &metav1.Status{
					Status: metav1.StatusSuccess,
					Code:   http.StatusOK,
				}, w)
			},
			args: args{
				namespace: "test-org",
				name:      "space-aaaa",
			},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			s := httptest.NewServer(tc.handler)
			defer s.Close()
			c := NewClient(up.NewConfig(func(cfg *up.Config) {
				cfg.Client = up.NewClient(func(u *up.HTTPClient) {
					u.BaseURL = parseURL(t, s.URL)
					u.HTTP = s.Client()
				})
			}))
			err := c.Delete(ctx, tc.args.namespace, tc.args.name, tc.args.opts)
			if diff := cmp.Diff(tc.want.err, err, test.EquateErrors()); diff != "" {
				t.Errorf("\n%s\nCreate(...): -want error, +got error:\n%s", tc.reason, diff)
			}
		})
	}
}

func writeObject(t *testing.T, obj runtime.Object, w io.Writer) {
	t.Helper()
	if err := codec.Encode(obj, w); err != nil {
		t.Fatal(err)
	}
}

func parseURL(t *testing.T, u string) *url.URL {
	t.Helper()
	url, err := url.Parse(u)
	if err != nil {
		t.Fatal(err)
	}
	return url
}
