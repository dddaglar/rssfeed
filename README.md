# RSSfeed

RSS blog aggregator pet project

## Requirements

- Go (see `go.mod` for the language version)
- PostgreSQL running and accessible

## Install

Install the CLI binary (recommended for production):

```sh
go install github.com/denizekindaglar/rssfeed@latest
```

After a successful `go install`, the `gator` binary will be in your Go bin directory (`$GOBIN` or `$GOPATH/bin`) and is a statically compiled executable you can run without the Go toolchain.

For development you can use:

```sh
go run .
```

**Note:** `go run .` is for development. `gator` (the compiled binary) is recommended for production.

## Configuration

Create a config file at `$HOME/.gatorconfig.json` with the database URL and optionally a current user:

```json
{
  "db_url": "postgres://user:password@localhost:5432/dbname?sslmode=disable",
  "current_user_name": "username"
}
```

The program reads/writes this file using the logic in `config.Config` — see `internal/config/config.go`.

Example existing config file used during development: `/Users/denizekindaglar/.gatorconfig.json`

## Usage

Build or install, then run the CLI:

```sh
gator <command> [args...]
```

### Commands

Commands implemented in this repo (handlers in code):

- `login <name>` — set current user (see `main.handlerLogin`)
- `register <name>` — create and set a new user (see `main.handlerRegister`)
- `reset` — delete all users (see `main.handlerReset`)
- `users` — list users (see `main.handlerUsers`)
- `addfeed <name> <url>` — add a feed (requires logged-in user) (see `main.handlerAddFeed`)
- `feeds` — list feeds (see `main.handlerFeeds`)
- `follow <feed_url>` — follow a feed (requires logged-in user) (see `main.handlerFollow`)
- `following` — list feeds you follow (requires logged-in user) (see `main.handlerFollowing`)
- `unfollow <feed_url>` — unfollow a feed (requires logged-in user) (see `main.handlerUnfollow`)
- `browse [limit]` — show recent posts for current user (see `main.handlerBrowse`)
- `agg <duration>` — continuously fetch feeds every given duration (e.g. `1m`) (see `main.handlerAgg`)

The command dispatcher and registrations are in `main.go`.

## Database

Create a Postgres database and run the SQL migrations from `sql/schema/` to create tables (users, feeds, feed_follows, posts, etc.). See `sql/schema/` for migration files (example: `sql/schema/001_users.sql`, `sql/schema/002_feeds.sql`, `sql/schema/005_posts.sql`).

## Notes

- The program uses the generated DB layer in `internal/database` (sqlc)
- `gator` is a compiled binary — build with `go build` or `go install` for production use

## Ideas for Extending

- Add sorting and filtering options to the `browse` command
- Add pagination to the `browse` command
- Add concurrency to the `agg` command so that it can fetch more frequently
- Add a search command that allows for fuzzy searching of posts
- Add bookmarking or liking posts
- Add a TUI that allows you to select a post in the terminal and view it in a more readable format (either in the terminal or open in a browser)
- Add an HTTP API (and authentication/authorization) that allows other users to interact with the service remotely
- Write a service manager that keeps the `agg` command running in the background and restarts it if it crashes