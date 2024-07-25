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
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"path"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/upbound/up-sdk-go"
	upboundv1alpha1 "github.com/upbound/up-sdk-go/apis/upbound/v1alpha1"
)

const (
	basePath  = "apis/upbound.io/v1alpha1/namespaces"
	spacePath = "spaces"
)

// NewClient creates a new spaces client.
func NewClient(cfg *up.Config) *Client {
	return &Client{cfg.Client.With(func(c *up.HTTPClient) {
		c.ErrorHandler = &kubeErrorHandler{}
	})}
}

// Client is a spaces client.
type Client struct {
	uc up.Client
}

// Create creates a space.
func (c *Client) Create(ctx context.Context, namespace string, space *upboundv1alpha1.Space, opts *metav1.CreateOptions) (*upboundv1alpha1.Space, error) {
	var params url.Values
	if opts != nil {
		p, err := parameterCodec.EncodeParameters(opts, metav1.SchemeGroupVersion)
		if err != nil {
			return nil, err
		}
		params = p
	}
	urlPath := path.Join(basePath, namespace, spacePath)
	if len(params) > 0 {
		urlPath += "?" + params.Encode()
	}
	req, err := c.uc.NewRequest(ctx, http.MethodPost, "", urlPath, space)
	if err != nil {
		return nil, err
	}
	res := &upboundv1alpha1.Space{}
	err = c.uc.Do(req, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// List lists all accessible space.
func (c *Client) List(ctx context.Context, namespace string, opts *metav1.ListOptions) (*upboundv1alpha1.SpaceList, error) {
	var params url.Values
	if opts != nil {
		p, err := parameterCodec.EncodeParameters(opts, metav1.SchemeGroupVersion)
		if err != nil {
			return nil, err
		}
		params = p
	}
	urlPath := path.Join(basePath, namespace, spacePath)
	if len(params) > 0 {
		urlPath += "?" + params.Encode()
	}
	req, err := c.uc.NewRequest(ctx, http.MethodGet, "", urlPath, nil)
	if err != nil {
		return nil, err
	}
	res := &upboundv1alpha1.SpaceList{}
	err = c.uc.Do(req, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Delete deletes a space.
func (c *Client) Delete(ctx context.Context, namespace, name string, opts *metav1.DeleteOptions) error {
	var params url.Values
	if opts != nil {
		p, err := parameterCodec.EncodeParameters(opts, metav1.SchemeGroupVersion)
		if err != nil {
			return err
		}
		params = p
	}
	urlPath := path.Join(basePath, namespace, spacePath, name)
	if len(params) > 0 {
		urlPath += "?" + params.Encode()
	}
	req, err := c.uc.NewRequest(ctx, http.MethodDelete, "", urlPath, nil)
	if err != nil {
		return err
	}
	return c.uc.Do(req, nil)
}

var _ up.ResponseErrorHandler = (*kubeErrorHandler)(nil)

type kubeErrorHandler struct{}

// Handle implements up.ResponseErrorHandler and doesn't intercept errors.
func (k *kubeErrorHandler) Handle(res *http.Response) error {
	if res.StatusCode >= http.StatusOK && res.StatusCode < http.StatusMultipleChoices {
		return nil
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	out, _, err := codec.Decode(b, nil, nil)
	if err != nil {
		return err
	}
	// any status besides StatusSuccess is considered an error.
	if st, ok := out.(*metav1.Status); ok && st.Status != metav1.StatusSuccess {
		return errors.FromObject(st)
	}
	// reset response body for response reader.
	res.Body = io.NopCloser(bytes.NewBuffer(b))
	return nil
}
