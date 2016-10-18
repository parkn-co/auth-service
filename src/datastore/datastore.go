package datastore

import (
	"github.com/parkn-co/parkn-server/src/config"
	"github.com/parkn-co/parkn-server/src/datastore/mongo"
	mgo "gopkg.in/mgo.v2"
)

var dbname = config.App.Database.MongoDB.Name

// DataStore is the type for a database session
type DataStore struct {
	mongoSession *mgo.Session
	Users        *mongo.UsersCollection
	Sessions     *mongo.SessionsCollection
}

func initialDataStore(mongoSession *mgo.Session) *DataStore {
	return &DataStore{
		mongoSession: mongoSession,
		Users:        nil,
		Sessions:     nil,
	}
}

// NewDataStore returns a new datastore with a copied session
func (ds *DataStore) NewDataStore() *DataStore {
	session := ds.mongoSession.Copy()

	return &DataStore{
		mongoSession: session,
		Users:        mongo.NewUsersCollection(collectionFromSession(session, "Users")),
		Sessions:     mongo.NewSessionsCollection(collectionFromSession(session, "Sessions"), session),
	}
}

// Close the session
func (ds *DataStore) Close() {
	ds.mongoSession.Close()
}

func collectionFromSession(session *mgo.Session, name string) *mgo.Collection {
	return session.DB(dbname).C(name)
}
