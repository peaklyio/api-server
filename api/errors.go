package api

import "fmt"

type ApiError struct {
	Message string
}

func E(m string, a ...interface{}) error {
	return &ApiError{
		Message: fmt.Sprintf(m, a...),
	}
}

func (a *ApiError) Error() string {
	return a.Message
}
