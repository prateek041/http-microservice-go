package handlers

import (
	"log"
	"net/http"
)

// defining a hello type to attach methods to them.
type Hello struct {
	l *log.Logger
}

// Now we are doing dependency injection.
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	log.Println("this is the hello handler", r.URL.Path)
}
