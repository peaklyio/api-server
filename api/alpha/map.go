package alpha

import (
	"net/http"

	"github.com/kubicorn/kubicorn/pkg/logger"
	"github.com/peaklyio/api-server/api/alpha/user"
)

var (
	routes = map[string]http.HandlerFunc{
		"/":     RootHandler,
		"/user": user.UserHandler,
		//"/profile": profile.ProfileHandler,
	}
)

func RegisterHandler() {
	for endpoint, f := range routes {
		logger.Info("Loading endpoint [%s]", endpoint)
		http.HandleFunc(endpoint, f)
	}
}

func nothingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("nothing..."))
}
