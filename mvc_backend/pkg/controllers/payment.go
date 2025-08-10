package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	auth "github.com/kartikgoyal137/MVC/pkg/middleware"
	"github.com/kartikgoyal137/MVC/pkg/types"
	"github.com/kartikgoyal137/MVC/pkg/utils"
)

type PaymentHandler struct {
	store     types.PaymentStore
	UserStore types.UserStore
}

func NewPaymentHandler(store types.PaymentStore, userstore types.UserStore) *PaymentHandler {
	return &PaymentHandler{store: store, UserStore: userstore}
}

func (h *PaymentHandler) HandleGetAllPayments(w http.ResponseWriter, r *http.Request) {

	payments, err := h.store.GetAllPayments()
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.UnMarshal(w, http.StatusOK, payments)
}

func (h *PaymentHandler) HandleGetPayByUser(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	userID := ctx.Value(auth.UserKey).(int)
	payments, err := h.store.PaymentsByUserId(userID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.UnMarshal(w, http.StatusOK, payments)
}

func (h *PaymentHandler) HandleNewPayment(w http.ResponseWriter, r *http.Request) {

	var payload types.MakePayment
	if err := utils.Marshal(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err := h.store.CreateNewPayment(&payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.UnMarshal(w, http.StatusCreated, map[string]string{"message": "Payment created successfully"})
}

func (h *PaymentHandler) ChangePaymentStatus(w http.ResponseWriter, r *http.Request) {

	var payload types.ChangePaymentStatusPayload
	if err := utils.Marshal(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err := h.store.ChangePayStatus(payload.OrderId, payload.Status)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.UnMarshal(w, http.StatusOK, map[string]string{"message": "Status updated successfully"})
}

func (h *PaymentHandler) HandleCalculateTotal(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, ok := vars["order_id"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("failed to fetch"))
		return
	}
	orderID, err := strconv.Atoi(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	total, err := h.store.CalculateTotal(orderID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.UnMarshal(w, http.StatusOK, total)
}
