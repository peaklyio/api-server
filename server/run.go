package server

import (
	"fmt"
	"net/http"

	"github.com/kubicorn/kubicorn/pkg/logger"
	api "github.com/peaklyio/api-server/api/alpha"
	"github.com/peaklyio/api-server/db"
	"github.com/peaklyio/api-server/db/mongo"
)

func ListenAndServe(options *ServerOptions) error {
	err := mongo.NewConnection(options.DatabaseOptions.MongoOptions)
	if err != nil {
		return fmt.Errorf("Unable to connect to Mongo: %v", err)
	}
	db.Set(mongo.GetMongo(), options.Domain)
	api.RegisterHandler()
	addr := fmt.Sprintf("%s:%d", options.BindAddress, options.BindPort)
	logger.Info("Starting server [%s]", addr)
	if options.EnableTLS {
		logger.Info("Running HTTPs server...")
		return fmt.Errorf("TLS not implemented!")
	} else {
		logger.Info("Running HTTP server...")
		err := http.ListenAndServe(addr, nil)
		if err != nil {
			return fmt.Errorf("Failure during serving: %v", err)
		}
	}
	return nil
}
