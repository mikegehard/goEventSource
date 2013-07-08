package eventSource

import (
	"net/http"
)

type Conn struct {
	response http.ResponseWriter
}

func (c Conn) Write (out []byte) {
	c.response.Write(out)
}

type Handler func(*Conn)

// ServeHTTP implements the http.Handler interface for a WebSocket
func (h Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h(&Conn{w})
}
