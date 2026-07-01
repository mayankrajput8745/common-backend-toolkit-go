package fasthttp

import (
	"context"
	"fmt"
	"time"

	"github.com/mayankrajput8745/common-backend-toolkit-go/httpserver/contract"

	fhr "github.com/fasthttp/router"
	fh "github.com/valyala/fasthttp"
)

const (
	defaultReadTimeout  = 5 * time.Second
	defaultWriteTimeout = 10 * time.Second
	defaultIdleTimeout  = 60 * time.Second
)

// Config is the fasthttp server's native tuning surface. Set any of its
// fields (Concurrency, MaxConnsPerIP, TLSConfig, ...); Handler is set
// internally and any value passed for it is overwritten on Start.
type Config = fh.Server

type fasthttpServer struct {
	server      *fh.Server
	router      *fhr.Router
	port        int32
	middlewares []contract.MiddlewareFunc
}

type fasthttpGroup struct {
	group *fhr.Group
}

func New(port int32, cfg contract.ServerConfig) contract.Server {
	serverCfg, _ := cfg.(Config)

	if serverCfg.ReadTimeout <= 0 {
		serverCfg.ReadTimeout = defaultReadTimeout
	}
	if serverCfg.WriteTimeout <= 0 {
		serverCfg.WriteTimeout = defaultWriteTimeout
	}
	if serverCfg.IdleTimeout <= 0 {
		serverCfg.IdleTimeout = defaultIdleTimeout
	}

	return &fasthttpServer{
		router: fhr.New(),
		port:   port,
		server: &serverCfg,
	}
}

func (s *fasthttpServer) Use(middlewares ...contract.MiddlewareFunc) {
	s.middlewares = append(s.middlewares, middlewares...)
}

func (s *fasthttpServer) GET(path string, handler contract.HandlerFunc) {
	s.router.GET(path, wrap(handler))
}

func (s *fasthttpServer) POST(path string, handler contract.HandlerFunc) {
	s.router.POST(path, wrap(handler))
}

func (s *fasthttpServer) PUT(path string, handler contract.HandlerFunc) {
	s.router.PUT(path, wrap(handler))
}

func (s *fasthttpServer) PATCH(path string, handler contract.HandlerFunc) {
	s.router.PATCH(path, wrap(handler))
}

func (s *fasthttpServer) Group(prefix string) contract.Router {
	return &fasthttpGroup{group: s.router.Group(prefix)}
}

func (s *fasthttpServer) Start() error {
	addr := fmt.Sprintf(":%d", s.port)
	fmt.Printf("Server Started and Listening on PORT: %s\n", addr)

	s.server.Handler = s.chainedHandler()

	return s.server.ListenAndServe(addr)
}

func (s *fasthttpServer) Close() error {
	if s.server == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.server.ShutdownWithContext(ctx)
}

func (s *fasthttpServer) chainedHandler() fh.RequestHandler {
	return func(ctx *fh.RequestCtx) {
		c := newContext(ctx)
		chain(c, s.middlewares, func() {
			s.router.Handler(ctx)
		})()
	}
}

func chain(ctx contract.Context, middlewares []contract.MiddlewareFunc, final func()) func() {
	if len(middlewares) == 0 {
		return final
	}
	return func() {
		middlewares[0](ctx, chain(ctx, middlewares[1:], final))
	}
}

func (g *fasthttpGroup) GET(path string, handler contract.HandlerFunc) {
	g.group.GET(path, wrap(handler))
}

func (g *fasthttpGroup) POST(path string, handler contract.HandlerFunc) {
	g.group.POST(path, wrap(handler))
}

func (g *fasthttpGroup) PUT(path string, handler contract.HandlerFunc) {
	g.group.PUT(path, wrap(handler))
}

func (g *fasthttpGroup) PATCH(path string, handler contract.HandlerFunc) {
	g.group.PATCH(path, wrap(handler))
}

func (g *fasthttpGroup) Group(prefix string) contract.Router {
	return &fasthttpGroup{group: g.group.Group(prefix)}
}

func wrap(handler contract.HandlerFunc) fh.RequestHandler {
	return func(ctx *fh.RequestCtx) {
		handler(newContext(ctx))
	}
}
