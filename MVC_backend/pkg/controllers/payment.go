package controller

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/kartikgoyal137/MVC/pkg/middleware"
	"github.com/kartikgoyal137/MVC/pkg/types"
	"github.com/kartikgoyal137/MVC/pkg/utils"
)

type PayHandler struct {
	store types.PaymentStore
	userStore types.UserStore
}

func NewPayHandler(store types.PaymentStore, userstore types.UserStore) *PayHandler {
	return &PayHandler{store: store, userStore: userstore}
}

func (h *PayHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/admin/allpayments", auth.JWTauth(auth.AdminAuth(h.HandleGetAllPayments,h.userStore), h.userStore)).Methods("GET")
	router.HandleFunc("/mypayments", auth.JWTauth(h.HandleGetPayByUser , h.userStore)).Methods("GET")
	router.HandleFunc("/payment", auth.JWTauth(h.HandleNewPayment , h.userStore)).Methods("GET")
	router.HandleFunc("/admin/paystatus", auth.JWTauth(auth.AdminAuth(h.ChangePaymentStatus, h.userStore), h.userStore)).Methods("POST")
}

func (h *PayHandler) HandleGetAllPayments(w http.ResponseWriter, r *http.Request) {
	
	payments, err := h.store.GetAllPayments()
	if err!=nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, payments)
}

func (h *PayHandler) HandleGetPayByUser(w http.ResponseWriter, r *http.Request) {
	
	ctx := r.Context()
	userID := ctx.Value(auth.UserKey).(int)
	payments, err := h.store.PaymentsByUserId(userID)
	if err!=nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, payments)
}

func (h *PayHandler) HandleNewPayment(w http.ResponseWriter, r *http.Request) {
	
	var payload types.MakePayment
	if err:=utils.ParseJSON(r, &payload); err!=nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	} 

	err := h.store.CreateNewPayment(&payload)

	utils.WriteJSON(w, http.StatusOK, err)
}

func (h *PayHandler) ChangePaymentStatus(w http.ResponseWriter, r *http.Request) {

	var payload types.ChangePaymentStatusPayload
	if err:=utils.ParseJSON(r, &payload); err!=nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err := h.store.ChangePayStatus(payload.OrderId, payload.Status)
	if err!=nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, err)
}
