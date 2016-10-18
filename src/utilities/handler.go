package utilities

import "net/http"

// Handler is used to create handlers for middlewares. It implements the http.Handler interface
type Handler struct {
	serveHTTP func(http.ResponseWriter, *http.Request)
}

// NewHandler creates a new Handler
func NewHandler(handle func(http.ResponseWriter, *http.Request)) Handler {
	return Handler{handle}
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.serveHTTP(w, r)
}
