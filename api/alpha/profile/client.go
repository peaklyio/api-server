package profile

import (
	"github.com/peaklyio/api-server/db"
	"gopkg.in/mgo.v2/bson"
)

func GetProfile(query *Profile) (*Profile, error) {
	dbi := db.Get()
	p := &Profile{}
	dbi.Get(db.Domain, Namespace, query, p)
	return p, nil
}

func SaveProfile(newProfile *Profile) (*Profile, error) {
	dbi := db.Get()
	id, err := dbi.Save(db.Domain, Namespace, newProfile)
	if err != nil {
		return nil, err
	}
	newProfile.ID = bson.ObjectIdHex(id)
	return newProfile, nil
}
