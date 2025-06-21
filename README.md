# SAST App

This project provides a simple web application for managing SAST (Static Application Security Testing) tasks and displaying scan reports. The backend is written in Go and uses MySQL for data storage. A minimal HTML page is included for interacting with the API.

## Project Layout

- **cmd/server** – main application entry point.
- **internal/** – backend packages (database setup, HTTP handlers and models).
- **web/** – static files served to the client.

## Building

Ensure Go is installed and run:

```bash
go mod tidy
go build ./cmd/server
```

`MYSQL_DSN` environment variable controls the database connection (default is `root:password@tcp(localhost:3306)/sast`).

## Running

Start the server and open `web/index.html` in a browser:

```bash
./server
```

The page allows creating new tasks and viewing reports in a clean and simple interface.
