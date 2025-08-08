package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kartikgoyal137/MVC/pkg/types"
	"github.com/kartikgoyal137/MVC/pkg/utils"
	"github.com/kartikgoyal137/MVC/pkg/middleware"
)

type OrderHandler struct {
	store types.OrderStore
	userStore types.UserStore
}

func NewOrderHandler(store types.OrderStore, userStore types.UserStore) *OrderHandler {
	return &OrderHandler{store: store, userStore: userStore}
}

func (h *OrderHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/placeorder", auth.JWTauth(h.PlaceOrder, h.userStore)).Methods("POST")
	router.HandleFunc("/startorder", auth.JWTauth(h.CreateOrderHandler, h.userStore)).Methods("POST")
	router.HandleFunc("/allorders", auth.JWTauth(auth.ChefAuth(h.HandleGetAllOrders,h.userStore), h.userStore)).Methods("GET")
	router.HandleFunc("/allorders", auth.JWTauth(auth.ChefAuth(h.ChangeOrderStatus,h.userStore), h.userStore)).Methods("GET")
	router.HandleFunc("/admin/paystatus", auth.JWTauth(auth.AdminAuth(h.ChangeOrderStatus, h.userStore), h.userStore)).Methods("POST")

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

func (h *OrderHandler) ChangeOrderStatus(w http.ResponseWriter, r *http.Request) {

	var payload types.ChangeOrderStatusPayload
	if err:=utils.ParseJSON(r, &payload); err!=nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err := h.store.ChangeStatus(payload.OrderID, payload.Status)
	if err!=nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, err)
}
