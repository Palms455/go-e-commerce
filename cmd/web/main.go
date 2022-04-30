package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "0.0.1"
const cssVersion = "1"


func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

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
	app.infoLog.Printf(fmt.Sprintf("Start http server in %s mode on port %d", app.config.env, app.config.port))

	return srv.ListenAndServe()
}

func main(){
	var cfg config
	flag.IntVar(&cfg.port, "port", 5000, "Port to listen on")
	flag.StringVar(&cfg.env, "env", "dev", "Environment: dev or prod")
	flag.StringVar(&cfg.api, "api", "http://localhost:5001", "Api url")

	flag.Parse()

	cfg.stripe.key, _ = os.LookupEnv("STRIPE_KEY")
	cfg.stripe.secret, _ = os.LookupEnv("STRIPE_SECRET")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	tc := make(map[string]*template.Template)

	app := &application{
		config: cfg,
		infoLog: infoLog,
		errorLog: errorLog,
		templateCache: tc,
		version: version,
	}
	err := app.serve()
	if err != nil {
		app.errorLog.Printf("Server error: %s", err)
		return
	}
}