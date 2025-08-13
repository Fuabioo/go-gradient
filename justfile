# Go Gradient Justfile

# Default recipe to display available commands
default:
    @just --list

# Run all tests
test:
    go test ./...

# Run a specific test by name
test-run name:
    go test -run ^{{name}}$

# Build the project
build:
    go build .

# Format code
fmt:
    gofmt -w .
    goimports -w .

# Tidy dependencies
tidy:
    go mod tidy

# Lint code (requires golangci-lint)
lint:
    golangci-lint run

# Install golangci-lint
install-linter:
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run example
example:
    cd example && go run main.go

# Clean build artifacts
clean:
    go clean

# Run all checks (format, tidy, lint, test)
check: fmt tidy lint test

# Show module info
mod-info:
    go list -m all