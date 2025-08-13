# CRUSH

This file provides instructions for agentic coding agents to interact with this repository.

## Commands

- **Test:** `go test ./...`
- **Run a specific test:** `go test -run ^TestName$`
- **Lint:** `golangci-lint run` (This is a common Go linter, but not explicitly defined in the project. You might need to install it: `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest`)
- **Build:** `go build .`
- **Tidy:** `go mod tidy`

## Code Style

- **Imports:** Standard Go import grouping is preferred.
- **Formatting:** Use `gofmt` or `goimports` to format the code.
- **Types:** Use `PascalCase` for exported types and `camelCase` for unexported types.
- **Naming Conventions:**
    - Use descriptive names for variables and functions.
    - Follow Go's idiomatic naming conventions (e.g., `err` for errors, short variable names for short-lived variables).
- **Error Handling:** Use `fmt.Errorf` with `%w` to wrap errors and provide context.
- **Comments:** Add comments to explain complex logic, not to describe what the code does.
- **Dependencies:** Use `go mod tidy` to manage dependencies.

## Project Structure

- The main logic is in `gradient.go`.
- Tests are in `gradient_test.go`.
- The `example/` directory contains an example of how to use the library.
