# !!DEPRECATED!! Please consider to use your language's native HTTP client to integrate with Variance

We are now encouraging our customers to use the native HTTP client to call our endpoints directly, given we only have
one endpoint and it is not worth having a SDK as an extra layer of indirection that could be a point of failure.


# Intrinsic Go Library

[![fern shield](https://img.shields.io/badge/%F0%9F%8C%BF-SDK%20generated%20by%20Fern-brightgreen)](https://github.com/fern-api/fern)
[![go shield](https://img.shields.io/badge/go-docs-blue)](https://pkg.go.dev/github.com/intrinsic-org/intrinsic-go)

The Intrinsic Go library provides convenient access to the Intrinsic API from Go.

## Requirements

This module requires Go version >= 1.13.

# Installation

Run the following command to use the Intrinsic Go library in your module:

```sh
go get github.com/intrinsic-org/intrinsic-go
```

## Usage

```go

package main

import (
	"context"
	"log"

	intrinsicgoclient "github.com/intrinsic-org/intrinsic-go/client"
	"github.com/intrinsic-org/intrinsic-go/option"
)

func main() {
	client := intrinsicgoclient.NewClient(
		option.WithAPIKey("<YOUR_API_KEY>"),
	)

	response, err := client.Events.CreateEventSync(
		context.TODO(),
		"new_text",
		map[string]interface{}{
			"caption": "text_content",
		},
	)

	if err != nil {
		log.Fatalf("Error creating event: %v", err)
	}

	// Assuming the 'response' variable contains useful information you'd like to log or use
	log.Printf("Created event : %+v", response)
}

```

## Timeouts

Setting a timeout for each individual request is as simple as using the standard
`context` library. Setting a one second timeout for an individual API call looks
like the following:

```go
ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
defer cancel()

response, err := client.EventTypes.CreateEventType(
  ctx,
  &intrinsic.CreateEventTypeRequest{
    Name: "example",
    Fields: []*intrinsic.EventTypeField{
      {
        FieldName: "field",
        Type:      intrinsic.EventTypeFieldTypeString,
      },
    },
  },
)
```

## Request Options

A variety of request options are included to adapt the behavior of the library, which includes
configuring authorization tokens, or providing your own instrumented `*http.Client`. Both of
these options are shown below:

```go
client := intrinsicgoclient.NewClient(
  option.WithAPIKey("<YOUR_API_KEY>"),
  option.WithHTTPClient(
    &http.Client{
      Timeout: 5 * time.Second,
    },
  ),
)
```

These request options can either be specified on the client so that they're applied on _every_
request (shown above), or for an individual request like so:

```go
response, err := client.Events.CreateEventSync(
		context.TODO(),
		"new_text",
		map[string]interface{}{
			"caption": "text_content",
		},
  option.WithAPIKey("<YOUR_API_KEY>"),
)
```

> Providing your own `*http.Client` is recommended. Otherwise, the `http.DefaultClient` will be used,
> and your client will wait indefinitely for a response (unless the per-request, context-based timeout
> is used).

## Automatic Retries

The Intrinsic Go client is instrumented with automatic retries with exponential backoff. A request will be
retried as long as the request is deemed retriable and the number of retry attempts has not grown larger
than the configured retry limit (default: 2).

A request is deemed retriable when any of the following HTTP status codes is returned:

- [408](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/408) (Timeout)
- [409](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/409) (Conflict)
- [429](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429) (Too Many Requests)
- [5XX](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500) (Internal Server Errors)

You can use the `option.WithMaxAttempts` option to configure the maximum retry limit to
your liking. For example, if you want to disable retries for the client entirely, you can
set this value to 1 like so:

```go
client := intrinsicgoclient.NewClient(
  option.WithMaxAttempts(1),
)
```

This can be done for an individual request, too:

```go
	response, err := client.Events.CreateEventSync(
		context.TODO(),
		"new_text",
		map[string]interface{}{
			"caption": "text_content",
		},
  option.WithMaxAttempts(1),
)
```

## Errors

Structured error types are returned from API calls that return non-success status codes. For example,
you can check if the error was due to a bad request (i.e. status code 400) with the following:

```go
response, err := client.Events.CreateEventSync(
  context.TODO(),
  "not_a_real_event_type",
  map[string]interface{}{
    "caption": "text_content",
  },
)

if err != nil {
  if notFoundError, ok := err.(*intrinsic.NotFoundError); ok {
    // Do something with the bad request ...
    log.Printf("Not Found : %+v", notFoundError)
  }
  log.Fatalf("Error creating event : %+v", err)
}
```

These errors are also compatible with the `errors.Is` and `errors.As` APIs, so you can access the error
like so:

```go

response, err := client.Events.CreateEventSync(
  context.TODO(),
  "not_a_real_event_type",
  map[string]interface{}{
    "caption": "text_content",
  },
)
if err != nil {
  var notFoundError *intrinsic.NotFoundError
  if errors.As(err, &notFoundError) {
    log.Printf("Not Found Error : %+v", notFoundError)
  }
  log.Fatalf("Failed to create event : %s", err)
}
```


## Contributing

While we value open-source contributions to this SDK, this library is generated programmatically.
Additions made directly to this library would have to be moved over to our generation code,
otherwise they would be overwritten upon the next generated release. Feel free to open a PR as
a proof of concept, but know that we will not be able to merge it as-is. We suggest opening
an issue first to discuss with us!

On the other hand, contributions to the `README.md` are always very welcome!
