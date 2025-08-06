package api

import (
	"log"
	"database/sql"
	"net/http"
	"github.com/kartikgoyal137/MVC/pkg/controllers"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db *sql.DB
} 

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr : addr,
		db : db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)
	log.Printf("Starting server on %s\n", s.addr)

	return http.ListenAndServe(s.addr, router)
}