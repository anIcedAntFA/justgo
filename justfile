# justgo — task runner (https://github.com/casey/just)
# Run `just` with no arguments to list every recipe.

# List all recipes
default:
    @just --list

# One-time setup: install pinned tools + git hooks.
setup:
    mise install
    lefthook install
    @echo "✅ setup complete — tools installed, hooks active"

# Format everything in place (Go + docs/config)
fmt:
    gofumpt -w .
    dprint fmt

# Run every check without writing — this is exactly what CI runs
check:
    #!/usr/bin/env bash
    set -euo pipefail
    echo "▶ gofumpt";       test -z "$(gofumpt -l .)" || { echo "run 'just fmt'"; gofumpt -l .; exit 1; }
    echo "▶ go vet";        go vet ./...
    echo "▶ golangci-lint"; golangci-lint run ./...
    echo "▶ dprint";        dprint check
    echo "▶ test";          go test -race ./...
    just secrets

# Run the test suite (with the race detector)
test:
    go test -race ./...

# Lint Go only
lint:
    golangci-lint run ./...

# Scan the repository (working tree + history) for leaked secrets
secrets:
    gitleaks git . --no-banner --redact

# Tidy every module's dependencies
tidy:
    go mod tidy
