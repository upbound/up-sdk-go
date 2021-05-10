package controlplanes

import (
	"context"
	"net/http"
	"path"

	"github.com/google/uuid"

	"github.com/upbound/up-sdk-go"
)

const (
	basePath     = "v1/controlPlanes"
	viewOnlyPath = "viewOnly"
)

// Client is a control planes client.
type Client struct {
	*up.Config
}

// NewClient build a control planes client from the passed config.
func NewClient(cfg *up.Config) *Client {
	return &Client{
		cfg,
	}
}

// Create a control plane on Upbound Cloud.
func (c *Client) Create(ctx context.Context, params *ControlPlaneCreateParameters) (*ControlPlaneResponse, error) {
	req, err := c.Client.NewRequest(ctx, http.MethodPost, basePath, "", params)
	if err != nil {
		return nil, err
	}
	cp := &ControlPlaneResponse{}
	err = c.Client.Do(req, &cp)
	if err != nil {
		return nil, err
	}
	return cp, nil
}

// Get a control plane on Upbound Cloud.
func (c *Client) Get(ctx context.Context, id uuid.UUID) (*ControlPlaneResponse, error) { // nolint:interfacer
	req, err := c.Client.NewRequest(ctx, http.MethodGet, basePath, id.String(), nil)
	if err != nil {
		return nil, err
	}
	cp := &ControlPlaneResponse{}
	err = c.Client.Do(req, &cp)
	if err != nil {
		return nil, err
	}
	return cp, nil
}

// Delete a control plane on Upbound Cloud.
func (c *Client) Delete(ctx context.Context, id uuid.UUID) error { // nolint:interfacer
	req, err := c.Client.NewRequest(ctx, http.MethodDelete, basePath, id.String(), nil)
	if err != nil {
		return err
	}
	return c.Client.Do(req, nil)
}

// SetViewOnly sets the view-only value of a control plane on Upbound Cloud.
func (c *Client) SetViewOnly(ctx context.Context, id uuid.UUID, viewOnly bool) error { // nolint:interfacer
	req, err := c.Client.NewRequest(ctx, http.MethodPut, basePath, path.Join(id.String(), viewOnlyPath), &controlPlaneViewOnlyParameters{
		IsViewOnly: viewOnly,
	})
	if err != nil {
		return err
	}
	return c.Client.Do(req, nil)
}
