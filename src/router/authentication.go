package router

import (
	"github.com/gorilla/mux"
	"github.com/parkn-co/parkn-server/src/controllers"
	"github.com/parkn-co/parkn-server/src/datastore"
)

func setAuthenticationRoutes(router *mux.Router, ds *datastore.DataStore) {
	authController := controllers.NewAuthController(ds)
	sub := router.PathPrefix("/auth/").Subrouter()

	sub.HandleFunc("/signup", authController.SignUp).Methods("POST")
	sub.HandleFunc("/signin", authController.SignIn).Methods("POST")
}
