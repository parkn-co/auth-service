package mongo

import (
	"errors"

	"github.com/parkn-co/parkn-server/src/services/authentication"
	"github.com/parkn-co/parkn-server/src/types"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// SessionsCollection is the struct we use to interact with the sessions collection
type SessionsCollection struct {
	baseCollection
	session *mgo.Session
}

// NewSessionsCollection creates a new SessionsCollection
func NewSessionsCollection(collection *mgo.Collection, session *mgo.Session) *SessionsCollection {
	return &SessionsCollection{baseCollection{collection}, session}
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

// GetSessionByToken gets a session by a token
func (sessions *SessionsCollection) GetSessionByToken(token string, session *types.Session) error {
	err := sessions.collection.Find(bson.M{"token": token}).One(session)
	if err != nil {
		return err
	}

	if !session.IsValid() {
		return errors.New("Token is invalid")
	}

	user := types.User{}
	sessions.session.FindRef(&session.UserRef).One(&user)

	session.User = user

	return nil
}
