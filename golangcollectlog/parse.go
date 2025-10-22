package main

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// --- Các phương thức (methods) cho struct LoggingResponseWriter ---
// (Struct đã được định nghĩa ở model.go)

func NewLoggingResponseWriter(w http.ResponseWriter) *LoggingResponseWriter {
	// Tạo struct (định nghĩa ở model.go)
	return &LoggingResponseWriter{w, http.StatusOK, 0}
}

func (lrw *LoggingResponseWriter) WriteHeader(code int) {
	lrw.StatusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func (lrw *LoggingResponseWriter) Write(b []byte) (int, error) {
	size, err := lrw.ResponseWriter.Write(b)
	lrw.BytesSent += size
	return size, err
}

// --- Middleware ghi log (Làm nhiệm vụ của Logstash) ---

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		lrw := NewLoggingResponseWriter(w)
		next.ServeHTTP(lrw, r)

		duration := time.Since(startTime)

		log.WithFields(log.Fields{
			"http.request.method":       r.Method,
			"url.path":                  r.URL.Path,
			"http.response.status_code": lrw.StatusCode,
			"http.response.bytes":       lrw.BytesSent,
			"http.request.referrer":     r.Referer(),
			"user_agent.original":       r.UserAgent(),	
			"client.ip":                 r.RemoteAddr,
			"event.duration":            duration.Nanoseconds(),
		}).Info("http_access_log")
	})
}
