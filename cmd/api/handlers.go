package main

import (
	"net/http"
)

type jsonResponse struct {
	Error  bool   `json:"error"`
	Mesage string `json:"message"`
}

func (app *application) Login(w http.ResponseWriter, r *http.Request) {

	type credentials struct {
		UserName string `json:"email"`
		Password string `json:"password"`
	}

	var creds credentials
	var payload jsonResponse

	err := app.readJson(w, r, &creds)
	if err != nil {
		app.errorLog.Println(err)
		payload.Error = true
		payload.Mesage = "invalid json"
		_ = app.writeJson(w, http.StatusBadRequest, payload)
	}
	// authenticate
	app.infoLog.Println(creds.Password, creds.UserName)

	// send back a response
	payload.Error = false
	payload.Mesage = "signed in"
	err = app.writeJson(w, http.StatusOK, payload)
	if err != nil {
		app.errorLog.Println(err)
	}
}
