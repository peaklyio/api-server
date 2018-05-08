package api

import (
	"net/http"

	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/mitchellh/mapstructure"
)

func RequestToObject(r *http.Request, v interface{}) error {
	hashMap := r.URL.Query()
	if len(hashMap) < 1 {
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return fmt.Errorf("unable to read body after empty query values: %v", err)
		}
		err = json.Unmarshal(bytes, v)
		if err != nil {
			return fmt.Errorf("unable to unmarshal request - invalid encoding: %v", err)
		}
		return nil
	}

	// Todo - sanitize data here

	queryMap := make(map[string]string)
	for k, v := range hashMap {
		queryMap[k] = v[0]
	}
	return mapstructure.Decode(queryMap, v)
}
