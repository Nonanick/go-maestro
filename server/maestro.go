package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Nonanick/go-maestro/pipes"
	"github.com/Nonanick/go-maestro/routes"
)

type Maestro struct {
	Address string
	Port    uint16
	Name    string

	Routes  []routes.Provider
	OnError []routes.Error

	RequestPipes  []pipes.Request
	ResponsePipes []pipes.Response
}

func (server *Maestro) ServeHTTP(http.ResponseWriter, *http.Request) {
	// Function to match url pattern ?

	// Handle 404 error if no resolver found

	// Apply pipes to request

	// Create a context for the current request

	// Generate custom request struct

	// Apply requets proxies

	// Invoke request resolver and wait for a response

	// Check for resolve errors

	// Apply response proxies

	// Apply response pipes

	// End

}

func (server *Maestro) Listen() (err error) {
	fmt.Println("[Maestro] Starting HTTP server on " + server.Address + ":" + strconv.Itoa(int(server.Port)))

	return http.ListenAndServe(
		server.Address+":"+strconv.FormatInt(int64(server.Port), 10),
		server,
	)
}
