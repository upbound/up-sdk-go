package namespaces

import (
	"context"
	"net/http"
	"path"

	"github.com/upbound/up-sdk-go"
	"github.com/upbound/up-sdk-go/service/controlplanes"
)

const (
	basePath          = "v1/namespaces"
	controlPlanesPath = "controlPlanes"
)

// Client is a namespaces client.
type Client struct {
	*up.Config
}

// NewClient builds a namespaces client from the passed config.
func NewClient(cfg *up.Config) *Client {
	return &Client{
		cfg,
	}
}

// Get a namespace on Upbound Cloud.
func (c *Client) Get(ctx context.Context, name string) (*NamespaceResponse, error) {
	req, err := c.Client.NewRequest(ctx, http.MethodGet, basePath, name, nil)
	if err != nil {
		return nil, err
	}
	ns := &NamespaceResponse{}
	err = c.Client.Do(req, &ns)
	if err != nil {
		return nil, err
	}
	return ns, nil
}

// List all namespaces for the authenticated user on Upbound Cloud.
func (c *Client) List(ctx context.Context) ([]NamespaceResponse, error) {
	req, err := c.Client.NewRequest(ctx, http.MethodGet, basePath, "", nil)
	if err != nil {
		return nil, err
	}
	ns := []NamespaceResponse{}
	err = c.Client.Do(req, &ns)
	if err != nil {
		return nil, err
	}
	return ns, nil
}

// ListControlPlanes lists all control planes in the given namespace on Upbound Cloud.
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
