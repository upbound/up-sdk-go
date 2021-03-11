# SDK Package Structure

* Owner: Dan Mangum (@hasheddan)
* Reviewers: Upbound Go SDK Maintainers
* Status: Draft

## Motivation

The package structure of the Upbound Go SDK should facilitate the following
goals, in order:

1. Enable consumers to programmatically access public Upbound Cloud APIs with as
   minimal friction as possible.
2. Incur least amount of maintenance burden on SDK maintainers as possible.

The second goal should only be sacrificed in service of the first.

## Introduction

The design of the following components of an SDK are typically what either
enables or prohibits a good user experience:

- Consistent authentication patterns across services
- Clear versioning and release of services
- Intuitive and robust error handling
- Extensibility

Other critical features, such as performance, client-side rate-limiting, and
pagination, are important, but can typically be implemented regardless of SDK
package structure.

There are a few primary schools of thought as to how to implement a Go SDK.
These can be broadly categorized into:

- Single Client - Many Methods
- Single Client - Many Services
- Client per Service - Services Versioned Together
- Client per Service - Services Versioned Independently
- Mixed

### Single Client - Many Methods

Example: [Docker](https://github.com/moby/moby/tree/master/client)

This is the simplest and least scalable approach to build an SDK. It involves
typically having a single `Client` struct, which has exported methods that
indicate the service and operation being performed. For example, in the Docker
client, `func (c *Client)CreateContainer(...)` is the method for the container
"service" that creates a new container. Because this method is defined on the
single `Client` type, the configuration and functionality to actually perform
the HTTP request, handle errors, etc. can be shared between each of the methods.
This also necessitates that all "services" be versioned and released together,
due to the fact that they all share a common client.

Advantages:
- Simple implementation (maintainer +)
- Simple versioning / release (maintainer +)

Disadvantages:
- Tightly coupled versioning (consumer -)
- Code organization at scale (maintainer -)
- Method discoverability (consumer -)

Because our primary goal is a frictionless experience for SDK consumers, this is
clearly not a viable approach.

### Single Client - Many Services

Example: [go-github](https://github.com/google/go-github)

This is very similar to the previous strategy, but instead embeds service
structs in a common client, meaning that to access the methods of the service
you have to access the service first (e.g. `client.Organizations.GetMembers()`).
Methods still get all the benefits of a shared client, as it is piped down into
each of the service structs, but have clearer organization and discoverability.
Services may or may not be separated into their own packages.

Advantages:
- Simple implementation (maintainer +)
- Simple versioning / release (maintainer +)
- Code organization at scale (maintainer +)
- Method discoverability (consumer +)

Disadvantages:
- Tightly coupled versioning (consumer -)

This is a much more suitable approach, but still mostly favors maintainers,
which we have established as the secondary goal of the package structure.
However, if the number of services is relatively small and changes to existing
APIs are infrequent, tightly coupled versioning becomes much less of an issue.

### Client per Service - Services Versioned Together

Example: [stripe-go](https://github.com/stripe/stripe-go)

This strategy differs fairly significantly from the previous two as it offers a
different `Client` for each of the services. Frequently, the implementation of
each of these clients looks quite similar, and there is sometimes a practice of
having every `NewClient()` function accept a common `Config` argument from a
shared package, and a service specific `Options`. These options are used for
operations that are common across all clients, such as setting the underlying
transport and selecting authentication methods.

Advantages:
- Simple versioning / release (maintainer +)
- Code organization at scale (maintainer +)
- Method discoverability (consumer +)

Disadvantages:
- Code duplication (maintainer -)
- Potentially diverging client implementations between services (consumer -)
- Tightly coupled versioning (consumer -)

In this model, maintainers are paying the price of potentially many different
client implementations, without providing consumers the benefits of full
separation. However, separate client implementations does serve to facilitate
method discovery and customization that is specific to a given service, which
can reduce bloat of a single shared client.

### Client per Service - Services Versioned Independently

Example: [aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2)

This strategy is similar to the previous client per service approach, but goes
farther to break each service into its own independently versioned [Go
module](https://golang.org/ref/mod). This gives consumers the advantage of being
able to only import modules for services that they need, which can cut down on
dependencies. This is especially relevant when a dependency with an unacceptable
license is introduced, which [can
prohibit](https://github.com/getsentry/sentry-go/issues/297) an organization
from using an entire module. However, its primary benefit is allowing for
independently versioned releases for each service, which allows consumers to
pick when they want to upgrade each module, without forcing update of others. It
also enables a future in which service teams can own the versioning of their
given submodule.

Advantages:
- Code organization at scale (maintainer +)
- Method discoverability (consumer +)
- Independent versioning (consumer +)

Disadvantages:
- Code duplication (maintainer -)
- Complex versioning (maintainer -)
- Potentially diverging client implementations between services (consumer -)

This strategy notably has the most benefits for consumers, our stated primary
goal. In addition, since Go 1.13, supporting multiple submodules in a single
repository has become less of a chore, so the "complex versioning" disadvantage
is somewhat overstated.

> In short, maintaining multiple submodules primarily consists of using subpath
> git tags (e.g. `organizations/v0.1.0`) and using [local
> replace](https://golang.org/ref/mod#go-mod-file-replace) for shared packages.
> See the [aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) or
> [golang/tools](https://github.com/golang/tools) repositories for examples.

That being said, it does require considerable thought and effort compared to
other approaches. More information on this pattern can be found in the [Golang
FAQ](https://github.com/golang/go/wiki/Modules#faqs--multi-module-repositories).
Do note the following
[comment](https://github.com/golang/go/issues/26664#issuecomment-455232444) from
Russ Cox:

> For all but power users, you probably want to adopt the usual convention that
> one repo = one module. It's important for long-term evolution of code storage
> options that a repo can contain multiple modules, but it's almost certainly
> not something you want to do by default.

### Mixed

Example: [google-cloud-go](https://github.com/googleapis/google-cloud-go)

This strategy is mostly mentioned for completeness, and consists of a mixture of
the previous two strategies. It typically leads to confusion for consumers as to
which services have independent versioning and which do not. The advantages and
disadvantages are a union of the two previous sections.

## Consumption Model

Out of the strategies described above, the consumption model is primarily
impacted by the single client vs. multiple client distinction. In the single
client scenario, initialization only happens once, which can be useful when the
consumer is using multiple services.

```go
package main

import (
    "context"
    "fmt"
    "net/http"

    upsdk "github.com/upbound/up-sdk-go"
)

func main() {
    httpClient := &http.Client{}
    client := upsdk.NewClient(httpClient)
    orgs, _ := client.Organizations.List(ctx)
    fmt.Println(orgs)
    platforms, _ := client.Platforms.List(ctx)
    fmt.Println(platforms)
}
```

On the other hand, when there are multiple clients, the consumer must
instantiate each of them individually, but they may share underlying HTTP
transport and other configuration options (such as credentials source, etc.).

```go
package main

import (
    "context"
    "fmt"
    "net/http"

    upsdk "github.com/upbound/up-sdk-go"
   "github.com/upbound/up-sdk-go/organizations"
   "github.com/upbound/up-sdk-go/platforms"
)

func main() {
    client := &http.Client{}
    config := upsdk.NewConfig(client)
    orgClient := organizations.NewClient(config, organizations.Options{})
    orgs, _ := orgClient.ListOrganizations(ctx)
    fmt.Println(orgs)
    paltformsClient := platforms.NewClient(config, platforms.Options{})
    ps, _ := platformsClient.ListPlatforms(ctx)
    fmt.Println(ps)
}
```

When using multiple clients, they may or may not be versioned independently. In
the case that they are, each of the `up-sdk-go` imports above would reflect a
different module which could be obtained at an individual version. Do note that
the _range_ of acceptable versions between all of the imported modules would not
be unbounded as they share a common dependency on the core `up-sdk-go` module.

## Proposal

Strategy:
- Short Term: `Client per Service - Services Versioned Together`
- Long Term: `Client per Service - Services Versioned Independently`

Given the advantages and disadvantages outlined in the introduction, using the
independently versioned client per service approach provides the most
flexibility for future API changes while giving consumers the most control over
the APIs they are using. However, while the Upbound API is still maturing, the
delineation of services is not abundantly clear, so the overhead of maintaining
a multi-module repo that is likely to drop modules and add new ones on a
frequent basis is likely not worth the short-term hassle.

Instituting a pattern of multiple clients, but not breaking them out into their
own submodules until the Upbound API has reached stability allows us to invest
in patterns and tooling that scales, without incurring the cost of complex
versioning in the short term.

The closest alternative in terms of equivalent trade-offs would be the single
client with many services, but going that direction would lead to more
significant breaking changes in the future if we wanted to pursue an independent
versioning strategy.

### Error Handling

With the chosen approach, error handling can be handled at both the individual
service and shared `Config` level. This allows for common error handling such as
`NotFound` errors, while also supplementing with information specific to the
service being implemented.

### Extensibility

For an SDK, extensibility typically looks like consumers being able to supply
their own HTTP transport implementation, as well as any middleware, such as
logging, they would like to be injected. All of the strategies described in the
introduction facilitate this functionality, with the chosen approach doing so
with the common `Config` that are passed to all clients.
