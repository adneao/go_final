package service

import (
	"encoding/json"
	"net/http"
)

func Make(w http.ResponseWriter, data interface{}, err error) {
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		makeResponseError(w, err)
		return
	}

	if data == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	responseContent, err := json.Marshal(data)
	if err != nil {
		makeResponseError(w, err)
		return
	}

	w.Write(responseContent)
}

func makeResponseError(w http.ResponseWriter, err error) {
	resp := ErrorWriter{err: err}
	resp.write(w)
}

type ErrorWriter struct {
	err error
}

func (e ErrorWriter) write(w http.ResponseWriter) {
	if e.err == nil {
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	response, err := json.Marshal(ErrorResponse{Message: e.err.Error()})
	if err != nil {
		return
	}

	w.Write(response)
}

type ErrorResponse struct {
	Message string `json:"message"`
}
