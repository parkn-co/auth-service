package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/Schema"
	"github.com/gorilla/context"
	"github.com/parkn-co/parkn-server/src/datastore"
	"github.com/parkn-co/parkn-server/src/types"
	"github.com/parkn-co/parkn-server/src/utilities"
)

// baseController is the base controller for all controllers in this api
type baseController struct {
	DataStore *datastore.DataStore
	decoder   *schema.Decoder
}

var decoder = schema.NewDecoder()

func newBaseController(ds *datastore.DataStore) baseController {
	return baseController{ds, decoder}
}

// SendJSON marshals v to a json struct and sends appropriate headers to w
func (c *baseController) SendJSON(w http.ResponseWriter, r *http.Request, v interface{}, code int) {
	w.Header().Add("Content-Type", "application/json")
	b, err := json.Marshal(v)

	if err != nil {
		log.Print(fmt.Sprintf("Error while encoding JSON: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "Internal server error")
	} else {
		w.WriteHeader(code)
		io.WriteString(w, string(b))
	}
}

// SendInternalError reponds to the request with a 500 error
func (c *baseController) SendInternalError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	io.WriteString(w, "Internal server error")
}

func (c *baseController) SendBadRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	io.WriteString(w, "Bad Request")
}

func (c *baseController) SendNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	io.WriteString(w, "Not Found")
}

func (c *baseController) SendUnauthorized(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
	io.WriteString(w, "Unauthorized")
}

// RequireTokenAuthentication is a middleware for requiring authentication
func (c *baseController) RequireTokenAuthentication(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	ds := c.DataStore.NewDataStore()
	defer ds.Close()

	token := r.Header.Get("Authorization")
	session := &types.Session{}
	err := ds.Sessions.GetSessionByToken(token, session)
	if err != nil {
		fmt.Println("Error getting session by token", err)
		c.SendUnauthorized(w, r)
		return
	}

	fmt.Println("Session found: ", session)

	next(w, r)
}

func (c *baseController) RequireAuthentication(h utilities.Handler) utilities.Handler {
	return utilities.NewHandler(func(w http.ResponseWriter, r *http.Request) {
		ds := c.DataStore.NewDataStore()
		defer ds.Close()

		token := r.Header.Get("Authorization")
		session := &types.Session{}
		err := ds.Sessions.GetSessionByToken(token, session)
		if err != nil {
			fmt.Println("Error getting session by token", err)
			c.SendUnauthorized(w, r)
			return
		}

		context.Set(r, "user", session.User.(types.User))

		h.ServeHTTP(w, r)
	})
}
