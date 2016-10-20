package controllers

import (
	"net/http"

	"github.com/parkn-co/parkn-server/src/datastore"
	"github.com/parkn-co/parkn-server/src/utilities/router_utils"
)

// Spots is the controller for routes dealing with parking spots
type Spots struct {
	baseController
}

// NewSpotsController returns a new Spots controller
func NewSpotsController(ds *datastore.DataStore) *Spots {
	return &Spots{newBaseController(ds)}
}

// GetAll is the handler for signing up from a client
func (c *Spots) GetAll(w http.ResponseWriter, r *http.Request) (int, interface{}) {
	return http.StatusOK, routerutils.Response("We are authenticated!")
}
