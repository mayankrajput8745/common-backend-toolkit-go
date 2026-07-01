package httpserver

import (
	"fmt"

	"github.com/mayankrajput8745/common-backend-toolkit-go/httpserver/frameworks/fasthttp"
)

func InitHTTPServer(framework string, port int32, cfg ServerConfig) (Server, error) {
	switch framework {
	case "fasthttp":
		return fasthttp.New(port, cfg), nil
	default:
		return nil, fmt.Errorf("unsupported framework: %s", framework)
	}
}
