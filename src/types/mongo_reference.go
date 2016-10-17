package types

import "gopkg.in/mgo.v2/bson"

// mongoReference is a struct for mongo referencts.
// For more info, see https://docs.mongodb.com/v3.2/reference/database-references/
// In most cases, we should do our references in code.
// As of right now, this is only used for quickly getting the user in a session.
type mongoReference struct {
	// Holds the collection name
	Ref string `bson:"$ref"`
	// Is the _id of the object we are refering to
	ID bson.ObjectId `bson:"$id"`
}
