// Copyright 2023 Upbound Inc.
// All rights reserved

// Package headers provides HTTP header name constants for the headers we care
// about.
package headers

// RequestIDHeader is used to uniquely identify a request as it tranverses
// through the system. See the envoy docs for more information.
// Ref: https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_conn_man/headers#x-request-id
const RequestIDHeader = "x-request-id"
