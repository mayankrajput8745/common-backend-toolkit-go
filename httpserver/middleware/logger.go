package middleware

import (
	"log"
	"os"
	"time"

	"github.com/mayankrajput8745/common-backend-toolkit-go/httpserver/contract"
)

// apiLogger is a custom logger for API request logging

var apiLogger = log.New(os.Stdout, "[API] ", log.LstdFlags)

func Logger() contract.MiddlewareFunc {
	return func(ctx contract.Context, next func()) {
		start := time.Now()
		next()
		latencyMs := time.Since(start).Milliseconds()
		apiLogger.Printf("%s %s | %d | %dms\n",
			ctx.Method(),
			ctx.Path(),
			ctx.StatusCode(),
			latencyMs,
		)
	}
}
