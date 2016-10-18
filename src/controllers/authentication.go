package controllers

import (
	"fmt"
	"net/http"

	"github.com/parkn-co/parkn-server/src/datastore"
	"github.com/parkn-co/parkn-server/src/types"
)

// Authentication is the controller for authentication routes
type Authentication struct {
	baseController
}

// NewAuthController returns a new Authentication controller
func NewAuthController(ds *datastore.DataStore) *Authentication {
	return &Authentication{newBaseController(ds)}
}

// SignUp is the handler for signing up from a client
func (c *Authentication) SignUp(w http.ResponseWriter, r *http.Request) {
	ds := c.DataStore.NewDataStore()
	defer ds.Close()

	err := r.ParseForm()
	if err != nil {
		// Handle error
		fmt.Print("Error parsing form: ")
		fmt.Println(err)
	}

	user := types.NewUser{}
	err = c.decoder.Decode(&user, r.PostForm)
	if err != nil {
		c.SendNotFound(w, r)
		return
	}

	errors, ok := user.Validate()
	if !ok {
		c.SendJSON(
			w,
			r,
			map[string]map[string]string{"errors": errors},
			http.StatusBadRequest,
		)

		return
	}

	// check to make sure user doesn't already exist here
	if c.DataStore.Users.UserExistsByEmail(user.Email) {
		c.SendJSON(
			w,
			r,
			map[string]string{"error": "Email is associated with an existing account"},
			http.StatusBadRequest,
		)

		return
	}

	id, err := c.DataStore.Users.CreateUser(&user)
	if err != nil {
		c.SendInternalError(w, r)
		return
	}

	token, err := ds.Sessions.NewSession(id)
	if err != nil {
		c.SendInternalError(w, r)
		return
	}

	c.SendJSON(
		w,
		r,
		map[string]interface{}{"success": true, "token": token},
		http.StatusOK,
	)
}

// SignIn is the handler for signing in
func (c *Authentication) SignIn(w http.ResponseWriter, r *http.Request) {
	ds := c.DataStore.NewDataStore()
	defer ds.Close()

	err := r.ParseForm()

	loginRequest := &types.LoginRequest{}
	err = c.decoder.Decode(loginRequest, r.PostForm)
	if err != nil {
		c.SendNotFound(w, r)
		return
	}

	if loginRequest.Email == "" || loginRequest.Password == "" {
		c.SendJSON(
			w,
			r,
			map[string]string{"error": "User name and password is required"},
			http.StatusBadRequest,
		)
		return
	}

	user := &types.User{}
	err = ds.Users.GetUserByLogin(loginRequest, user)
	// No user was found with that email and password
	if err != nil {
		c.SendNotFound(w, r)
		return
	}

	token, err := ds.Sessions.NewSession(user.ID)
	if err != nil {
		c.SendInternalError(w, r)
		return
	}

	c.SendJSON(
		w,
		r,
		map[string]interface{}{"success": true, "token": token},
		http.StatusOK,
	)
}
