# CLAUDE.md

This file provides guidance to AI Agents such as Claude Code or ChatGPT Codex when working with code in this repository.

## General guidelines

Use idiomatic Go as possible. Prefer simple code than complex.

## Project Overview

Gemserve is a simple Gemini protocol server written in Go that serves static files over TLS-encrypted connections. The Gemini protocol is a lightweight, privacy-focused alternative to HTTP designed for serving text-based content.

### Development Commands

```bash
# Build, test, and format everything
make

# Run tests only
make test

# Build binaries to ./dist/ (gemserve, gemget, gembench)
make build

# Format code with gofumpt and gci
make fmt

# Run golangci-lint
make lint

# Run linter with auto-fix
make lintfix

# Clean build artifacts
make clean

# Run the server (after building)
./dist/gemserve

# Generate TLS certificates for development
certs/generate.sh
```

### Architecture

Core Components

- **cmd/gemserve/gemserve.go**: Entry point with TLS server setup, signal handling, and graceful shutdown
- **cmd/gemget/**: Gemini protocol client for fetching content
- **cmd/gembench/**: Benchmarking tool for Gemini servers
- **server/**: Request processing, file serving, and Gemini protocol response handling
- **gemini/**: Gemini protocol implementation (URL parsing, status codes, path normalization)
- **config/**: CLI-based configuration system
- **lib/logging/**: Structured logging package with context-aware loggers
- **lib/apperrors/**: Application error handling (fatal vs non-fatal errors)
- **uid/**: Connection ID generation for logging (uses external vendor package)

Key Patterns

- **Security First**: All file operations use `filepath.IsLocal()` and path cleaning to prevent directory traversal
- **Error Handling**: Uses structured errors via `lib/apperrors` package distinguishing fatal from non-fatal errors
- **Logging**: Structured logging with configurable levels via internal logging package
- **Testing**: Table-driven tests with parallel execution, heavy focus on security edge cases

Request Flow

1. TLS connection established on port 1965
2. Read up to 1KB request (Gemini spec limit)
3. Parse and normalize Gemini URL
4. Validate path security (prevent traversal)
5. Serve file or directory index with appropriate MIME type
6. Send response with proper Gemini status codes

Configuration

Server configured via CLI flags:
- `--listen`: Server address (default: localhost:1965)
- `--root-path`: Directory to serve files from
- `--dir-indexing`: Enable directory browsing (default: false)
- `--log-level`: Logging verbosity (debug, info, warn, error; default: info)
- `--response-timeout`: Response timeout in seconds (default: 30)
- `--tls-cert`: TLS certificate file path (default: certs/server.crt)
- `--tls-key`: TLS key file path (default: certs/server.key)
- `--max-response-size`: Maximum response size in bytes (default: 5242880)

Testing Strategy

- **server/server_test.go**: Path security and file serving tests
- **gemini/url_test.go**: URL parsing and normalization tests
- Focus on security edge cases (Unicode, traversal attempts, malformed URLs)
- Use parallel test execution for performance

Security Considerations

- All connections require TLS certificates (stored in certs/)
- Path traversal protection is critical - test thoroughly when modifying file serving logic
- Request size limited to 1KB per Gemini specification
- Input validation on all URLs and paths
