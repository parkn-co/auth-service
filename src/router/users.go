package router

import (
	"github.com/gorilla/mux"
	"github.com/parkn-co/parkn-server/src/controllers"
	"github.com/parkn-co/parkn-server/src/datastore"
	"github.com/parkn-co/parkn-server/src/utilities"
)

func setUsersRoutes(router *mux.Router, ds *datastore.DataStore) {
	controller := controllers.NewUsersController(ds)
	sub := router.PathPrefix("/users/").Subrouter()

	sub.Handle("/", utilities.
		ApplyMiddlewares(controller.UserProfile, controller.RequireAuthentication)).
		Methods("GET")
}
