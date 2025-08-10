package controller

import (
	"fmt"
	"net/http"
	"strconv"
	sqldriver "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/kartikgoyal137/MVC/pkg/types"
	"github.com/kartikgoyal137/MVC/pkg/utils"
)

type CartHandler struct {
	store     types.CartStore
	UserStore types.UserStore
}

func NewCartHandler(store types.CartStore, userStore types.UserStore) *CartHandler {
	return &CartHandler{store: store, UserStore: userStore}
}


func (h *CartHandler) AddToCartHandler(w http.ResponseWriter, r *http.Request) {

	var item types.CartItem
	if err := utils.Marshal(r, &item); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if item.Quantity <= 0 || item.Quantity > 100 {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("quantity must be between 1 and 100"))
		return
	}

	err := h.store.AddToCart(item)
	if err != nil {
		if mysqlErr, ok := err.(*sqldriver.MySQLError); ok && mysqlErr.Number == 1062 {
			err2 := h.store.UpdateCartItemQuantity(item)
			utils.UnMarshal(w, http.StatusOK, map[string]string{"message": "Item updated successfully"})
			if err2 != nil {
				utils.WriteError(w, http.StatusBadRequest, err2)
				return
			}
			return
		} else {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}
	}

	utils.UnMarshal(w, http.StatusCreated, map[string]string{"message": "Item added successfully"})
}

func (h *CartHandler) DeleteCartItemHandler(w http.ResponseWriter, r *http.Request) {

	var item types.CartItem
	if err := utils.Marshal(r, &item); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	err := h.store.DeleteCartItem(item)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.UnMarshal(w, http.StatusOK, map[string]string{"message": "Item deleted successfully"})
}

func (h *CartHandler) UpdateCartHandler(w http.ResponseWriter, r *http.Request) {

	var item types.CartItem
	if err := utils.Marshal(r, &item); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if item.Quantity <= 0 || item.Quantity > 100 {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("quantity must be between 1 and 100"))
		return
	}

	err := h.store.UpdateCartItemQuantity(item)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.UnMarshal(w, http.StatusOK, map[string]string{"message": "updated quantity successfully"})
}

func (h *CartHandler) GetCartItemsHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	orderID, err := strconv.Atoi(vars["orderid"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	items, err := h.store.GetCartItems(int(orderID))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.UnMarshal(w, http.StatusAccepted, items)
}
