package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/OmarElouardi99/BOOKS-API/internal/driver"
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
	db       *driver.DB
}

// main entery point to the api
func main() {

	var cfg config
	cfg.port, _ = strconv.Atoi(os.Getenv("API_PORT"))

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	dsn := os.Getenv("DSN")
	db, err := driver.ConnectSQL(dsn)
	if err != nil {
		log.Fatal("Cannot connect to the DB")
	}
	app := &application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
		db:       db,
	}

	err = app.serve()
	if err != nil {
		log.Fatal(err)
	}

}

// serve start the web server
func (app *application) serve() error {

	app.infoLog.Println("API listening on port", app.config.port)
	router := app.routes()

	return http.ListenAndServe(fmt.Sprintf(":%d", app.config.port), router)
}
