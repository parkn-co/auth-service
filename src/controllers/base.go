package controllers

import (
	"net/http"

	"github.com/gorilla/Schema"
	"github.com/gorilla/context"
	routing "github.com/parkn-co/go-routing"
	"github.com/parkn-co/parkn-server/src/datastore"
	"github.com/parkn-co/parkn-server/src/types"
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

// Definitions for global middlewares

// RequireAuthentication is a middleware for requiring authentication
func (c *baseController) RequireAuthentication(res http.ResponseWriter, req *http.Request) (int, interface{}) {
	ds := c.DataStore.NewDataStore()
	defer ds.Close()

	token := req.Header.Get("Authorization")
	session := &types.Session{}
	err := ds.Sessions.GetSessionByToken(token, session)
	if err != nil {
		return routing.Unauthorized()
	}

	context.Set(req, "user", session.User.(types.User))

	return 0, nil
}
