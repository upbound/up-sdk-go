# Upbound Go SDK

`up-sdk-go` is the official [Go] SDK for [Upbound]. It is currently under active
development and subject to breaking changes.

## Supported Services

The following services are currently supported:
- Accounts
- Control Planes
- Tokens

## Authentication

`up-sdk-go` currently defers authentication to the consumer by passing a
configured `http.Client`. The [_examples] directory contains examples of how
this can be accomplished with a `cookiejar` implementation and session tokens.

<!-- Named Links -->
[Go]: https://golang.org/
[Upbound]: https://cloud.upbound.io/
[_examples]: ./_examples
