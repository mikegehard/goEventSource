package eventsource

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHttpWritesTheProperHeaders(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/foo", nil)
	if err != nil {
		t.Errorf("Failed to create request")
	}

	recorder := httptest.NewRecorder()

	eventHandler := func(es *Conn) {}

	httpHandler := Handler(eventHandler)

	httpHandler.ServeHTTP(recorder, req)

	if recorder.Header().Get("Content-Type") != "text/event-stream" {
		t.Errorf("Content-Type header not set to text/event-stream")
	}

	if recorder.Header().Get("Cache-Control") != "no-cache" {
		t.Errorf("Cache-Control header not set to no-cache")
	}

	if recorder.Header().Get("Connection") != "keep-alive" {
		t.Errorf("Connection header not set to keep-alive")
	}

	if recorder.Header().Get("Transfer-Encoding") != "chunked" {
		t.Errorf("Transfer-Encoding header not set to chunked")
	}
}

func TestServeHttpSetsResponseToOK(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/foo", nil)
	if err != nil {
		t.Errorf("Failed to create request")
	}

	recorder := httptest.NewRecorder()

	eventHandler := func(es *Conn) {}

	httpHandler := Handler(eventHandler)

	httpHandler.ServeHTTP(recorder, req)

	if recorder.Code != 200 {
		t.Errorf("Response was not a 200")
	}
}

func TestServeHttpFlushes(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/foo", nil)
	if err != nil {
		t.Errorf("Failed to create request")
	}

	recorder := httptest.NewRecorder()

	eventHandler := func(es *Conn) {}

	httpHandler := Handler(eventHandler)

	httpHandler.ServeHTTP(recorder, req)

	if !recorder.Flushed {
		t.Errorf("Not flushed")
	}
}

func TestConnWritingAMessage(t *testing.T) {
	recorder := httptest.NewRecorder()
	connection := &Conn{recorder, recorder}

	connection.Write("Hello World")

	expectedBody := "data: Hello World\n\n"

	switch recorder.Body.String() {
	case expectedBody:
		// body is equal so no need to do anything
	default:
		t.Errorf("Body (%s) did not match expectation (%s).", recorder.Body.String(), expectedBody)
	}

	if !recorder.Flushed {
		t.Errorf("Writer did not get flushed.")
	}
}
