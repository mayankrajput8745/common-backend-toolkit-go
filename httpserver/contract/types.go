package contract

// ServerConfig carries framework-specific server tuning options. Each
// framework implementation defines and type-asserts its own concrete
// config type (e.g. fasthttp.Config), so this stays a plain pass-through.
type ServerConfig any

type HandlerFunc func(ctx Context)

type MiddlewareFunc func(ctx Context, next func())

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

type Router interface {
	GET(path string, handler HandlerFunc)
	POST(path string, handler HandlerFunc)
	PUT(path string, handler HandlerFunc)
	PATCH(path string, handler HandlerFunc)
	Group(prefix string) Router
}

type Server interface {
	Router
	Use(middlewares ...MiddlewareFunc)
	Start() error
	Close() error
}
