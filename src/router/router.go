package router

import (
	"github.com/gorilla/mux"
	"github.com/parkn-co/parkn-server/src/authentication"
)

// InitRouter returns the router object
func InitRouter() *mux.Router {
	r := mux.NewRouter()
	sub := r.PathPrefix("/api/v1/").Subrouter()

	// Routes handling
	authController := &authentication.Controller{}
	authSubRouter := sub.PathPrefix("/authentication/").Subrouter()
	authSubRouter.HandleFunc("/signin", authController.SignIn).Methods("GET")

	return r
}
