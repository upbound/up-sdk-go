// Copyright 2021 Upbound Inc
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
