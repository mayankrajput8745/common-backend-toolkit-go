package middleware

import (
	"fmt"
	"time"

	"github.com/mayankrajput8745/common-backend-toolkit-go/httpserver/contract"
)

func Logger() contract.MiddlewareFunc {
	return func(ctx contract.Context, next func()) {
		start := time.Now()
		next()
		fmt.Printf("[%s] %s | %d | %s\n",
			ctx.Method(),
			ctx.Path(),
			ctx.StatusCode(),
			time.Since(start),
		)
	}
}
