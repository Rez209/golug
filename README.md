# GoLug 🚀

**GoLug** is a fast, standalone CLI tool for scaffolding production-ready microservices — no config, no boilerplate hunting, no copy-pasting from your last project.

One command generates a working service with Docker, a database connection, health checks, and (for Go) a full local observability stack — ready to `docker compose up` and start coding.

```bash
golug new auth-service --lang go --port 8081
cd auth-service && docker compose up --build
```

That's it. You have a running microservice with Postgres, health checks, and monitoring wired up.

## Why

Every new backend project starts the same way: set up the folder structure, wire up the database, write a Dockerfile, configure env vars, add a health check — the same 30 minutes of setup, every time, in every language. GoLug automates that so you can get to the actual logic immediately.

## Features

- **Multi-language templates:** Go, Python, JavaScript (Node.js), and C++
- **Docker-first:** generates a working `Dockerfile` and `docker-compose.yml` per service
- **Database-ready:** pre-wired PostgreSQL connection and `.env` config (Go, Python, JS)
- **Health checks built in:** every generated service ships with a `/api/health` endpoint
- **Zero runtime dependencies:** GoLug itself is a single compiled binary — all templates are embedded, nothing to download at generation time
- **Git-ready:** automatically runs `git init` in the new project

## Installation

**Option 1 — go install (recommended if you have Go):**

```bash
go install github.com/Rez209/golug@latest
```

**Option 2 — download a binary:**

Grab the latest build for your OS from the [Releases](https://github.com/Rez209/golug/releases) page.

**Option 3 — build from source:**

```bash
git clone https://github.com/Rez209/golug.git
cd golug
go build -o golug .
```

## Usage

```bash
golug new <service-name> --lang <language> --port <port>
```

| Flag | Description | Default |
|------|-------------|---------|
| `--lang`, `-l` | Target language: `go`, `python`, `js`, `cpp` | `go` |
| `--port`, `-p` | Port the service will listen on | `8080` |

### Examples

```bash
# Go microservice with Postgres + Prometheus/Grafana monitoring
golug new auth-service --lang go --port 8081

# Python microservice (FastAPI + SQLAlchemy + Postgres)
golug new data-service --lang python --port 8082

# Node.js microservice (Express + Postgres)
golug new gateway-service --lang js --port 8083
```

## Supported languages

| Flag | Stack |
|------|-------|
| `go` | Go, PostgreSQL, Prometheus + Grafana, standard project layout |
| `python` | FastAPI, SQLAlchemy, PostgreSQL |
| `js` | Node.js, Express, PostgreSQL |
| `cpp` | CMake, basic Dockerfile |

## Running a generated service

```bash
cd auth-service
docker compose up --build
curl http://localhost:8081/api/health
```

## Roadmap

- [ ] TypeScript template
- [ ] Interactive mode (`golug new` with prompts, no flags required)
- [ ] Custom module path support for Go templates
- [ ] Optional auth/JWT scaffolding

## Contributing

Issues and pull requests are welcome — especially template improvements for C++ and JS, and reports of anything that breaks on your machine.

## License

MIT