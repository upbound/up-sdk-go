package accounts

import (
	"context"
	"net/http"
	"path"

	"github.com/upbound/up-sdk-go"
	"github.com/upbound/up-sdk-go/service/controlplanes"
)

const (
	basePath          = "v1/accounts"
	controlPlanesPath = "controlPlanes"
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

// Get a account on Upbound Cloud.
func (c *Client) Get(ctx context.Context, name string) (*AccountResponse, error) {
	req, err := c.Client.NewRequest(ctx, http.MethodGet, basePath, name, nil)
	if err != nil {
		return nil, err
	}
	ns := &AccountResponse{}
	err = c.Client.Do(req, &ns)
	if err != nil {
		return nil, err
	}
	return ns, nil
}

// List all accounts for the authenticated user on Upbound Cloud.
func (c *Client) List(ctx context.Context) ([]AccountResponse, error) {
	req, err := c.Client.NewRequest(ctx, http.MethodGet, basePath, "", nil)
	if err != nil {
		return nil, err
	}
	ns := []AccountResponse{}
	err = c.Client.Do(req, &ns)
	if err != nil {
		return nil, err
	}
	return ns, nil
}

// ListControlPlanes lists all control planes in the given account on Upbound Cloud.
func (c *Client) ListControlPlanes(ctx context.Context, name string) ([]controlplanes.ControlPlaneResponse, error) {
	req, err := c.Client.NewRequest(ctx, http.MethodGet, basePath, path.Join(name, controlPlanesPath), nil)
	if err != nil {
		return nil, err
	}
	cp := []controlplanes.ControlPlaneResponse{}
	err = c.Client.Do(req, &cp)
	if err != nil {
		return nil, err
	}
	return cp, nil
}
