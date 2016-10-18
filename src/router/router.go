package router

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/parkn-co/parkn-server/src/datastore"
)

// InitRouter returns the router object
func InitRouter() http.Handler {
	r := mux.NewRouter()
	sub := r.PathPrefix("/api/v1/").Subrouter()

	ds := datastore.Connect()

	// Routes handling
	setAuthenticationRoutes(sub, ds)
	setUsersRoutes(sub, ds)
	setSpotsRoutes(sub, ds)

	routerWithMiddlewares := handlers.LoggingHandler(os.Stdout, r)
	routerWithMiddlewares = handlers.RecoveryHandler()(routerWithMiddlewares)

	return routerWithMiddlewares
}
