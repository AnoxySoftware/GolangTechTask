# Golang Tech Test

## Task

- Provided a `Go` implementation of the `gRPC` service in the `cmd/` directory.

- Implement a `DynamoDB` based store for this `gRPC` service

  This is implemented in `internal/server` and uses the dynamodb model in `internal/models`

- Add pagination to the `ListVoteables` RPC call

  This is included as part of the DynamoDB server, but not currently the mock server

- Provide adequate test coverage for this simple service

  There are currently tests to ensure the gRPC service is working, but not automated tests for the DynamoDB side, however there is an included [magefile](https://magefile.org/) in `magefile.go` which allows for easily running [grpcurl](https://github.com/fullstorydev/grpcurl)

  Examples:
  - `mage grpcurlList` - lists the voteables
  - `mage grpcurlCreate` - creates a new voteable, can optionally take `QUESTION` and `ANSWERS` environment variables
  - `mage grpcurlCast` - casts a vote, uses `INDEX` and `UUID` environment variables

## How to impress

- Observability

  - [Apex log](https://github.com/apex/log) is used to provide a structured logging API

  - [OpenTelemetry](https://github.com/open-telemetry/opentelemetry-go-contrib) is linked into the gRPC server to provide tracing, however there are no exporters configured currently.

  - I ran out of time to implement metrics, but most of the required code should be there.

- Configuration / Secrets management

  - [ff](https://github.com/peterbourgon/ff) is used for neater configuration of flags, and also allows reading from the environment

## References

Here are some useful links I used to aid my work:

- [Behavior of server.GracefulStop() in golang](https://stackoverflow.com/questions/55797865/behavior-of-server-gracefulstop-in-golang)
- [Testing a gRPC service](https://stackoverflow.com/questions/42102496/testing-a-grpc-service)
- [Common Design Patterns - List Pagination](https://cloud.google.com/apis/design/design_patterns#list_pagination)
