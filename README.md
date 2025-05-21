# Blog Aggregator

- Simple RSS feed aggregator in GO
- CLI tools that allows users to:
  - Add RSS feeds from across the internet to be collected
  - Store the collected posts in a PostgreSQL database
  - Follow and unfollow RSS feeds that other users have added
  - View summaries of the aggregated posts in the terminal,
    with a link to the full post

## Learning Goals

- Learn how to integrate a `Go` application with a `PostgreSQL` database
- Practice using SQL skills to query and migrate database
- Learn how to write a long-running service that continuously fetches new posts
  from RSS feeds and stores them in the database

## Package INFO

- internal/config/
  - Responsible for reading and writing the JSON file.
- internal/database/
  - auto generated sql.go by sqlc
- sql/
  - sql files for queried & schema. (used for sqlc)

## Installation

Make sure you have the latest [Go toolchain](https://golang.org/dl/) installed as well as a local Postgres database. You can then install `gator` with:

```bash
go install ...
```

- You need to install `goose` for migrate DB
  - `go install github.com/pressly/goose/v3/cmd/goose@latest`
  - move to sql/schema directiry when you need to migrate your database
    - `goose postgres postgres://username:@localhost:5432/gator up`

## Config

Create a `~/.gatorconfig.json` file in your home directory with the following structure:

```json
{
  "db_url": "postgres://username:@localhost:5432/database?sslmode=disable"
}
```

Replace the values with your database connection string.

## Usage

Create a new user:

```bash
gator register <name>
```

Add a feed:

```bash
gator addfeed <url>
```

Start the aggregator:

```bash
gator agg 30s
```

View the posts:

```bash
gator browse [limit]
```

There are a few other commands:

- `gator login <name>` - Log in as a user that already exists
- `gator users` - List all users
- `gator feeds` - List all feeds
- `gator follow <url>` - Follow a feed that already exists in the database
- `gator unfollow <url>` - Unfollow a feed that already exists in the database
