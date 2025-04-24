# Variance Event CLI

A simple command-line tool to test the Variance (formerly Intrinsic) Go client E2E

## Usage

```
go run cmd/create_event_async/main.go \
  --event-type $EVENT_TYPE \
  --api-key $API_KEY \
  --base-url https://intrinsicapi.com \
  --idempotency-key $OPTIONAL_IDEMPOTENCY_KEY
```