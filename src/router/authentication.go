package router

import (
	"github.com/gorilla/mux"
	"github.com/parkn-co/parkn-server/src/controllers"
	"github.com/parkn-co/parkn-server/src/datastore"
)

func setAuthenticationRoutes(router *mux.Router, ds *datastore.DataStore) {
	authController := &controllers.Authentication{DataStore: ds}
	authSubRouter := router.PathPrefix("/auth/").Subrouter()
	authSubRouter.HandleFunc("/signup", authController.SignUp).Methods("POST")
	authSubRouter.HandleFunc("/signin", authController.SignIn).Methods("POST")
}
