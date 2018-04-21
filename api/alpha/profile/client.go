package profile

import (
	"github.com/peaklyio/api-server/db"
)

func GetProfile(query *Profile) (*Profile, error) {
	dbi := db.Get()
	p := &Profile{}
	dbi.Get(db.Domain, Namespace, query, p)
	return p, nil
}

func SaveProfile(newProfile *Profile) (*Profile, error) {
	newProfile.Uniq = newProfile.EmailAddress
	dbi := db.Get()
	err := dbi.Save(db.Domain, Namespace, newProfile.Uniq, newProfile)
	if err != nil {
		return nil, err
	}
	return newProfile, nil
}
