package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/intrinsic-org/intrinsic-go/events"
	"github.com/intrinsic-org/intrinsic-go/option"
)

func main() {
	// Define command line flags
	baseURL := flag.String("base-url", "", "Base URL for the API")
	eventTypeName := flag.String("event-type", "", "The type of event being created")
	dataJson := flag.String("data", "", "JSON file containing event data")
	apiKey := flag.String("api-key", "", "API key for authentication")
	idempotencyKey := flag.String("idempotency-key", "", "Idempotency key for the request")
	flag.Parse()

	if *baseURL == "" {
		log.Fatal("--base-url is required")
	}

	if *eventTypeName == "" {
		log.Fatal("--event-type is required")
	}

	// Parse event data from JSON string if provided, otherwise use empty object
	var eventData map[string]interface{}
	if *dataJson != "" {
		if err := json.Unmarshal([]byte(*dataJson), &eventData); err != nil {
			log.Fatalf("Failed to parse JSON data: %v", err)
		}
	} else {
		eventData = make(map[string]interface{})
	}

	// Create options for the client
	var opts []option.RequestOption
	opts = append(opts, option.WithBaseURL(*baseURL))

	// Add API key if provided
	if *apiKey != "" {
		opts = append(opts, option.WithAPIKey(*apiKey))
	}

	// Create a client with the options
	client := events.NewClient(opts...)

	// Add idempotency key if provided
	if *idempotencyKey != "" {
		headers := make(http.Header)
		headers.Set("X-Idempotency-Key", *idempotencyKey)
		opts = append(opts, option.WithHTTPHeader(headers))
	}

	// Create the event
	ctx := context.Background()
	resp, err := client.CreateEventAsync(ctx, *eventTypeName, eventData, opts...)
	if err != nil {
		log.Fatalf("Failed to create event: %v", err)
	}

	// Print the response
	prettyResp, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatalf("Failed to format response: %v", err)
	}
	fmt.Printf("Event created successfully:\n%s\n", string(prettyResp))
}
