package routes

import (
	"io"

	"github.com/Nonanick/go-maestro/proxy"
	"github.com/Nonanick/go-maestro/server/packet"
)

type Resolver struct {
	RequestProxies  []proxy.Request
	ResponseProxies []proxy.Response

	Resolver func(request packet.Request) packet.Response

	CastBody    func(body io.ReadCloser) interface{}
	Validations []func() (bool, error)

	Schema interface{}
}
