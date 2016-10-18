package types

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/parkn-co/parkn-server/src/config"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var dbname = config.App.Database.MongoDB.Name

// Session is the type we use to store client sessions
type Session struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
	UserRef mgo.DBRef
	User    interface{}
	Token   string
}

// NewSession creates a new session object
func NewSession(userID bson.ObjectId, token string) *Session {
	return &Session{
		UserRef: mgo.DBRef{
			Collection: "Users",
			Id:         userID,
			Database:   dbname,
		},
		Token: token,
		User:  nil,
	}
}

// jwt.Token https://godoc.org/github.com/dgrijalva/jwt-go#Token
// type Token struct {
//     Raw       string                 // The raw token.  Populated when you Parse a token
//     Method    SigningMethod          // The signing method used or to be used
//     Header    map[string]interface{} // The first segment of the token
//     Claims    Claims                 // The second segment of the token
//     Signature string                 // The third segment of the token.  Populated when you Parse a token
//     Valid     bool                   // Is the token valid?  Populated when you Parse/Verify a token
// }

// IsValid checks if a session's token is still valid
func (s *Session) IsValid() bool {
	token, err := s.parseToken()
	if err != nil || !token.Valid {
		return false
	}

	return true
}

func (s *Session) parseToken() (*jwt.Token, error) {
	return jwt.Parse(s.Token, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.App.Security.Secret), nil
	})
}
