package models

import "gopkg.in/mgo.v2/bson"

// User is the user type.
type User struct {
	ID           bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Email        string        `json:"email"`
	FirstName    string        `json:"firstName"`
	LastName     string        `json:"lastName"`
	PasswordHash string
}
