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