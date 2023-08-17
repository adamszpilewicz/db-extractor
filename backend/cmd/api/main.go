// main.go
package main

import (
	"fmt"
	"log"
	"os"
)

// config is the type for all application configuration
type config struct {
	port     int // what port do we want the web server to listen on
	username string
	password string
	database string
	host     string
}

// application is the type for all data we want to share with the
// various parts of our application.
type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
}

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

	r := newRouter(app)
	r.setupRoutes()

	err := r.start(fmt.Sprintf(":%d", app.config.port))
	if err != nil {
		log.Fatal(err)
	}
}
