package types

// LoginRequest is a struct to hold a new login request
type LoginRequest struct {
	Email    string `json:"email" schema:"email"`
	Password string `json:"email" schema:"password"`
}
