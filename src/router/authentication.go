package router

import (
	"github.com/gorilla/mux"
	"github.com/parkn-co/parkn-server/src/controllers"
	"github.com/parkn-co/parkn-server/src/datastore"
	"github.com/parkn-co/parkn-server/src/utilities/router_utils"
)

func setAuthenticationRoutes(router *mux.Router, ds *datastore.DataStore) {
	authController := controllers.NewAuthController(ds)
	sub := router.PathPrefix("/auth/").Subrouter()

	sub.Handle("/signup/", routerutils.NewHandler(authController.SignUp)).
		Methods("POST")

	sub.Handle("/signin/", routerutils.NewHandler(authController.SignIn)).
		Methods("POST")
}
