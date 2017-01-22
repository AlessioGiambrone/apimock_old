package main

import (
	"fmt"
	"net/http"
)

// HeaderLogger is a middleware handler that logs the request headers.
type HeaderLogger struct{}

// newHeaderLogger returns a new Logger instance
func newHeaderLogger() *HeaderLogger {
	return &HeaderLogger{}
}

func (l *HeaderLogger) ServeHTTP(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {

	for k, v := range req.Header {
		fmt.Println(k, v)
	}

	fmt.Fprintln(rw, "Chiamata registrata.")

}
