package profile

import (
	"github.com/peaklyio/api-server/object"
)

type Profile struct {
	object.Object
	FirstName string
	LastName  string
}
