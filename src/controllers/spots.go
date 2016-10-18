package controllers

import (
	"net/http"

	"github.com/parkn-co/parkn-server/src/datastore"
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
func (c *Spots) GetAll(w http.ResponseWriter, r *http.Request) {
	c.SendJSON(
		w,
		r,
		map[string]interface{}{"success": true, "data": "We are authenticated!"},
		http.StatusOK,
	)
}
