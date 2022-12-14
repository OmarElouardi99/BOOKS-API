package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/OmarElouardi99/BOOKS-API/internal/data"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func (app *application) routes() http.Handler {

	mux := mux.NewRouter()

	headersOk := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Authorization", "X-CSRF-Token"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	fmt.Println(os.Getenv("ORIGIN_ALLOWED"))
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	headersExpose := handlers.ExposedHeaders([]string{"Links"})
	credentialsOk := handlers.AllowCredentials()
	maxAge := handlers.MaxAge(300)
	mux.Use(handlers.CORS(originsOk, headersOk, methodsOk, headersExpose, credentialsOk, maxAge))

	mux.HandleFunc("/users/login", app.Login).Methods("GET")
	mux.HandleFunc("/users/login", app.Login).Methods("POST")
	mux.HandleFunc("/users/all", func(w http.ResponseWriter, r *http.Request) {
		var user data.User
		var err error
		users, err := user.GetAll()
		if err != nil {
			app.errorLog.Println("Can't get users", err)
			return
		}
		app.writeJson(w, http.StatusOK, users)

	})
	return mux
}
