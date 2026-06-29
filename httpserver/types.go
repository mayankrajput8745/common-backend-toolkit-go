package httpserver

import "github.com/mayankrajput8745/common-backend-toolkit-go/httpserver/contract"

type (
	HandlerFunc    = contract.HandlerFunc
	MiddlewareFunc = contract.MiddlewareFunc
	Context        = contract.Context
	Router         = contract.Router
	Server         = contract.Server
)
