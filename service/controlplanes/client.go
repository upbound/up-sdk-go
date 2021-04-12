package controlplanes

import (
	"context"
	"net/http"

	"github.com/google/uuid"

	"github.com/upbound/up-sdk-go"
)

const basePath = "v1/controlPlanes"

// Client is a control planes client.
type Client struct {
	*up.Config
}

// NewControlPlanesClient build a control planes client from the passed config.
func NewControlPlanesClient(cfg *up.Config) *Client {
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
func (c *Client) Get(ctx context.Context, id uuid.UUID) (*ControlPlaneResponse, error) {
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
func (c *Client) Delete(ctx context.Context, id uuid.UUID) error {
	req, err := c.Client.NewRequest(ctx, http.MethodDelete, basePath, id.String(), nil)
	if err != nil {
		return err
	}
	return c.Client.Do(req, nil)
}
