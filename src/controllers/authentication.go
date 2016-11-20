package controllers

import (
	"fmt"
	"net/http"

	routing "github.com/parkn-co/go-routing"
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
func (c *Authentication) SignUp(w http.ResponseWriter, r *http.Request) (int, interface{}) {
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
		return routing.NotFound()
	}

	errors, ok := user.Validate()
	if !ok {
		return http.StatusBadRequest, routing.FormErrorResponse(errors)
	}

	// check to make sure user doesn't already exist here
	if c.DataStore.Users.UserExistsByEmail(user.Email) {
		errs := "Email is associated with an existing account"

		return http.StatusConflict, routing.ErrorResponse(errs)
	}

	id, err := c.DataStore.Users.CreateUser(&user)
	if err != nil {
		return routing.InternalError()
	}

	token, err := ds.Sessions.NewSession(id)
	if err != nil {
		return routing.InternalError()
	}

	return http.StatusCreated, routing.Response(map[string]interface{}{"token": token})
}

// SignIn is the handler for signing in
func (c *Authentication) SignIn(w http.ResponseWriter, r *http.Request) (int, interface{}) {
	ds := c.DataStore.NewDataStore()
	defer ds.Close()

	err := r.ParseForm()

	loginRequest := &types.LoginRequest{}
	err = c.decoder.Decode(loginRequest, r.PostForm)
	if err != nil {
		return routing.NotFound()
	}

	errs, ok := loginRequest.Validate()
	if !ok {
		return http.StatusBadRequest, routing.FormErrorResponse(errs)
	}

	user := &types.User{}
	err = ds.Users.GetUserByLogin(loginRequest, user)
	// No user was found with that email and password
	if err != nil {
		return routing.NotFound()
	}

	token, err := ds.Sessions.NewSession(user.ID)
	if err != nil {
		return routing.InternalError()
	}

	return http.StatusOK, routing.Response(map[string]interface{}{"token": token})
}

// SignOut destroys a session by the given token
func (c *Authentication) SignOut(res http.ResponseWriter, req *http.Request) (int, interface{}) {
	ds := c.DataStore.NewDataStore()
	defer ds.Close()

	token := req.Header.Get("Authorization")
	if token != "" {
		_ = ds.Sessions.DestroySession(token)
	}

	return http.StatusOK, routing.Response(nil)
}
