package api

import (
	"encoding/json"
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

func NewHttpError(err error, status int, msg string) *HttpError {
	return &HttpError{
		Status: status,
		Msg:    msg,
		Cause:  nil,
	}
}

func HandleError(w http.ResponseWriter, httpErr *HttpError) {
	res, _ := json.Marshal(map[string]string{
		"message": httpErr.Msg,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpErr.Status)
	w.Write(res)
}
