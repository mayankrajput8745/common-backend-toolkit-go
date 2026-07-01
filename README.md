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
- Fasthttp adapter (high-performance) with pass-through native config (timeouts, concurrency, TLS, etc.)
- Graceful shutdown via `Server.Close()`
- Simple server initialization with `InitHTTPServer()`
- Standardized JSON response helpers (success, error, paginated, custom)

## Installation

```bash
go get github.com/mayankrajput8745/common-backend-toolkit-go
```

## Quick Start

```go
package main

import (
	"log"

	"github.com/mayankrajput8745/common-backend-toolkit-go/httpserver"
	"github.com/mayankrajput8745/common-backend-toolkit-go/httpserver/middleware"
)

func main() {
	srv, err := httpserver.InitHTTPServer("fasthttp", 8080, nil)
	if err != nil {
		log.Fatal(err)
	}

	srv.Use(middleware.Logger())

	srv.GET("/health", func(ctx httpserver.Context) {
		httpserver.SuccessResponse(ctx, map[string]string{"status": "ok"})
	})

	srv.GET("/users/:id", func(ctx httpserver.Context) {
		id := ctx.Param("id")
		httpserver.SuccessResponse(ctx, map[string]string{"user_id": id})
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

## Server Initialization

```go
func InitHTTPServer(framework string, port int32, cfg ServerConfig) (Server, error)
```

`cfg` is an `any` pass-through (`ServerConfig`) that each framework adapter type-asserts to its own native config type. Pass `nil` to use framework defaults.

```go
import fasthttp "github.com/mayankrajput8745/common-backend-toolkit-go/httpserver/frameworks/fasthttp"

cfg := &fasthttp.Config{
	Concurrency:   256 * 1024,
	MaxConnsPerIP: 100,
}

srv, err := httpserver.InitHTTPServer("fasthttp", 8080, cfg)
```

For the fasthttp adapter, `fasthttp.Config` is an alias for `fasthttp.Server` (from `valyala/fasthttp`) — set any of its fields (`ReadTimeout`, `WriteTimeout`, `Concurrency`, `TLSConfig`, ...). Must be passed by pointer. `Handler` is set internally and overwritten on `Start()`. If `ReadTimeout`, `WriteTimeout`, or `IdleTimeout` are left unset, they default to 5s, 10s, and 60s respectively.

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

- `middleware.Logger()` — Logs method, path, status code, and latency (ms)

Example:
```go
srv.Use(middleware.Logger())
```

## Response Helpers

`httpserver` ships a set of helpers over `Context.WriteJSON` for common response shapes, all returning the standard `Response{Success, Msg, Data}` envelope (or `PaginatedDataResponse` for pagination):

| Helper | Status Code |
|---|---|
| `SuccessResponse(ctx, data)` | 200 |
| `SuccessMsgResponse(ctx, msg)` | 202 |
| `CreatedResponse(ctx, data)` | 201 |
| `PaginatedResponse(ctx, data, pagination)` | 200 |
| `BadRequestResponse(ctx, msg)` | 400 |
| `UnauthorizedResponse(ctx, msg)` | 401 |
| `ForbiddenResponse(ctx, msg)` | 403 |
| `NotFoundResponse(ctx, msg)` | 404 |
| `TooManyRequestsResponse(ctx, msg)` | 429 |
| `InternalServerErrorResponse(ctx, msg)` | 500 |
| `CustomResponse(ctx, statusCode, success, msg, data)` | custom |
| `TextResponse(ctx, text)` | 200, `text/plain` |

Example:
```go
httpserver.SuccessResponse(ctx, user)
httpserver.NotFoundResponse(ctx, "user not found")
httpserver.PaginatedResponse(ctx, users, httpserver.Pagination{
	TotalCount: 120,
	Limit:      20,
	Page:       1,
	NextCursor: "abc123",
})
```

## Project Structure

```
common-backend-toolkit-go/
├── go.mod
├── go.sum
├── httpserver/
│   ├── main.go              # Factory (InitHTTPServer)
│   ├── types.go             # Type aliases
│   ├── response.go          # Response envelope + helpers
│   ├── contract/
│   │   └── types.go         # Core interfaces
│   ├── frameworks/
│   │   └── fasthttp/
│   │       ├── server.go    # Fasthttp implementation + Config
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
- [x] Configurable server tuning (timeouts, max connections, etc.) via `ServerConfig`
- [ ] Add tests

## License

MIT
