package main

import (
	"encoding/json"
	"net/http"

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

	//mux.Get("/search", nil)
	mux.Get("/search", func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("Welcome to my website!"))
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("done-by", "Cris")

		res := map[string]interface{}{"message": "Hello search"}

		_ = json.NewEncoder(w).Encode(res)
	})
	return mux
}
