package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type EasyStatus struct {
	encoder ResponseEncoder
}

type ResponseEncoder interface {
	Encode(interface{}) ([]byte, error)
	Decode([]byte, interface{}) error
}

type JSONEncoder struct {
}

func (j *JSONEncoder) Encode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (j *JSONEncoder) Decode(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func NewEasyStatus(encoder ResponseEncoder) *EasyStatus {
	return &EasyStatus{
		encoder: encoder,
	}
}

func (e *EasyStatus) Status200Okay(w http.ResponseWriter, r *http.Request, v interface{}) {
	bytes, err := e.encoder.Encode(v)
	if err != nil {
		e.Status500InternalServerError(w, r, v)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func (e *EasyStatus) Status400BadRequest(w http.ResponseWriter, r *http.Request, v interface{}) {
	bytes, err := e.encoder.Encode(v)
	if err != nil {
		e.Status500InternalServerError(w, r, v)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write(bytes)
}

func (e *EasyStatus) Status404NotFound(w http.ResponseWriter, r *http.Request, v interface{}) {
	bytes, err := e.encoder.Encode(v)
	if err != nil {
		e.Status500InternalServerError(w, r, v)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write(bytes)
}

func (e *EasyStatus) Status405MethodNotAllowed(w http.ResponseWriter, r *http.Request, v interface{}) {
	bytes, err := e.encoder.Encode(v)
	if err != nil {
		e.Status500InternalServerError(w, r, v)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write(bytes)
}

func (e *EasyStatus) Status500InternalServerError(w http.ResponseWriter, r *http.Request, v interface{}) {
	bytes, err := e.encoder.Encode(v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// We are failing recursively on encoding - default to JSON and hope for the best
		w.Write([]byte(fmt.Sprintf(`
{
	"Error": "Major error while encoding the response: %v"

}
`, err)))
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(bytes)
}
