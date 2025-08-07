package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kartikgoyal137/MVC/pkg/types"
)

type OrderHandler struct {
	store types.OrderStore
}

func NewOrderHandler(store types.OrderStore) *OrderHandler {
	return &OrderHandler{store: store}
}

func (h *OrderHandler) RegisterRoutes(router *mux.Router) {
	
}



func (h *OrderHandler) handleA(w http.ResponseWriter, r *http.Request) {
	
}

func (h *OrderHandler) handleB(w http.ResponseWriter, r *http.Request) {
 
}
