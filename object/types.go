package object

import (
	"gopkg.in/mgo.v2/bson"
)

type Object struct {
	Type string
	ID   bson.ObjectId
}
