package api

import "gopkg.in/mgo.v2/bson"

type ApiObject struct {
	Type    string        `bson:"Type,omitempty"`
	MongoID bson.ObjectId `bson:"_id,omitempty"`
	Name    string        `bson:"Name,omitempty"`
	Uniq    int           `bson:"Uniq,omitempty"`
}
