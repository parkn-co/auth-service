package router

import (
	"github.com/gorilla/mux"
	"github.com/parkn-co/parkn-server/src/datastore"
	"github.com/urfave/negroni"
)

// InitRouter returns the router object
func InitRouter() *negroni.Negroni {
	r := mux.NewRouter()
	sub := r.PathPrefix("/api/v1/").Subrouter()

	ds := datastore.Connect()

	// Routes handling
	setAuthenticationRoutes(sub, ds)

	n := negroni.Classic()
	n.UseHandler(r)

	return n
}
