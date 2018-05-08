package user

import (
	"github.com/peaklyio/api-server/api"
)

type User struct {
	api.ApiObject
	FirstName    string `bson:"FirstName,omitempty"`
	LastName     string `bson:"LastName,omitempty"`
	EmailAddress string `bson:"EmailAddress,omitempty"`
}
