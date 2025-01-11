package main

import (
	"net/http"

	v1 "github.com/nyybl/scrapynato/routes/api/v1"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	router *chi.Mux
}

func NewServer() *Server {
	s := Server{
		router: chi.NewRouter(),
	}
	s.router.Use(middleware.Logger)

	v1.RegisterV1(s.router)
	return &s
}

func (s *Server) Listen(listenAddr string) {
	http.ListenAndServe(listenAddr, s.router)
}