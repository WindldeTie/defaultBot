# WARP.md

This file provides guidance to WARP (warp.dev) when working with code in this repository.

## Commands you’ll use most

All commands assume Go is installed (go.mod declares `go 1.25.0`).

- Install deps
  ```
  go mod download
  ```
- Build (module-wide)
  ```
  go build ./...
  ```
- Build binary in repo root
  ```
  go build -o bin/defaultBot .
  ```
- Run the bot (loads .env automatically via godotenv)
  ```
  go run .
  ```
  - Alternatively, set env vars for a single run (PowerShell):
    ```
    $env:BOT_TOKEN={{BOT_TOKEN}}; $env:ADMIN_ID='123456789'; go run .
    ```
- Format and basic static checks (no external linters configured)
  ```
  go fmt ./...
  go vet ./...
  ```
- Test (no tests currently present, but patterns below apply once added)
  ```
  go test ./...
  # Run tests in one package
  go test ./handler
  # Run a single test by name (regex match)
  go test ./handler -run "^TestHandleUpdate$"
  ```

## High-level architecture and flow

- Entry point (`main.go`)
  - Loads environment from `.env` via `github.com/joho/godotenv`.
  - Requires `BOT_TOKEN`; process exits if missing.
  - Initializes Telegram client (`github.com/go-telegram-bot-api/telegram-bot-api/v5`).
  - Deletes any existing webhook with `DropPendingUpdates: true` to switch to long polling.
  - Instantiates `handler.Handler` and calls `Start(false)` (debug logging off).

- Update processing (`handler` package)
  - `Start(debug bool)` configures `bot.Debug` and begins reading from `GetUpdatesChan` with a 60s timeout.
  - Spawns a background console goroutine that reads stdin:
    - `exit` cleanly terminates the process; unknown commands are logged.
  - For each incoming `tgbotapi.Update`, `HandleUpdate` routes by message text:
    - `/start` → `handleStart` replies with a greeting (currently Russian: "Привет!").
    - Any other text is echoed back verbatim.

- Configuration and environment
  - `.env` is expected at repo root (already tracked). Relevant variables:
    - `BOT_TOKEN` (required): Telegram bot token.
    - `ADMIN_ID` (optional): parsed to `int64` in `Start`; currently passed to `HandleUpdate` but not enforced—intended for future admin-only logic.

- Debugging and modes
  - Debug logs from the Telegram client are controlled by the `debug` argument to `Start` (currently hardcoded to `false` in `main.go`). Set to `true` to enable verbose request/response logging.

## Notes specific to this repo

- The bot uses long polling (webhooks are explicitly cleared on startup).
- Command routing is centralized in `handler.HandleUpdate` via a `switch` on the first token of the message. Add new commands by extending this switch and extracting handlers as needed.
