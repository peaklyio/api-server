package alpha

import (
	"net/http"

	"github.com/kubicorn/kubicorn/pkg/logger"
)

var (
	routes = map[string]http.HandlerFunc{
		"/": nothingHandler,
	}
)

func RegisterHandler() {
	for endpoint, f := range routes {
		logger.Info("Loading endpoint [%s]", endpoint)
		http.HandleFunc(endpoint, f)
	}
}

func nothingHandler(w http.ResponseWriter, r *http.Request) {

}
