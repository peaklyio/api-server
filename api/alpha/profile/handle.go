package profile

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kubicorn/kubicorn/pkg/logger"
	"gopkg.in/mgo.v2/bson"
)

const Namespace = "profile"

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	logger.Debug("/profile [%s]", r.Method)
	switch r.Method {
	case "GET":
		p, err := getProfile(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("400 - Bad request: %v\n", err)))
			return
		}
		pp, err := GetProfile(p)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("400 - Bad request: %v\n", err)))
			return
		}
		bytes, err := json.Marshal(pp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("500 - Internal server error: %v\n", err)))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(fmt.Sprintf("405 - Method [%s] not allowed\n", r.Method)))
		return
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 - Not found\n"))
}

func getProfile(r *http.Request) (*Profile, error) {
	p := &Profile{}
	values := r.URL.Query()
	p.FirstName = values.Get("FirstName")
	p.LastName = values.Get("LastName")
	p.ID = bson.ObjectIdHex(values.Get("ID"))
	return p, nil
}

func getPostProfile(r *http.Request) (*Profile, error) {
	p := &Profile{}
	logger.Debug("Parsing body...")
	d := json.NewDecoder(r.Body)
	err := d.Decode(p)
	if err != nil {
		return nil, fmt.Errorf("Unable to decode. Invalid JSON: %v", err)
	}
	return p, nil
}
