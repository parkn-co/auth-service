package mongo

import mgo "gopkg.in/mgo.v2"

type baseCollection struct {
	collection *mgo.Collection
}
