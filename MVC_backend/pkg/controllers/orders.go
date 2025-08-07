package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kartikgoyal137/MVC/pkg/types"
	"github.com/kartikgoyal137/MVC/pkg/utils"
	"github.com/kartikgoyal137/MVC/pkg/middleware/auth"
)

type OrderHandler struct {
	store types.OrderStore
	userStore types.UserStore
}

func NewOrderHandler(store types.OrderStore, userStore types.UserStore) *OrderHandler {
	return &OrderHandler{store: store, userStore: userStore}
}

func (h *OrderHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/api/order", auth.JWTauth(h.PlaceOrder, h.userStore)).Methods("POST")
}



func (h *OrderHandler) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	var order types.CreateOrder
	if err:=utils.ParseJSON(r, &order); err!=nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err := h.store.UpdateOrder(order)
	if err!=nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
}

func (h *OrderHandler) EmptyOrder(w http.ResponseWriter, r *http.Request) {
	var order types.User
	if err := h.store.CreateEmptyOrder(order); err!=nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	
	utils.WriteJSON(w, http.StatusAccepted, "ok")
}

func (h *OrderHandler) EditOrder(w http.ResponseWriter, r *http.Request) {

}
