# Go Hello World Actions

A simple Go Hello World web server application demonstrating the use of different GitHub Action types wrapped with [witness-run-action](https://github.com/testifysec/witness-run-action) for attesting the build and test process.

## Application Details

This application is a minimal HTTP server that responds with "Hello, World!" on port 8080.

## GitHub Actions Workflow

The workflow in this repository demonstrates:

1. Using `witness-run-action` to wrap various GitHub Action types:
   - JavaScript actions (actions/checkout)
   - Composite actions (actions/setup-go)
   - Direct command execution
   - Docker container actions (goreleaser/goreleaser-action)

2. Creating attestations for each step in the process, including:
   - Environment attestations
   - GitHub attestations
   - SLSA attestations

3. Signing attestations using Sigstore and storing them in Archivista

## Running Locally

```bash
# Run the server
go run main.go

# Run tests
go test -v ./...
```

## Building with GoReleaser

```bash
goreleaser release --snapshot --clean
```