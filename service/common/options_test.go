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

package common

import (
	"context"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestWithPage(t *testing.T) {
	o := WithPage(50)
	r, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://localhost:8080", nil)
	o(r)
	if diff := cmp.Diff("50", r.URL.Query().Get(PageParam)); diff != "" {
		t.Errorf("\n%s\nWithPage(...): -want, +got:\n%s", "should set page query parameter correctly", diff)
	}
}

func TestWithSize(t *testing.T) {
	o := WithSize(30)
	r, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "https://localhost:8080", nil)
	o(r)
	if diff := cmp.Diff("30", r.URL.Query().Get(SizeParam)); diff != "" {
		t.Errorf("\n%s\nWithSize(...): -want, +got:\n%s", "should set size query parameter correctly", diff)
	}
}
