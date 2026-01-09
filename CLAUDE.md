# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Common commands
- Run server: `make run` (uses `cmd/server`)
- Build binary: `make build` (writes to `cmd/server`)
- Run all tests: `go test ./...`
- Run single test package: `go test ./internal/<package>`

## Architecture overview
- Entry point is `cmd/server/main.go`, which calls `internal/initialize.Run()`.
- Startup flow in `internal/initialize/run.go`: load config (Viper + `.env`), init logger, init Postgres (GORM), init Redis, then start Gin HTTP server with graceful shutdown via `WaitForSignal()` and `Shutdown()`.
- Configuration structs live in `pkg/setting/section.go`; defaults and validation in `internal/initialize/config.go`.
- Global singletons are in `global/` (config, logger, DB, Redis).
- HTTP routing is in `internal/routers/router.go` with middleware chain from `internal/middlewares/` and versioned routes under `/api`.
- Request flow: controller (`internal/controllers`) → service (`internal/services`) → repository (`internal/repositories`).
- GORM models are in `internal/po/`; models register themselves and are migrated in `internal/initialize/database.go` via `po.All()`.
- Shutdown hooks are registered via `internal/initialize/shutdown.go` (`OnShutdown` and `Shutdown`).
