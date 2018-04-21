package object

import (
	"gopkg.in/mgo.v2/bson"
)

type Object struct {
	Type    string
	MongoID bson.ObjectId `bson:"_id,omitempty"`
	Name    string
	Uniq    string
}
