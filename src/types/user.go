package types

import (
	"errors"

	"gopkg.in/mgo.v2/bson"
)

// User is the user type.
type User struct {
	ID           bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Email        string        `json:"email" schema:"email"`
	FirstName    string        `json:"firstName" schema:"firstName"`
	LastName     string        `json:"lastName" shema:"lastName"`
	PasswordHash string
}

// Validate is used to validate a user
func (u *User) Validate() error {
	if u.Email == "" {
		return errors.New("Email is required")
	}

	if u.FirstName == "" {
		return errors.New("First name is required")
	}

	if u.LastName == "" {
		return errors.New("Last name is required")
	}

	return nil
}
