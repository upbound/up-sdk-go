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
	"net/http"
	"strconv"
)

const (
	// PageParam is the name of the page query parameter.
	PageParam = "page"
	// SizeParam is the name fo the size query parameter.
	SizeParam = "size"
)

// ListOption modifies a list request.
type ListOption func(*http.Request)

// WithSize sets the maximum number of entities included in the list response.
func WithSize(size int) ListOption {
	return func(r *http.Request) {
		q := r.URL.Query()
		q.Add(SizeParam, strconv.FormatInt(int64(size), 10))
		r.URL.RawQuery = q.Encode()
	}
}

// WithPage sets the page of entities included in the list response.
func WithPage(page int) ListOption {
	return func(r *http.Request) {
		q := r.URL.Query()
		q.Add(PageParam, strconv.FormatInt(int64(page), 10))
		r.URL.RawQuery = q.Encode()
	}
}
