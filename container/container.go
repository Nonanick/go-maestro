package container

import (
	"github.com/Nonanick/go-maestro/controller"
	"github.com/Nonanick/go-maestro/proxy"
	"github.com/Nonanick/go-maestro/server/routes"
)

type Container struct {
	routes.Provider

	Name    string
	BaseURL string

	Controllers []controller.Controller

	RequestProxies  []proxy.Request
	ResponseProxies []proxy.Response

	Validations []
}
