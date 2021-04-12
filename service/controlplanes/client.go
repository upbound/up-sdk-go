package controlplanes

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/upbound/up-sdk-go"
)

const basePath = "v1/controlPlanes"

type Client struct {
	*up.Config
}

func NewControlPlanesClient(cfg *up.Config) *Client {
	return &Client{
		cfg,
	}
}

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

func (c *Client) Delete(ctx context.Context, id uuid.UUID) error {
	req, err := c.Client.NewRequest(ctx, http.MethodDelete, basePath, id.String(), nil)
	if err != nil {
		return err
	}
	return c.Client.Do(req, nil)
}
