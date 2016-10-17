package authentication

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/parkn-co/parkn-server/src/config"
	"github.com/parkn-co/parkn-server/src/types"
	"golang.org/x/crypto/bcrypt"
)

const (
	expireOffset    = 3600
	expirationDelta = 72
)

// PasswordHash generates a hash for a given password
func PasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hash), err
}

// Authenticate checks the login credentials
func Authenticate(user *types.User, loginRequest *types.LoginRequest) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginRequest.Password))

	return err == nil && loginRequest.Email == user.Email
}

// GenerateJwt generates a new token for the user
func GenerateJwt(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * time.Duration(expirationDelta)).Unix(),
		"iat": time.Now().Unix(),
		"sub": userID,
	})

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.App.Security.Secret))
	if err != nil {
		return "", err
	}

	return t, nil
}
