package up

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/pkg/errors"

	uerrors "github.com/upbound/up-sdk-go/errors"
)

type Client interface {
	NewRequest(ctx context.Context, method, prefix, urlPath string, body interface{}) (*http.Request, error)
	Do(req *http.Request, obj interface{}) error
}

type HTTPClient struct {
	// BaseURL is the base Upbound API URL.
	BaseURL *url.URL

	// ErrorHandler controls how the client handles errors.
	ErrorHandler ResponseErrorHandler

	// HTTP is the underlying HTTP client.
	HTTP *http.Client

	// User agent for communicating with the Upbound API.
	UserAgent string
}

type ResponseErrorHandler interface {
	Handle(res *http.Response) error
}

// NewRequest builds an HTTP request.
func (c *HTTPClient) NewRequest(ctx context.Context, method, prefix, urlPath string, body interface{}) (*http.Request, error) {
	u, err := c.BaseURL.Parse(path.Join(prefix, urlPath))
	if err != nil {
		return nil, err
	}
	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequestWithContext(ctx, method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", c.UserAgent)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return req, nil
}

// Do performs an HTTP request and reads the body into the provided interface.
func (c *HTTPClient) Do(req *http.Request, obj interface{}) error {
	res, err := c.HTTP.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if err := c.handleErrors(res); err != nil {
		return err
	}
	if obj != nil {
		b, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return json.Unmarshal(b, &obj)
	}
	return nil
}

func (c *HTTPClient) handleErrors(res *http.Response) error {
	return c.ErrorHandler.Handle(res)
}

// DefaultErrorHandler is the default operations for handling errors returned by
// the Upbound API.
type DefaultErrorHandler struct{}

// Handle handles HTTP response errors from the Upbound API. Caller is
// responsible for closing response body.
func (h *DefaultErrorHandler) Handle(res *http.Response) error {
	status := res.StatusCode
	if status >= 200 && status < 300 {
		return nil
	}
	errBody := "could not read error body"
	b, err := io.ReadAll(res.Body)
	if err == nil {
		errBody = string(b)
	}
	switch status {
	case http.StatusNotFound:
		return uerrors.New(errors.New(errBody), "resource not found", uerrors.ErrorTypeNotFound)
	case http.StatusForbidden:
		return uerrors.New(errors.New(errBody), "forbidden", uerrors.ErrorTypeForbidden)
	case http.StatusUnauthorized:
		return uerrors.New(errors.New(errBody), "permission denied", uerrors.ErrorTypeUnauthorized)
	default:
		return uerrors.New(errors.New(errBody), "unknown error", uerrors.ErrorTypeUnknown)
	}
}
