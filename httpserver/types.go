package httpserver

import "common-backend-toolkit/httpserver/contract"

type (
	HandlerFunc    = contract.HandlerFunc
	MiddlewareFunc = contract.MiddlewareFunc
	Context        = contract.Context
	Router         = contract.Router
	Server         = contract.Server
)
