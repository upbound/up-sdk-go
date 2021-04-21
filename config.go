package up

import (
	"github.com/crossplane/crossplane-runtime/pkg/logging"
)

// Config represents common configuration for Upbound SDK clients.
type Config struct {
	// Client is the underlying client.
	Client Client

	// Logger is the interface for structured logging.
	Logger logging.Logger
}

// A ConfigModifierFn modifies a Config.
type ConfigModifierFn func(*Config)

// NewConfig builds a new Config for communicating with the Upbound API.
func NewConfig(modifiers ...ConfigModifierFn) *Config {
	c := &Config{
		Client: NewClient(),
		Logger: logging.NewNopLogger(),
	}
	for _, m := range modifiers {
		m(c)
	}
	return c
}
