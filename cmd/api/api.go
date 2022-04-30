package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
	"webapp/internal/driver"
)
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

const version = "0.0.1"

type config struct{
	port int
	env string
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
	version string
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", app.config.port),
		Handler: app.Routes(),
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

	app := &application{
		config: cfg,
		infoLog: infoLog,
		errorLog: errorLog,
		version: version,
	}

	err = app.serve()
	if err != nil {
		app.errorLog.Printf("error in backend: %s", err)
		os.Exit(1)
	}


}
