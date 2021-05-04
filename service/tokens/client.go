package tokens

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/upbound/up-sdk-go"
)

const (
	basePath = "v1/tokens"

	errMissingParams = "parameters must be supplied"
)

// Client is a accounts client.
type Client struct {
	*up.Config
}

// NewClient builds a accounts client from the passed config.
func NewClient(cfg *up.Config) *Client {
	return &Client{
		cfg,
	}
}

// Create a token on Upbound Cloud.
func (c *Client) Create(ctx context.Context, params *TokenCreateParameters) (*TokenResponse, error) {
	body := &tokenCreateRequest{
		Data: tokenCreateParameters{
			Type:                  tokenBody,
			TokenCreateParameters: params,
		},
	}
	req, err := c.Client.NewRequest(ctx, http.MethodPost, basePath, "", body)
	if err != nil {
		return nil, err
	}
	ns := &TokenResponse{}
	err = c.Client.Do(req, &ns)
	if err != nil {
		return nil, err
	}
	return ns, nil
}

// Get a token on Upbound Cloud.
func (c *Client) Get(ctx context.Context, id uuid.UUID) (*TokenResponse, error) { // nolint:interfacer
	req, err := c.Client.NewRequest(ctx, http.MethodGet, basePath, id.String(), nil)
	if err != nil {
		return nil, err
	}
	ns := &TokenResponse{}
	err = c.Client.Do(req, &ns)
	if err != nil {
		return nil, err
	}
	return ns, nil
}

// Update a token on Upbound Cloud.
func (c *Client) Update(ctx context.Context, params *TokenUpdateParameters) (*TokenResponse, error) {
	body := &tokenUpdateRequest{
		Data: tokenUpdateParameters{
			Type:                  tokenBody,
			TokenUpdateParameters: params,
		},
	}
	if params == nil {
		return nil, errors.New(errMissingParams)
	}
	req, err := c.Client.NewRequest(ctx, http.MethodPatch, basePath, params.ID.String(), body)
	if err != nil {
		return nil, err
	}
	ns := &TokenResponse{}
	err = c.Client.Do(req, &ns)
	if err != nil {
		return nil, err
	}
	return ns, nil
}

// Delete a token on Upbound Cloud.
func (c *Client) Delete(ctx context.Context, id uuid.UUID) error { // nolint:interfacer
	req, err := c.Client.NewRequest(ctx, http.MethodDelete, basePath, id.String(), nil)
	if err != nil {
		return err
	}
	return c.Client.Do(req, nil)
}
