# gist-backend

An API for saving and retrieving code snippets (gists) from a database.

## Requirements

- Go 1.21

## Installation

1. Clone the repository
2. Initialize the database with `cat sql/create.sql | sqlite3 dev.sqlite3`
3. Run `go run cmd/api.go`
