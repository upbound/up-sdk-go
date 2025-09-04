# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is the official Go SDK for Upbound, providing programmatic access to Upbound's cloud services. The SDK is organized around multiple services including accounts, configurations, control planes, organizations, repositories, robots, tokens, and user management.

## Development Commands

### Building and Testing
- `make build` - Build source code and artifacts for host platform
- `make build.all` - Build for all platforms
- `make test` - Run unit tests
- `make e2e` - Run end-to-end integration tests
- `make lint` - Run linting with golangci-lint (very comprehensive configuration)
- `make generate` - Run code generation for all modules
- `make reviewable` - Validate PR readiness (combines lint, test, generate)
- `make check-diff` - Ensure reviewable doesn't create git diff

### API/Types Generation
- `make go.generate.apis` - Generate API types and client code in the apis/ module
- The project uses separate Go modules for `/apis` and root

### Code Quality
- Uses golangci-lint with extensive rule set (see .golangci.yaml)
- Enforces crossplane-runtime error patterns - use `github.com/crossplane/crossplane-runtime/v2/pkg/errors` instead of standard `errors` package
- Prohibits certain test libraries in favor of Go standard testing

## Architecture

### Multi-Module Structure
- **Root module** (`github.com/upbound/up-sdk-go`): Core HTTP client and service implementations
- **APIs module** (`github.com/upbound/up-sdk-go/apis`): Kubernetes API types and generated client code

### Core Components
- **HTTP Client** (`client.go`): Configurable HTTP client with error handling, request/response processing, and context propagation via `ContextTransport`
- **Service Layer** (`service/`): Individual service clients for each Upbound service
  - Each service follows pattern: `service/{name}/client.go` with types in `types.go`
  - Services: accounts, auth, configurations, controlplanes, gitsources, organizations, repositories, robots, spaces, teams, tokens, userinfo, users
- **Common Utilities** (`service/common/`): Shared functionality like pagination options
- **Error Handling** (`errors/`): Structured error types following Upbound API patterns
- **HTTP Utilities** (`http/`): Headers and request ID propagation
- **Mock/Fake** (`fake/`): Generated mocks for testing

### Authentication
- Authentication is deferred to consumers via configured `http.Client`
- Examples in `_examples/` show cookiejar and session token patterns
- Request ID propagation for tracing across service calls

### API Types (apis/ module)
- Contains Kubernetes-style API definitions
- Uses controller-runtime and crossplane patterns
- Generates deepcopy methods and client code
- Separate module to avoid dependency conflicts

## Key Patterns

### Service Client Pattern
Each service follows this structure:
```go
type Client struct {
    *up.Config  // Embedded config with HTTP client
}

func NewClient(cfg *up.Config) *Client
func (c *Client) Get(ctx context.Context, name string) (*Response, error)
```

### Error Handling
- All errors should use crossplane-runtime error wrapping
- HTTP errors are parsed into structured `uerrors.Error` types
- Request IDs are propagated for debugging

### Testing
- Use Go standard testing, avoid external assertion libraries
- Tests in `*_test.go` files alongside implementation
- Mock clients available in `fake/` directory

## Development Workflow

1. Make changes to code
2. Run `make generate` if API types changed
3. Run `make reviewable` to validate changes
4. Ensure `make check-diff` passes (no uncommitted generated code)

## Important Notes

- Go 1.24.6+ required
- Project uses build/ submodule for common Makefiles - run `make submodules` if missing
- The `apis/` module has its own go.mod and dependency management
- Linting is very strict - see .golangci.yaml for full configuration
- New-only linting enabled (only lint changed code from HEAD)