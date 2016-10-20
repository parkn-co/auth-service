package router

import (
	"io"
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

	sub.HandleFunc("/", rootHandler).Methods("GET")

	routerWithMiddlewares := handlers.LoggingHandler(os.Stdout, r)
	routerWithMiddlewares = handlers.RecoveryHandler()(routerWithMiddlewares)

	return routerWithMiddlewares
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Parkn API V1")
}
