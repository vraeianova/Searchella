package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes() *chi.Mux {

	mux := chi.NewMux()

	//Global Middleware
	mux.Use(
		middleware.Logger,
		middleware.Recoverer,
	)
}
