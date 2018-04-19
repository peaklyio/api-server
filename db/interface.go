package db

type Storable interface {
	Save(domain, namespace string, object interface{}) error
	Get(domain, namespace string, query, new interface{}) (interface{}, error)
}
