package controller

import (
	"net/http"
	auth "github.com/kartikgoyal137/MVC/pkg/middleware"
	"github.com/kartikgoyal137/MVC/pkg/types"
	"github.com/kartikgoyal137/MVC/pkg/utils"
)

type OrderHandler struct {
	store     types.OrderStore
	UserStore types.UserStore
}

func NewOrderHandler(store types.OrderStore, userStore types.UserStore) *OrderHandler {
	return &OrderHandler{store: store, UserStore: userStore}
}


func (h *OrderHandler) PlaceOrder(w http.ResponseWriter, r *http.Request) {

	var order types.CreateOrder
	if err := utils.Marshal(r, &order); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err := h.store.CreateOrder(order)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.UnMarshal(w, http.StatusCreated, map[string]string{"message": "Order created successfully"})

}

func (h *OrderHandler) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := ctx.Value(auth.UserKey).(int)
	orderID, err := h.store.CreateEmptyOrder(userID)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.UnMarshal(w, http.StatusCreated, map[string]int{"order_id": orderID})
}

func (h *OrderHandler) HandleGetAllOrders(w http.ResponseWriter, r *http.Request) {

	orders, err := h.store.GetAllOrders()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.UnMarshal(w, http.StatusOK, orders)
}

func (h *OrderHandler) HandleGetAllActiveOrders(w http.ResponseWriter, r *http.Request) {

	orders, err := h.store.GetAllActiveOrders()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.UnMarshal(w, http.StatusOK, orders)
}

func (h *OrderHandler) ChangeOrderStatus(w http.ResponseWriter, r *http.Request) {

	var payload types.ChangeOrderStatusPayload
	if err := utils.Marshal(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err := h.store.ChangeStatus(payload.OrderID, payload.Status)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.UnMarshal(w, http.StatusOK, map[string]string{"message": "Status updated successfully"})
}

func (h *OrderHandler) HandleMyOrders(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := ctx.Value(auth.UserKey).(int)

	item, err := h.store.OrdersByUserId(userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.UnMarshal(w, http.StatusOK, item)
}
