package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"net/http"
)

func (app *application) Routes() http.Handler {
	mux := chi.NewRouter()

	// Разрешение CORS запросов
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		MaxAge: 300,
	}))

	mux.Post("/api/v1/payment-intent", app.GetPaymentIntent)
	mux.Get("/api/v1/widget/{id}", app.GetWidgetByID)

	return mux
}