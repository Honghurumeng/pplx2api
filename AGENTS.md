# Repository Guidelines

## Project Structure & Modules
- `main.go`: entrypoint; starts Gin server and session refresher.
- `router/`: HTTP routes; mounts `/v1` and `/hf/v1`.
- `middleware/`: CORS and API key auth.
- `service/`: HTTP handlers (`ChatCompletionsHandler`, `ModelsHandler`, `HealthCheckHandler`).
- `core/`: Perplexity client, streaming, uploads.
- `config/`: env loading (`.env` via `godotenv`), model maps, runtime config.
- `job/`: session auto‑refresh; persists `sessions.json`.
- `utils/`, `logger/`: helpers and colored logging.
- Tooling: `.env.example`, `Dockerfile`, `README.md`.

## Build, Test, and Development Commands
- Setup env: `cp .env.example .env` then set `SESSIONS` and `APIKEY`.
- Run locally: `go run main.go` (listens on `ADDRESS`, default `0.0.0.0:8080`).
- Build binary: `go build -o bin/pplx2api ./main.go`.
- Format/Vet: `go fmt ./... && go vet ./...`.
- Docker: `docker build -t pplx2api:local .` then `docker run -p 8080:8080 --env-file .env pplx2api:local`.

## Coding Style & Naming Conventions
- Use `gofmt` (tabs); do not hand‑format.
- Package names: short lowercase; exported identifiers PascalCase; unexported camelCase.
- HTTP handlers end with `Handler` and live in `service/`.
- Filenames lowercase (e.g., `router.go`, `openai.go`).

## Testing Guidelines
- Framework: Go `testing` (no tests yet). Place `_test.go` in the same package.
- Names: `TestXxx`; prefer table tests for core logic.
- Run: `go test -v ./...` (optional `-cover` to inspect coverage).

## Commit & Pull Request Guidelines
- Commit style in history is short, imperative (e.g., "add gpt-5", "fix photo upload"). Keep messages focused; optionally use `feat:`/`fix:` prefixes.
- PRs should include: purpose, what/why, run instructions, env vars touched, linked issues, and screenshots/logs or `curl` for API changes.

## Security & Configuration Tips
- Never commit secrets. Keep `.env` local; update `.env.example` when adding vars.
- Required: `SESSIONS`, `APIKEY`. Optional knobs: `PROXY`, `IS_INCOGNITO`, `MAX_CHAT_HISTORY_LENGTH`, etc. See `README.md` for full list.
- Avoid logging sensitive values; redact tokens in examples and logs.

