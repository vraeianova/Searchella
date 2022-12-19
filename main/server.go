package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

type SearchellaServer struct {
	server *http.Server
}

func NewServer(mux *chi.Mux) *http.Server {
	s := &http.Server{
		Addr:           ":9000",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return s
}

func (s *SearchellaServer) Run() {
	s.server.ListenAndServe()
}
