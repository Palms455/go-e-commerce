package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"flag"
	"time"
)
const version = "0.0.1"
const cssVersion = "1"

type config struct{
	port int
	env, api string
	db struct{
		dsn string
	}
	stripe struct{
		secret, key string
	}
}

type application struct{
	config config
	infoLog *log.Logger
	errorLog *log.Logger
	templateCache map[string]*template.Template
	version string
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", app.config.port),
		Handler: app.routes(),
		IdleTimeout: 30 * time.Second,
		ReadTimeout: 10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	app.infoLog.Printf(fmt.Sprintf("Start backend server in %s mode on port %d", app.config.env, app.config.port))

	return srv.ListenAndServe()
}


func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 5001, "Port to listen on")
	flag.StringVar(&cfg.env, "env", "dev", "Environment: dev or prod")

	flag.Parse()

	cfg.stripe.key = os.Getenv("STRIPE_KEY")
	cfg.stripe.secret = os.Getenv("STRIPE_SECRET")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		config: cfg,
		infoLog: infoLog,
		errorLog: errorLog,
	}

	err := app.serve()
	if err != nil {
		app.errorLog.Printf("error in backend: %s", err)
		os.Exit(1)
	}


}
