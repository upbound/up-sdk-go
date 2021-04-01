package errors

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
)

func TestIsNotFound(t *testing.T) {
	cases := map[string]struct {
		reason string
		err    error
		want   bool
	}{
		"TestNotErrIsNotFound": {
			reason: "Error with type other than ErrorTypeNotFound should return false.",
			err: &Error{
				Type: ErrorTypeUnknown,
			},
			want: false,
		},
		"TestIsErrIsNotFound": {
			reason: "Error with type ErrorTypeNotFound should return true.",
			err: &Error{
				Type: ErrorTypeNotFound,
			},
			want: true,
		},
		"TestNotUpError": {
			reason: "Error that does not implement upError should return false.",
			err:    errors.New("other error"),
			want:   false,
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if diff := cmp.Diff(tc.want, IsNotFound(tc.err)); diff != "" {
				t.Errorf("\n%s\nIsNotFound(...): -want, +got:\n%s", tc.reason, diff)
			}
		})
	}
}
