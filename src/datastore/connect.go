package datastore

import (
	"fmt"

	"github.com/parkn-co/parkn-server/src/config"
	mgo "gopkg.in/mgo.v2"
)

// Connect connects to the database
func Connect() *DataStore {
	mongoSession, err := mgo.Dial(config.App.Database.MongoDB.URL)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB")

	ds := initialDataStore(mongoSession)

	return ds.NewDataStore()
}
