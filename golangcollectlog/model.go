package main

import "net/http"

type LoggingResponseWriter struct {
	http.ResponseWriter
	StatusCode int
	BytesSent int
}