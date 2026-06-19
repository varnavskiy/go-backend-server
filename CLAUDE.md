# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Run the server (must be run from cmd/server/ so the relative .system.env path resolves)
cd cmd/server && go run server.go

# Build
go build -o backendServer ./cmd/server/

# Run tests
go test ./...

# Run a single test
go test ./internal/configs/... -run TestFunctionName -v

# Docker build
docker build -t backendsrv .
```

## Architecture

This is a Go HTTP server using [Gin](https://github.com/gin-gonic/gin) with graceful shutdown.

**Entry point:** `cmd/server/server.go` — wires config, logger, Gin router, and the HTTP server lifecycle.

**Config:** `internal/configs/config.go` loads all configuration from `.system.env` via `godotenv`. The `init()` function uses a relative path `../../.system.env`, so the binary must be run from `cmd/server/` (or the path adjusted). Config covers port, Gin mode, and lumberjack log rotation settings.

**Logging:** `lumberjack` rotates log files under `cmd/server/logs/gin.log`. Logs go to both the file and stdout via `io.MultiWriter`.

**Key dependencies declared but not yet wired:**
- `go.mongodb.org/mongo-driver/v2` — MongoDB client is in `go.mod` but no usage exists yet.
