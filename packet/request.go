package packet

import (
	"net/http"
)

type Request struct {
	http.Request
}
