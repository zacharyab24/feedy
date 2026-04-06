# Feedy

A self-hosted RSS/Atom feed reader built with Go, HTMX, and Tailwind CSS.

## Features

- Add and manage RSS/Atom feeds
- Browse feed items in a unified timeline
- Mark items as read/unread
- Star items for later
- Filter items by feed
- Manual feed refresh

## Tech Stack

- **Go** (stdlib `net/http`, `html/template`)
- **HTMX** for dynamic UI without writing JavaScript
- **Tailwind CSS** for styling
- **SQLite** via `modernc.org/sqlite` (pure Go, no CGO)
- **gofeed** for RSS/Atom parsing

## Getting Started

### Prerequisites

- Go 1.22+
- Docker (optional)

### Run locally

```sh
go mod download
go run main.go
```

The app will be available at `http://localhost:8081`. The SQLite database is created automatically in `data/`.

### Run with Docker (future)

I havent done the docker stuff yet but will support docker compose

## Project Structure

```
feedy/
├── main.go
├── internal/
│   ├── db/          # Database setup and migrations
│   ├── models/      # Feed and Item structs
│   ├── store/       # Database CRUD operations
│   ├── fetch/       # RSS/Atom feed fetching and parsing
│   └── handler/     # HTTP handlers
├── templates/       # Go HTML templates
│   └── partials/    # HTMX fragment templates
└── static/          # Static assets
```
