package router

import (
	"github.com/gorilla/mux"
	routing "github.com/parkn-co/go-routing"
	"github.com/parkn-co/parkn-server/src/controllers"
	"github.com/parkn-co/parkn-server/src/datastore"
)

func setSpotsRoutes(router *mux.Router, ds *datastore.DataStore) {
	controller := controllers.NewSpotsController(ds)
	sub := router.PathPrefix("/spots/").Subrouter()

	sub.Handle("/", routing.NewHandler(controller.GetAll).
		UseMiddlewares(controller.RequireAuthentication)).
		Methods("GET")

	sub.Handle("/1", routing.NewHandler(controller.GetAll)).
		Methods("GET")
}
