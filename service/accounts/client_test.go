package accounts

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/pkg/errors"

	"github.com/upbound/up-sdk-go"
	"github.com/upbound/up-sdk-go/fake"
	"github.com/upbound/up-sdk-go/service/controlplanes"
)

func TestGet(t *testing.T) {
	errBoom := errors.New("boom")

	cases := map[string]struct {
		reason string
		cfg    *up.Config
		name   string
		want   *AccountResponse
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
			want: &AccountResponse{},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			c := NewClient(tc.cfg)
			res, err := c.Get(context.Background(), tc.name)
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

	cases := map[string]struct {
		reason string
		cfg    *up.Config
		want   []AccountResponse
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
			want: []AccountResponse{},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			c := NewClient(tc.cfg)
			res, err := c.List(context.Background())
			if diff := cmp.Diff(tc.err, err, cmpopts.EquateErrors()); diff != "" {
				t.Errorf("\n%s\nList(...): -want error, +got error:\n%s", tc.reason, diff)
			}
			if diff := cmp.Diff(tc.want, res); diff != "" {
				t.Errorf("\n%s\nList(...): -want, +got:\n%s", tc.reason, diff)
			}
		})
	}
}

func TestListControlPlanes(t *testing.T) {
	errBoom := errors.New("boom")

	cases := map[string]struct {
		reason string
		cfg    *up.Config
		name   string
		want   []controlplanes.ControlPlaneResponse
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
			want: []controlplanes.ControlPlaneResponse{},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			c := NewClient(tc.cfg)
			res, err := c.ListControlPlanes(context.Background(), tc.name)
			if diff := cmp.Diff(tc.err, err, cmpopts.EquateErrors()); diff != "" {
				t.Errorf("\n%s\nListControlPlanes(...): -want error, +got error:\n%s", tc.reason, diff)
			}
			if diff := cmp.Diff(tc.want, res); diff != "" {
				t.Errorf("\n%s\nListControlPlanes(...): -want, +got:\n%s", tc.reason, diff)
			}
		})
	}
}
