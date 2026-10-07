# URL Shortener

A simple URL shortening service built with Go, PostgreSQL, and Redis.

The service converts long URLs into short identifiers using Base62 encoding. URL mappings are stored in PostgreSQL, while Redis is used as a caching layer for faster URL resolution.

## Tech Stack

- Go
- PostgreSQL
- Redis
- Docker
- Base62 Encoding

## How It Works

### Create a Short URL

A long URL is sent to the `/shorten` endpoint. The service generates a short identifier and stores the URL mapping in PostgreSQL.

```text
Long URL
   ↓
Base62
   ↓
Short URL
```

### Redirect

When a short URL is requested, the service first checks Redis. If the URL is not cached, it retrieves the mapping from PostgreSQL, stores it in Redis, and redirects the client.

```text
Request
   ↓
Redis
   ├── Hit  → Redirect
   └── Miss → PostgreSQL → Redis → Redirect
```

## Project Structure

```text
url-shortener/
├── cache/       # Redis
├── encoder/     # Base62 encoding
├── handlers/    # HTTP handlers
├── store/       # PostgreSQL
├── main.go      # Application entry point
├── go.mod
├── go.sum
└── .env.example
```

## Running Locally

Clone the repository:

```bash
git clone https://github.com/swastik-gautam/url-shortener.git
cd url-shortener
```

Install dependencies:

```bash
go mod download
```

Configure the required environment variables in `.env` and make sure PostgreSQL and Redis are running.

Start the application:

```bash
go run .
```

The server runs on:

```text
http://localhost:8080
```
