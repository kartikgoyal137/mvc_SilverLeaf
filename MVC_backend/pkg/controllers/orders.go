package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kartikgoyal137/MVC/pkg/types"
	"github.com/kartikgoyal137/MVC/pkg/utils"
	auth "github.com/kartikgoyal137/MVC/pkg/middleware"
)

type OrderHandler struct {
	store types.OrderStore
	userStore types.UserStore
}

func NewOrderHandler(store types.OrderStore, userStore types.UserStore) *OrderHandler {
	return &OrderHandler{store: store, userStore: userStore}
}

func (h *OrderHandler) RegisterRoutes(router *mux.Router) {
	ChefHandler1 := auth.ChefAuth(h.HandleGetAllActiveOrders, h.userStore)
	jwtChefHandler1 := auth.JWTauth(ChefHandler1, h.userStore)

	ChefHandler2 := auth.ChefAuth(h.ChangeOrderStatus, h.userStore)
	jwtChefHandler2 := auth.JWTauth(ChefHandler2, h.userStore)

	AdminHandler3 := auth.AdminAuth(h.HandleGetAllOrders, h.userStore)
	jwtAdminHandler3 := auth.JWTauth(AdminHandler3, h.userStore)


	router.HandleFunc("/placeorder", auth.JWTauth(h.PlaceOrder, h.userStore)).Methods("POST")
	router.HandleFunc("/orders/user", auth.JWTauth(h.HandleMyOrders, h.userStore)).Methods("GET")
	router.HandleFunc("/orders/start", auth.JWTauth(h.CreateOrderHandler, h.userStore)).Methods("POST")
	router.HandleFunc("/orders/chef/active", jwtChefHandler1).Methods("GET")
	router.HandleFunc("/orders/chef/status", jwtChefHandler2).Methods("POST")
	router.HandleFunc("/orders/chef/all", jwtAdminHandler3).Methods("GET")
}



func (h *OrderHandler) PlaceOrder(w http.ResponseWriter, r *http.Request) {

	var order types.CreateOrder
	if err:=utils.ParseJSON(r, &order); err!=nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err := h.store.CreateOrder(order)
	if err!=nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

}

func (h *OrderHandler) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := ctx.Value(auth.UserKey).(int)
	orderID, err := h.store.CreateEmptyOrder(userID)

	if err!=nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, orderID)
}

func (h *OrderHandler) HandleGetAllOrders(w http.ResponseWriter, r *http.Request) {

	orders, err := h.store.GetAllOrders()
	if err!=nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, orders)
}

func (h *OrderHandler) HandleGetAllActiveOrders(w http.ResponseWriter, r *http.Request) {

	orders, err := h.store.GetAllActiveOrders()
	if err!=nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, orders)
}

func (h *OrderHandler) ChangeOrderStatus(w http.ResponseWriter, r *http.Request) {

	var payload types.ChangeOrderStatusPayload
	if err:=utils.ParseJSON(r, &payload); err!=nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err := h.store.ChangeStatus(payload.OrderID, payload.Status)
	if err!=nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Status updated successfully"})
}

func (h *OrderHandler) HandleMyOrders(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := ctx.Value(auth.UserKey).(int)

	item, err := h.store.OrdersByUserId(userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusAccepted, item)
}
