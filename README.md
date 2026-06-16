# Marginalia

A self-hosted personal library — [Readarr](https://github.com/Readarr/Readarr) replacement with the functionality of Goodreads.

Marginalia ships as a single static Go binary with an embedded React frontend, so deployment
is one file (or one container) with no external runtime dependencies.

> **Status:** **Very** early development.

## Why?

I'm using this project to scratch a few different itches.

1. With Readarr stopping development, I haven't found a replacement that I like, or at least, one that works how I want it to.
2. I've been putting off learning Go. This is a great real-world project to help me learn.

## Notes on AI

**Is this vibe-coded:** No. I don't consider it vibe-coded.

I've been a full-stack software engineer for nearly 20 years. My favorite development stack is Ruby on Rails, but I wanted to learn Go and I wanted a small single binary app for Marginalia. I'm using Claude to supplement my lack of Go skills and to automate parts I enjoy the least, like writing tests. After all, this is a project for me to fulfill my own need and _to have fun_. If it ends up being used by others, **great!** If not, I'm ok with that too.

If you're a Go developer that would like to help with the backend, I'd be more than happy to collaborate with someone more experienced than me!

## Features

- Single static binary — frontend is embedded via `go:embed`, no separate web server needed.
- Pure-Go SQLite (`modernc.org/sqlite`) to keep things small and fast.
- Graceful shutdown with configurable timeouts, fully driven by environment variables.
- Multi-stage Docker build and Compose setup with a persistent data volume.
- Out-of-the-box Unraid compatibility.

## Tech stack

| Layer      | Choice                                                                                    |
| ---------- | ----------------------------------------------------------------------------------------- |
| Backend    | Go 1.25,[`chi`](https://github.com/go-chi/chi) router                                     |
| Frontend   | React +[Vite+](https://viteplus.dev/) (`vp`), pnpm, TypeScript                            |
| Database   | SQLite via[`modernc.org/sqlite`](https://pkg.go.dev/modernc.org/sqlite) (pure Go, no CGO) |
| Migrations | [`pressly/goose/v3`](https://github.com/pressly/goose) (embedded SQL, library mode)       |
| Packaging  | Multi-stage Docker image, Docker Compose                                                  |

## Project layout

```text
cmd/marginalia/main.go          entrypoint: open DB -> migrate -> start server
internal/server/
    server.go                   Server type, New(cfg), Run(ctx) with graceful shutdown
    config.go                   Config + ConfigFromEnv() + defaults
    routes.go                   apiRouter() (mounts /api)
    handlers.go                 handleHealth (GET /api/health)
    spa.go                      spaHandler() + go:embed web/dist (SPA + 404 logic)
    web/                        React/Vite+ frontend; builds to web/dist
internal/store/
    store.go                    Open(ctx, path) *sql.DB (WAL, foreign_keys, busy_timeout)
    migrate.go                  Migrate(ctx, db) via goose; embeds migrations/*.sql
    migrations/                 versioned SQL migrations
docker/
    Dockerfile                  3-stage build: frontend -> go build -> alpine runtime
    docker-compose.yml          builds image, surfaces env vars, mounts /data volume
```

## Getting started

### Prerequisites

- [Go](https://go.dev/) 1.25+
- [pnpm](https://pnpm.io/) (the frontend uses the [Vite+](https://viteplus.dev/) `vp` toolchain)

### Run locally

The frontend `web/dist` directory is embedded at build time via `go:embed`, so it must exist
before you build or run the server.

```bash
# 1. Build the frontend
cd internal/server/web && pnpm install && pnpm build && cd -

# 2. Run the server
go run ./cmd/marginalia
```

The server listens on `http://localhost:8090` by default. Health check:

```bash
curl http://localhost:8090/api/health
```

### Build a static binary

```bash
CGO_ENABLED=0 go build -o marginalia ./cmd/marginalia
```

### Run with Docker

The image builds the frontend and Go binary in separate stages, so you don't need Go or
pnpm installed locally — only Docker.

```bash
docker compose -f docker/docker-compose.yml up --build
```

The Compose setup runs as a non-root user and persists the SQLite database in the
`marginalia-data` volume mounted at `/data`.

## Configuration

All configuration is via environment variables. Every value has a sensible default.

| Variable                | Default                                           | Purpose                  |
| ----------------------- | ------------------------------------------------- | ------------------------ |
| `HOST`                  | all interfaces                                    | Listen host              |
| `PORT`                  | `8090`                                            | Listen port              |
| `DATABASE_PATH`         | `marginalia.db` (`/data/marginalia.db` in Docker) | SQLite database file     |
| `HTTP_READ_TIMEOUT`     | `15s`                                             | Server read timeout      |
| `HTTP_WRITE_TIMEOUT`    | `15s`                                             | Server write timeout     |
| `HTTP_IDLE_TIMEOUT`     | `60s`                                             | Server idle timeout      |
| `HTTP_SHUTDOWN_TIMEOUT` | `10s`                                             | Graceful shutdown budget |

## Development

```bash
# Build and vet everything
go build ./... && go vet ./...

# Verify a CGO-free build (required — must stay CGO_ENABLED=0 compatible)
CGO_ENABLED=0 go build ./...
```

Frontend tooling is driven by Vite+ (`vp`) from `internal/server/web`. Use `pnpm dev` for a
hot-reloading frontend dev server, and `pnpm build` to produce the embedded `web/dist`.

## License

[MIT](LICENSE)
