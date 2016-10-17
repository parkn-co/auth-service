package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// baseController is the base controller for all controllers in this api
type baseController struct {
}

// SendJSON marshals v to a json struct and sends appropriate headers to w
func (c *baseController) SendJSON(w http.ResponseWriter, r *http.Request, v interface{}, code int) {
	w.Header().Add("Content-Type", "application/json")
	b, err := json.Marshal(v)

	if err != nil {
		log.Print(fmt.Sprintf("Error while encoding JSON: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"error": "Internal server error"}`)
	} else {
		w.WriteHeader(code)
		io.WriteString(w, string(b))
	}
}

// SendInternalError reponds to the request with a 500 error
func (c *baseController) SendInternalError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	io.WriteString(w, `{"error": "Internal server error"}`)
}

func (c *baseController) SendBadRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	io.WriteString(w, `{"error": "Bad Request"}`)
}

func (c *baseController) SendNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	io.WriteString(w, `{"error": "Not Found"}`)
}
