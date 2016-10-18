package router

import (
	"github.com/gorilla/mux"
	"github.com/parkn-co/parkn-server/src/controllers"
	"github.com/parkn-co/parkn-server/src/datastore"
	"github.com/parkn-co/parkn-server/src/utilities"
)

func setSpotsRoutes(router *mux.Router, ds *datastore.DataStore) {
	controller := controllers.NewSpotsController(ds)
	sub := router.PathPrefix("/spots/").Subrouter()

	sub.Handle("/", utilities.
		ApplyMiddlewares(controller.GetAll, controller.RequireAuthentication)).
		Methods("GET")
}
