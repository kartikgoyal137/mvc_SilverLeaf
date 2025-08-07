package api

import (
	"log"
	"database/sql"
	"net/http"
	"github.com/kartikgoyal137/MVC/pkg/controllers"
	"github.com/kartikgoyal137/MVC/pkg/models"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db *sql.DB
	Server *http.Server
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

	userStore := models.NewUserDB(models.DB)
	userHandler := controller.NewUserHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	orderStore := models.NewOrderDB(models.DB)
	orderHandler := controller.NewOrderHandler(orderStore, userStore)
	orderHandler.RegisterRoutes(subrouter)

	menuStore := models.NewMenuDB(models.DB)
	menuHandler := controller.NewMenuHandler(menuStore, userStore)
	menuHandler.RegisterRoutes(subrouter)

	
	log.Printf("Starting server on %s\n", s.addr)
	s.Server = &http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	return s.Server.ListenAndServe()
}

