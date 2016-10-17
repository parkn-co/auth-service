package types

import "github.com/parkn-co/parkn-server/src/utilities"

// NewUser is used for validating a new user when signing up
type NewUser struct {
	Email     string `schema:"email"`
	FirstName string `schema:"firstName"`
	LastName  string `schema:"lastName"`
	Password  string `schema:"password"`
}

// Validate is used for validating a new user when signing up
func (u *NewUser) Validate() (map[string]string, bool) {
	errors := make(map[string]string)

	if u.Email == "" {
		errors["Email"] = "Email is required"
	} else if !utilities.IsEmailValid(u.Email) {
		errors["Email"] = "Please enter a valid email address"
	}

	if u.FirstName == "" {
		errors["FirstName"] = "First name is required"
	}

	if u.LastName == "" {
		errors["LastName"] = "Last name is required"
	}

	if u.Password == "" {
		errors["Password"] = "Password is required"
	}

	return errors, len(errors) == 0
}
