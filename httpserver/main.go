package httpserver

import (
	"fmt"

	"common-backend-toolkit/httpserver/frameworks/fasthttp"
)

func InitHTTPServer(framework string, port int32) (Server, error) {
	switch framework {
	case "fasthttp":
		return fasthttp.New(port), nil
	default:
		return nil, fmt.Errorf("unsupported framework: %s", framework)
	}
}
