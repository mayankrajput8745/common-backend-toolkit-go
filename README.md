# common-backend-toolkit-go

A lightweight, **framework-agnostic HTTP server toolkit** for Go.

This toolkit helps you build HTTP servers using clean interfaces so your core application logic remains decoupled from the underlying HTTP framework (fasthttp, Gin, std `net/http`, etc.).

## Why This Toolkit?

- **Decoupling**: Write handlers against interfaces, not concrete frameworks.
- **Flexibility**: Swap HTTP frameworks with minimal code changes.
- **Production Ready**: Includes graceful shutdown support.
- **Simple & Lightweight**: Focused on routing, middleware, and context abstraction.
- **Extensible**: Easy to add new framework adapters.

## Features

- Clean `Context`, `HandlerFunc`, `MiddlewareFunc`, `Router`, and `Server` interfaces
- Middleware chaining with `next()` pattern
- Route grouping (`Group()`)
- Fasthttp adapter (high-performance)
- Graceful shutdown via `Server.Close()`
- Simple server initialization with `InitHTTPServer()`

## Installation

```bash
go get github.com/mayankrajput8745/common-backend-toolkit-go
```

## Quick Start

```go
package main

import (
	"log"

	"common-backend-toolkit/httpserver"
	"common-backend-toolkit/httpserver/middleware"
)

func main() {
	srv, err := httpserver.InitHTTPServer("fasthttp", 8080)
	if err != nil {
		log.Fatal(err)
	}

	srv.Use(middleware.Logger())

	srv.GET("/health", func(ctx httpserver.Context) {
		ctx.WriteJSON(200, map[string]string{"status": "ok"})
	})

	srv.GET("/users/:id", func(ctx httpserver.Context) {
		id := ctx.Param("id")
		ctx.WriteJSON(200, map[string]string{"user_id": id})
	})

	log.Println("Starting server...")
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
```

## Graceful Shutdown

The toolkit supports graceful shutdown using `Close()`:

```go
// On shutdown signal (SIGTERM, SIGINT, or manual trigger)
if err := srv.Close(); err != nil {
	log.Printf("Error during graceful shutdown: %v", err)
}
```

### Manual Close on Critical Errors

You can manually close the server when unhandled errors occur:

```go
func handleCriticalError(srv httpserver.Server, err error) {
	log.Printf("Critical error: %v. Closing server...", err)
	
	if closeErr := srv.Close(); closeErr != nil {
		log.Printf("Failed to close server: %v", closeErr)
	}
}
```

## Available Interfaces

### `Server`

```go
type Server interface {
	Router
	Use(middlewares ...MiddlewareFunc)
	Start() error
	Close() error
}
```

### `Router`

```go
type Router interface {
	GET(path string, handler HandlerFunc)
	POST(path string, handler HandlerFunc)
	PUT(path string, handler HandlerFunc)
	PATCH(path string, handler HandlerFunc)
	Group(prefix string) Router
}
```

### `Context`

```go
type Context interface {
	Method() string
	Path() string
	Param(key string) string
	Query(key string) string
	Body() []byte
	Header(key string) string
	StatusCode() int

	Status(code int) Context
	WriteJSON(statusCode int, v any) error
	String(s string) error
	SetHeader(key, value string)
}
```

## Middleware

Currently available:

- `middleware.Logger()` — Logs method, path, status code, and latency

Example:
```go
srv.Use(middleware.Logger())
```

## Project Structure

```
common-backend-toolkit-go/
├── go.mod
├── httpserver/
│   ├── main.go              # Factory (InitHTTPServer)
│   ├── types.go             # Type aliases
│   ├── contract/
│   │   └── types.go         # Core interfaces
│   ├── frameworks/
│   │   └── fasthttp/
│   │       ├── server.go    # Fasthttp implementation
│   │       ├── context.go   # Context adapter
│   │       └── main.go
│   └── middleware/
│       └── logger.go
```

## Extending the Toolkit

You can add support for other frameworks (Gin, Echo, stdlib `net/http`, etc.) by:

1. Implementing the `contract.Server`, `contract.Router`, and `contract.Context` interfaces.
2. Registering the new framework in `httpserver/main.go` inside `InitHTTPServer()`.

## Roadmap

- [ ] Add more middlewares (Recovery, CORS, RequestID, Timeout)
- [ ] Add Gin adapter
- [ ] Add standard `net/http` adapter
- [ ] Improved configuration options for server (timeouts, max connections, etc.)
- [ ] Add tests

## License

MIT
