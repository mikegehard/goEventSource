package main

import (
	"fmt"
	"net/http"
	"github.com/msgehard/eventSource"
)

func eventHandler(es *eventSource.Conn) {
	es.Write([]byte("Hello world"))
}

func main() {
	http.Handle("/events", eventSource.Handler(eventHandler))
	http.ListenAndServe(":9001", nil)
	fmt.Printf("Hello world!")
}
