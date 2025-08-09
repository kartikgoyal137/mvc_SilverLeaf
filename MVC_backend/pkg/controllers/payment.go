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

type PayHandler struct {
	store types.PaymentStore
	userStore types.UserStore
}

func NewPayHandler(store types.PaymentStore, userstore types.UserStore) *PayHandler {
	return &PayHandler{store: store, userStore: userstore}
}


func (h *PayHandler) RegisterRoutes(router *mux.Router) {
	adminHandler1 := auth.AdminAuth(h.HandleGetAllPayments, h.userStore)
	jwtAdminHandler1 := auth.JWTauth(adminHandler1, h.userStore)

	adminHandler2 := auth.AdminAuth(h.ChangePaymentStatus, h.userStore)
	jwtAdminHandler2 := auth.JWTauth(adminHandler2, h.userStore)


	router.HandleFunc("/payments/admin/all", jwtAdminHandler1).Methods("GET")
	router.HandleFunc("/payments/admin/status", jwtAdminHandler2).Methods("POST")
	router.HandleFunc("/payments/user", auth.JWTauth(h.HandleGetPayByUser , h.userStore)).Methods("GET")
	router.HandleFunc("/payments/total/{order_id}", auth.JWTauth(h.HandleCalculateTotal , h.userStore)).Methods("GET")
	router.HandleFunc("/payments/new", auth.JWTauth(h.HandleNewPayment , h.userStore)).Methods("POST")
	
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
	if err!=nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Payment created successfully"})
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

	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Status updated successfully"})
}

func (h *PayHandler) HandleCalculateTotal(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, ok := vars["order_id"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("failed to fetch"))
		return
	}
	orderID, err := strconv.Atoi(id)
	if err!=nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	total, err := h.store.CalculateTotal(orderID)
	if err!=nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	

	utils.WriteJSON(w, http.StatusOK, total)
}
