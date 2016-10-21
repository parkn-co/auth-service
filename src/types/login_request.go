package types

// LoginRequest is a struct to hold a new login request
type LoginRequest struct {
	Email    string `json:"email" schema:"email"`
	Password string `json:"email" schema:"password"`
}

// Validate is used for validating a new user when signing up
func (u *LoginRequest) Validate() (map[string]string, bool) {
	errors := make(map[string]string)

	if u.Email == "" {
		errors["email"] = "Email is required"
	}

	if u.Password == "" {
		errors["password"] = "Password is required"
	}

	return errors, len(errors) == 0
}
