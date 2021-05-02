package routes

import (
	"io"

	"github.com/Nonanick/go-maestro/hook"
	"github.com/Nonanick/go-maestro/packet"
	"github.com/Nonanick/go-maestro/proxy"
	"github.com/Nonanick/go-maestro/validate"
)

type ExecutionPlan struct {
	TransformRequest func(body io.ReadCloser) interface{}
	Schema           interface{}

	ValidateRequest []validate.Request
	RequestProxies  []proxy.Request

	Resolver func(request packet.Request) packet.Response

	ResponseProxies     []proxy.Response
	ResponseValidations []validate.Response

	Hooks []hook.Hook
}
