package router

import (
	"github.com/gorilla/mux"
	"github.com/parkn-co/parkn-server/src/controllers"
	"github.com/parkn-co/parkn-server/src/datastore"
	"github.com/parkn-co/parkn-server/src/utilities/router_utils"
)

func setSpotsRoutes(router *mux.Router, ds *datastore.DataStore) {
	controller := controllers.NewSpotsController(ds)
	sub := router.PathPrefix("/spots/").Subrouter()

	sub.Handle("/", routerutils.NewHandler(controller.GetAll).
		UseMiddlewares(controller.RequireAuthentication)).
		Methods("GET")

	sub.Handle("/1", routerutils.NewHandler(controller.GetAll)).
		Methods("GET")
}
