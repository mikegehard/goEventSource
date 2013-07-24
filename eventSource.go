package eventsource

import (
	"fmt"
	"net/http"
)

/*
See :
http://html5doctor.com/server-sent-events/
http://www.html5rocks.com/en/tutorials/eventsource/basics/
http://cjihrig.com/blog/the-server-side-of-server-sent-events/
*/

type Conn struct {
	writer         http.ResponseWriter
	flusher        http.Flusher
	closeNotifiter http.CloseNotifier
}

func (c Conn) Write(message string) {
	c.writer.Write([]byte(fmt.Sprintf("data: %s", message)))
	c.writer.Write([]byte("\n\n"))
	c.flusher.Flush()
}

func (c Conn) CloseNotify() <-chan bool {
	return c.closeNotifiter.CloseNotify()
}

type Handler func(*Conn)

func (h Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	f, f_ok := w.(http.Flusher)

	if !f_ok {
		panic("ResponseWriter is not a Flusher")
	}

	cn, cn_ok := w.(http.CloseNotifier)

	if !cn_ok {
		panic("ResponseWriter is not a CloseNotifier")
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Transfer-Encoding", "chunked")

	f.Flush()

	h(&Conn{w, f, cn})
}
