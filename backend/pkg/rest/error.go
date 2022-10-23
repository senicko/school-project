package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HttpError struct {
	Status int
	Msg    string
	Cause  error
}

func (e HttpError) Error() string {
	if e.Cause == nil {
		return e.Msg
	}

	return e.Cause.Error()
}

// NewHttpError creates a new http error.
func NewHttpError(err error, status int, msg string) *HttpError {
	return &HttpError{
		Status: status,
		Msg:    msg,
		Cause:  err,
	}
}

// HandleError processes HttpError and sends response to the client.
func HandleError(w http.ResponseWriter, httpError *HttpError) {
	// Log the error for debugging purpouse
	fmt.Printf("error: %v\n", httpError.Cause)

	// Send response to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpError.Status)

	if httpError.Msg != "" {
		res, _ := json.Marshal(map[string]string{
			"message": httpError.Msg,
		})
		w.Write(res)
	}
}
