package router

import (
	"github.com/gorilla/mux"
	routing "github.com/parkn-co/go-routing"
	"github.com/parkn-co/parkn-server/src/controllers"
	"github.com/parkn-co/parkn-server/src/datastore"
)

func setAuthenticationRoutes(router *mux.Router, ds *datastore.DataStore) {
	authController := controllers.NewAuthController(ds)
	sub := router.PathPrefix("/auth/").Subrouter()

	sub.Handle("/signup/", routing.NewHandler(authController.SignUp)).
		Methods("POST")

	sub.Handle("/signin/", routing.NewHandler(authController.SignIn)).
		Methods("POST")

	sub.Handle("/signout/", routing.NewHandler(authController.SignOut)).
		Methods("POST")
}
