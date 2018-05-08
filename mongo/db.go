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
	Database string
}

type Mongo struct {
	Session *mgo.Session
	options *MongoOptions
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
		options: options,
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

func GetCollection(name string) *mgo.Collection {
	return m.Session.DB(m.options.Database).C(name)
}

type u struct {
	Uniq string
}

func (m *Mongo) Save(domain, namespace, uniq string, object interface{}) error {
	db := m.Session.DB(domain)
	collection := db.C(namespace)
	q := u{
		Uniq: uniq,
	}
	_, err := collection.Upsert(q, bson.M{"$set": object})
	if err != nil {
		return fmt.Errorf("Unable to upsert into mongo: %v", err)
	}
	return nil
}

func (m *Mongo) Get(domain, namespace, uniq string, new interface{}) (interface{}, error) {
	q := u{
		Uniq: uniq,
	}
	fmt.Printf("%+v\n", q)
	db := m.Session.DB(domain)
	collection := db.C(namespace)
	err := collection.Find(q).One(&new)
	if err != nil {
		return nil, fmt.Errorf("Unable to read from mongo: %v", err)
	}
	return new, nil
}
