package app

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"fmt"
)

const port = 8080

type Server struct {
	Router *mux.Router
}

func New() *Server {
	s := &Server{
		Router: mux.NewRouter(),
	}
	return s
}

func (s *Server) Start() error {
	srv := http.Server{
		Handler:s.Router,
		Addr:              fmt.Sprintf(":%d", port),
		WriteTimeout:      5 * time.Second,
		ReadTimeout:       5 * time.Second,
		IdleTimeout:       5 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
	}
	return srv.ListenAndServe()
}
