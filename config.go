package up

import (
	"net/http"
	"net/url"
	"time"

	"github.com/crossplane/crossplane-runtime/pkg/logging"
)

const (
	defaultBaseURL     = "https://api.upbound.io"
	defaultUserAgent   = "up-sdk-go"
	defaultHTTPTimeout = 10 * time.Second
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

// NewConfig build a new Config for communicating with the Upbound API.
func NewConfig(modifiers ...ConfigModifierFn) *Config {
	b, _ := url.Parse(defaultBaseURL)
	c := &Config{
		Client: &HTTPClient{
			BaseURL:      b,
			ErrorHandler: &DefaultErrorHandler{},
			HTTP: &http.Client{
				Timeout: defaultHTTPTimeout,
			},
			UserAgent: defaultUserAgent,
		},
		Logger: logging.NewNopLogger(),
	}
	for _, m := range modifiers {
		m(c)
	}
	return c
}
