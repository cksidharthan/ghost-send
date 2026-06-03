# GhostSend

A secure secret sharing platform. Secrets are encrypted at rest, expire automatically, and can be limited to a set number of views.

## Features

- Encrypted secret storage in PostgreSQL
- Configurable expiration (5 minutes to 7 days)
- View count limits
- Optional password protection

## Stack

- Backend: Go 1.23, Gin, PostgreSQL, sqlc, golang-migrate, Uber FX, Zap
- Frontend: Nuxt 3, TailwindCSS (embedded in the Go binary at build time)

## Getting Started

### Prerequisites

- Go 1.23+
- Node.js 22+
- PostgreSQL 17+
- [Task](https://taskfile.dev) (optional, for task runner)

### Environment Variables

Copy `.env` and fill in your values:

```
POSTGRES_HOST=localhost
POSTGRES_PORT=5433
POSTGRES_USER=username
POSTGRES_PASSWORD=password
POSTGRES_DB=ghostsend
POSTGRES_SSL_MODE=disable
PORT=8080
LOG_LEVEL=info
```

### Local Development

Start a PostgreSQL instance:

```bash
docker compose up -d postgres
```

Run the app (builds the frontend first):

```bash
task up
# or without Task:
cd frontend && npm install && npm run generate && cd ..
go run main.go
```

The application is available at http://localhost:8080.

### Build

```bash
task build
# or without Task:
cd frontend && npm run generate && cd ..
CGO_ENABLED=0 go build -ldflags "-w -s" -o ghost-send main.go
```

### Docker Compose (full stack)

```bash
docker compose up -d
```

This starts PostgreSQL (port 5433), the backend (port 8080), and the frontend (port 6666).

## API

| Method | Path | Description |
|--------|------|-------------|
| GET | /healthz | Health check |
| POST | /api/v1/secrets | Create a secret |
| POST | /api/v1/secrets/:id | Retrieve a secret |
| GET | /api/v1/secrets/:id/status | Check secret status |

### POST /api/v1/secrets

```json
{
  "secret_text": "my secret",
  "password": "required",
  "expiration": "1h",
  "views": 1
}
```

Valid `expiration` values: `5m`, `1h`, `1d`, `7d`. Defaults to `1h` if omitted.

### POST /api/v1/secrets/:id

Password is always required and must be sent in the request body.

```json
{
  "password": "required"
}
```

## Project Structure

```
cmd/          application entry point
db/           migrations and sqlc-generated code
frontend/     Nuxt 3 frontend (built output embedded in binary)
pkg/
  config/     environment config
  daemon/     background cleanup (expired secret janitor)
  logger/     zap logger setup
  postgres/   database connection
  router/     HTTP router and handlers
  secret/     secret domain (HTTP handlers + service layer)
  ui/         static file serving for embedded frontend
```

## License

MIT
