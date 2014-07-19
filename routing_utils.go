package main

import (
	"encoding/json"
	"net/http"
)

func Render(w http.ResponseWriter, object interface{}, status int) {
	json, err := json.Marshal(object)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	w.Write(json)
}

func Render400(w http.ResponseWriter) {
	Render(
		w,
		&Error{
			Status:  400,
			Message: "bad request, check json format",
		},
		404,
	)
}

func Render404(w http.ResponseWriter, req *http.Request) {
	Render(
		w,
		&Error{
			Status:  404,
			Message: "not found",
		},
		404,
	)
}

func Render405(w http.ResponseWriter) {
	Render(
		w,
		&Error{
			Status:  405,
			Message: "method not allowed",
		},
		405,
	)
}
