package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type MiddlewareFunc func(http.Handler) http.Handler

func (app *application) routes() http.Handler {

	mux := mux.NewRouter()
	mux.HandleFunc("/users/login", app.Login).Methods("GET")
	mux.HandleFunc("/users/login", app.Login).Methods("POST")
	return mux
}
