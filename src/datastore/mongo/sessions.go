package mongo

import (
	"github.com/parkn-co/parkn-server/src/services/authentication"
	"github.com/parkn-co/parkn-server/src/types"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// SessionsCollection is the struct we use to interact with the sessions collection
type SessionsCollection struct {
	baseCollection
}

// NewSessionsCollection creates a new SessionsCollection
func NewSessionsCollection(collection *mgo.Collection) *SessionsCollection {
	return &SessionsCollection{baseCollection{collection}}
}

// NewSession creates a new session in the database
func (sessions *SessionsCollection) NewSession(userID bson.ObjectId) (string, error) {
	token, err := authentication.GenerateJwt(userID.String())
	if err != nil {
		return "", err
	}

	err = sessions.collection.Insert(types.NewSession(userID, token))
	if err != nil {
		return "", err
	}

	return token, nil
}
