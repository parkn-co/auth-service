package utilities

import "net/http"

// Middleware is a helper type for middlewares
type middleware func(Handler) Handler

// ApplyMiddlewares applies one or more middlewares for a HandlerFunc
func ApplyMiddlewares(handlerFunc func(http.ResponseWriter, *http.Request), middlewares ...middleware) Handler {
	h := NewHandler(handlerFunc)

	for _, mw := range middlewares {
		h = mw(h)
	}

	return h
}
