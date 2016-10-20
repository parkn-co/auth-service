package controllers

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/parkn-co/parkn-server/src/datastore"
	"github.com/parkn-co/parkn-server/src/utilities/router_utils"
)

// Users is the controller for routes dealing with users
type Users struct {
	baseController
}

// NewUsersController returns a new Spots controller
func NewUsersController(ds *datastore.DataStore) *Users {
	return &Users{newBaseController(ds)}
}

// UserProfile is the handler for returning a user
func (c *Users) UserProfile(w http.ResponseWriter, r *http.Request) (int, interface{}) {
	user := context.Get(r, "user")

	return http.StatusOK, routerutils.Response(user)
}
