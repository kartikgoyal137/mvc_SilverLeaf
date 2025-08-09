package api

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/kartikgoyal137/MVC/pkg/controllers"
	"github.com/kartikgoyal137/MVC/pkg/models"
	"log"
	"net/http"
)

type APIServer struct {
	addr   string
	db     *sql.DB
	Server *http.Server
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := models.NewUserDB(s.db)
	userHandler := controller.NewUserHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	orderHandler := controller.NewOrderHandler(models.NewOrderDB(s.db), userStore)
	orderHandler.RegisterRoutes(subrouter)

	menuHandler := controller.NewMenuHandler(models.NewMenuDB(s.db), userStore)
	menuHandler.RegisterRoutes(subrouter)

	cartHandler := controller.NewCartHandler(models.NewCartDB(s.db), userStore)
	cartHandler.RegisterRoutes(subrouter)

	log.Printf("Starting server on %s\n", s.addr)
	s.Server = &http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	err := s.Server.ListenAndServe()

	return err
}
