package api

import (
	"database/sql"
	"log"
	"net/http"
)

type APIServer struct {
	addr   string
	db     *sql.DB
	router *http.ServeMux
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr:   addr,
		db:     db,
		router: http.NewServeMux(),
	}
}

func (s *APIServer) Run() error {
	s.registerServiceRoutes()

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, s.router)
}

func (s *APIServer) registerServiceRoutes() {
	
}
