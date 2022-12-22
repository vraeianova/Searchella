package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

type SearchellaServer struct {
	server *http.Server
}

func NewServer(mux *chi.Mux) *SearchellaServer {
	s := &http.Server{
		Addr:           ":9000",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return &SearchellaServer{s}
}

func (s *SearchellaServer) Run() {
	s.server.ListenAndServe()
}
