package router

import (
	"github.com/gorilla/mux"
	routing "github.com/parkn-co/go-routing"
	"github.com/parkn-co/parkn-server/src/controllers"
	"github.com/parkn-co/parkn-server/src/datastore"
)

func setUsersRoutes(router *mux.Router, ds *datastore.DataStore) {
	controller := controllers.NewUsersController(ds)
	sub := router.PathPrefix("/users/").Subrouter()

	sub.Handle("/", routing.NewHandler(controller.UserProfile).
		UseMiddlewares(controller.RequireAuthentication)).
		Methods("GET")
}
