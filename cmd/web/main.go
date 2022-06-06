package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/joho/godotenv"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
	"webapp/internal/driver"
	"webapp/internal/models"
)

const version = "0.0.1"
const cssVersion = "1"

var session *scs.SessionManager

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

type config struct {
	port     int
	env, api string
	db       struct {
		dsn string
	}
	stripe struct {
		secret, key string
	}
}

type application struct {
	config        config
	infoLog       *log.Logger
	errorLog      *log.Logger
	templateCache map[string]*template.Template
	version       string
	DB            models.DBModel
	Session       *scs.SessionManager
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}
	app.infoLog.Printf(fmt.Sprintf("Start http server in %s mode on port %d", app.config.env, app.config.port))

	return srv.ListenAndServe()
}

func main() {
	gob.Register(map[string]interface{}{})
	var cfg config
	flag.IntVar(&cfg.port, "port", 5000, "Port to listen on")
	flag.StringVar(&cfg.env, "env", "dev", "Environment: dev or prod")
	flag.StringVar(&cfg.api, "api", "http://localhost:5001", "Api url")

	flag.Parse()

	cfg.stripe.key, _ = os.LookupEnv("STRIPE_KEY")
	cfg.stripe.secret, _ = os.LookupEnv("STRIPE_SECRET")
	cfg.db.dsn, _ = os.LookupEnv("DSN")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	conn, err := driver.OpenDB(cfg.db.dsn)
	if err != nil {
		errorLog.Println(err)
	}

	defer conn.Close()

	// set SessionManager
	session = scs.New()
	session.Lifetime = 24 * time.Hour

	tc := make(map[string]*template.Template)

	app := &application{
		config:        cfg,
		infoLog:       infoLog,
		errorLog:      errorLog,
		templateCache: tc,
		version:       version,
		DB:            models.DBModel{DB: conn},
		Session:       session,
	}
	err = app.serve()
	if err != nil {
		app.errorLog.Printf("Server error: %s", err)
		return
	}
}
