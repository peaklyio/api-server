package mongo

import (
	"fmt"

	"github.com/kubicorn/kubicorn/pkg/logger"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoOptions struct {
	Address string
	Port    int
}

type Mongo struct {
	Session *mgo.Session
}

var (
	m *Mongo
)

func NewConnection(options *MongoOptions) error {
	addr := fmt.Sprintf("%s:%d", options.Address, options.Port)
	logger.Info("Mongo address: %v", addr)
	session, err := mgo.Dial(addr)
	if err != nil {
		return fmt.Errorf("Unable to connect to Mongo server: %v", err)
	}
	m = &Mongo{
		Session: session,
	}
	logger.Info("Connected to Mongo!")
	return nil
}

func GetMongo() *Mongo {
	if m == nil {
		logger.Critical("Unable to get mongo, connection not established!")
	}
	return m
}

func (m *Mongo) Save(domain, namespace, uniq string, object interface{}) error {
	db := m.Session.DB(domain)
	collection := db.C(namespace)
	q := struct {
		Uniq string
	}{
		Uniq: uniq,
	}
	_, err := collection.Upsert(q, bson.M{"$set": object})
	if err != nil {
		return fmt.Errorf("Unable to upsert into mongo: %v", err)
	}
	return nil
}

func (m *Mongo) Get(domain, namespace string, query, new interface{}) (interface{}, error) {
	db := m.Session.DB(domain)
	collection := db.C(namespace)
	err := collection.Find(query).One(&new)
	if err != nil {
		return nil, fmt.Errorf("Unable to read from mongo: %v", err)
	}
	return new, nil
}
