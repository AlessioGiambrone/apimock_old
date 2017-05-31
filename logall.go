package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

// EchoToStdout is a middleware handler that redirects any received request to stdout.
type EchoToStdout struct{}

// newEchoer returns a new EchoToStdout instance
func newEchoer() *EchoToStdout {
	return new(EchoToStdout)
}

func (l *EchoToStdout) ServeHTTP(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	body, err := httputil.DumpRequest(req, true)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	next(rw, req)
}
