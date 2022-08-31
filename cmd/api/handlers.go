package main

import (
	"encoding/json"
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

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		// send back error message
		app.errorLog.Println("Invalid Json")
		payload.Error = true
		payload.Mesage = "invalid json"
		out, err := json.MarshalIndent(payload, "", "\t")
		if err != nil {
			app.errorLog.Println(err)
		}

		w.Header().Set("Content type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(out)
		return
	}

	// authenticate
	app.infoLog.Println(creds.Password, creds.UserName)

	// send back a response
	payload.Error = false
	payload.Mesage = "signed in"
	out, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		app.errorLog.Println(err)
	}

	w.Header().Set("Content type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}
