package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// config is the type for all api configuration
type config struct {
	port int // server port
}

// application is the type for all data we want to share cross the api
type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
}

// main entery point to the api
func main() {

	var cfg config
	cfg.port = 8081

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	err := app.serve()
	if err != nil {
		log.Fatal(err)
	}

}

// serve start the web server
func (app *application) serve() error {

	app.infoLog.Println("API listening on port", app.config.port)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.config.port),
		Handler: app.routes(),
	}
	return srv.ListenAndServe()

}