package main

import (
	"fmt"

	"github.com/Nonanick/go-maestro/server"
)

func main() {

	maestro := server.Maestro{
		Address: "localhost",
		Port:    8081,
		Name:    "MaestroServer",
	}

	serverLaunchErr := maestro.Listen()

	if serverLaunchErr != nil {
		fmt.Println(serverLaunchErr)
	}
}
