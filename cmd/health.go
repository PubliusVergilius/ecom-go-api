package main

import "net/http"

/*
type Controller map[string]Route

type Route struct {
	handle []*http.HandlerFunc
	suburi string
}
*/

type HealthHandler struct{}

func (*HealthHandler) Pattern() string {
	return "/health"
}

// NewEchoHandler builds a new EchoHandler.
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// ServeHTTP handles an HTTP request to the /echo endpoint.
func (*HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is up and running!"))
}
