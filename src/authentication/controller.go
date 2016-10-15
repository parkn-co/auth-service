package authentication

import (
	"net/http"

	"github.com/parkn-co/parkn-server/src/common"
)

// Controller is the controller for authentication routes
type Controller struct {
	common.Controller
}

// SignIn is the handler for signing in from a client
func (c *Controller) SignIn(w http.ResponseWriter, r *http.Request) {
	c.SendJSON(
		w,
		r,
		map[string]string{"hello": "world"},
		http.StatusOK,
	)
}
