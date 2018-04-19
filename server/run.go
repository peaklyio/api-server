package server

import (
	"fmt"
	"net/http"

	"github.com/kubicorn/kubicorn/pkg/logger"
)

func ListenAndServe(options *ServerOptions) error {
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
