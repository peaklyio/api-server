package profile

import (
	"github.com/kubicorn/kubicorn/pkg/logger"
	"github.com/peaklyio/api-server/db"
)

func GetProfile(query *Profile) (*Profile, error) {
	logger.Debug("Lookup in mongo..")
	dbi := db.Get()
	p := &Profile{}
	dbi.Get(db.Domain, Namespace, query, p)
	return p, nil
}
