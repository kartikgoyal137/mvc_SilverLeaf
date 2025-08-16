package api

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/kartikgoyal137/MVC/pkg/controllers"
	"github.com/kartikgoyal137/MVC/pkg/models"
	auth "github.com/kartikgoyal137/MVC/pkg/middleware"
	"log"
	"os"
	"net/http"
	"github.com/rs/cors"
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
	s.RegisterUserRoutes(subrouter, userHandler)

	orderHandler := controller.NewOrderHandler(models.NewOrderDB(s.db), userStore)
	s.RegisterOrderRoutes(subrouter, orderHandler)

	menuHandler := controller.NewMenuHandler(models.NewMenuDB(s.db), userStore)
	s.RegisterMenuRoutes(subrouter, menuHandler)

	cartHandler := controller.NewCartHandler(models.NewCartDB(s.db), userStore)
	s.RegisterCartRoutes(subrouter, cartHandler)

	paymentHandler := controller.NewPaymentHandler(models.NewPaymentDB(s.db), userStore)
	s.RegisterPaymentRoutes(subrouter, paymentHandler)

	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		staticFile := "/app/static" + r.URL.Path
		if _, err := os.Stat(staticFile); os.IsNotExist(err) {
			http.ServeFile(w, r, "/app/static/index.html")
			return
		}
		http.ServeFile(w, r, staticFile)
	})

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8080"}, 
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowedHeaders: []string{"*"},
	})

	handler := c.Handler(router)

	log.Printf("Starting server on %s\n", s.addr)
	s.Server = &http.Server{
		Addr:    s.addr,
		Handler: handler,
	}

	err := s.Server.ListenAndServe()

	return err
}

func (s *APIServer) RegisterUserRoutes(router *mux.Router, h *controller.UserHandler) {
	adminHandler1 := auth.AdminAuth(h.HandleGetAllUsers, h.Store)
	jwtAdminHandler1 := auth.JWTauth(adminHandler1, h.Store)

	adminHandler2 := auth.AdminAuth(h.ChangeUserStatus, h.Store)
	jwtAdminHandler2 := auth.JWTauth(adminHandler2, h.Store)

	router.HandleFunc("/client/admin/all", jwtAdminHandler1).Methods("GET")
	router.HandleFunc("/client/admin/status/{role}/{user}", jwtAdminHandler2).Methods("PATCH")
	router.HandleFunc("/client/login", h.HandleLogin).Methods("POST")
	router.HandleFunc("/client/signup", h.HandleSignup).Methods("POST")
	router.HandleFunc("/client/userinfo", auth.JWTauth(h.HandleGetUser, h.Store)).Methods("GET")
}

func (s *APIServer) RegisterCartRoutes(router *mux.Router, h *controller.CartHandler) {
	router.HandleFunc("/cart/add", auth.JWTauth(h.AddToCartHandler, h.UserStore)).Methods("POST")
	router.HandleFunc("/cart/edit", auth.JWTauth(h.UpdateCartHandler, h.UserStore)).Methods("PATCH")
	router.HandleFunc("/cart/remove", auth.JWTauth(h.DeleteCartItemHandler, h.UserStore)).Methods("POST")
	router.HandleFunc("/cart/get/{orderid}", auth.JWTauth(h.GetCartItemsHandler, h.UserStore)).Methods("GET")
}

func (s *APIServer) RegisterMenuRoutes(router *mux.Router, h *controller.MenuHandler) {
	adminHandler1 := auth.AdminAuth(h.HandleAddMenuItem, h.UserStore)
	jwtAdminHandler1 := auth.JWTauth(adminHandler1, h.UserStore)

	adminHandler2 := auth.AdminAuth(h.HandleRemoveMenuItem, h.UserStore)
	jwtAdminHandler2 := auth.JWTauth(adminHandler2, h.UserStore)

	router.HandleFunc("/menu/cat/all", h.AllCategories).Methods("GET")
	router.HandleFunc("/menu/cat/{id}", h.MenuByCategory).Methods("GET")
	router.HandleFunc("/menu/add", jwtAdminHandler1).Methods("PATCH")
	router.HandleFunc("/menu/remove/{product_id}", jwtAdminHandler2).Methods("DELETE")
}

func (s *APIServer) RegisterOrderRoutes(router *mux.Router, h *controller.OrderHandler) {
	chefHandler1 := auth.ChefAuth(h.HandleGetAllActiveOrders, h.UserStore)
	jwtChefHandler1 := auth.JWTauth(chefHandler1, h.UserStore)

	chefHandler2 := auth.ChefAuth(h.ChangeOrderStatus, h.UserStore)
	jwtChefHandler2 := auth.JWTauth(chefHandler2, h.UserStore)

	AdminHandler3 := auth.AdminAuth(h.HandleGetAllOrders, h.UserStore)
	jwtAdminHandler3 := auth.JWTauth(AdminHandler3, h.UserStore)

	router.HandleFunc("/orders/place", auth.JWTauth(h.PlaceOrder, h.UserStore)).Methods("POST")
	router.HandleFunc("/orders/user", auth.JWTauth(h.HandleMyOrders, h.UserStore)).Methods("GET")
	router.HandleFunc("/orders/start", auth.JWTauth(h.CreateOrderHandler, h.UserStore)).Methods("POST")
	router.HandleFunc("/orders/chef/active", jwtChefHandler1).Methods("GET")
	router.HandleFunc("/orders/chef/status", jwtChefHandler2).Methods("POST")
	router.HandleFunc("/orders/admin/all", jwtAdminHandler3).Methods("GET")
}

func (s *APIServer) RegisterPaymentRoutes(router *mux.Router, h *controller.PaymentHandler) {
	adminHandler1 := auth.AdminAuth(h.HandleGetAllPayments, h.UserStore)
	jwtAdminHandler1 := auth.JWTauth(adminHandler1, h.UserStore)

	adminHandler2 := auth.AdminAuth(h.ChangePaymentStatus, h.UserStore)
	jwtAdminHandler2 := auth.JWTauth(adminHandler2, h.UserStore)

	router.HandleFunc("/payments/admin/all", jwtAdminHandler1).Methods("GET")
	router.HandleFunc("/payments/admin/status", jwtAdminHandler2).Methods("PATCH")
	router.HandleFunc("/payments/user", auth.JWTauth(h.HandleGetPayByUser, h.UserStore)).Methods("GET")
	router.HandleFunc("/payments/total/{order_id}", auth.JWTauth(h.HandleCalculateTotal, h.UserStore)).Methods("GET")
	router.HandleFunc("/payments/new", auth.JWTauth(h.HandleNewPayment, h.UserStore)).Methods("POST")

}
