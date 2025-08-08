package controller

import (
	"net/http"
    "strconv"
	"github.com/gorilla/mux"
	"github.com/kartikgoyal137/MVC/pkg/types"
	"github.com/kartikgoyal137/MVC/pkg/utils"
	"github.com/kartikgoyal137/MVC/pkg/middleware/auth"
)

type CartHandler struct {
	store types.CartStore
	userStore types.UserStore
}

func NewCartHandler(store types.CartStore, userStore types.UserStore) *CartHandler {
	return &CartHandler{store: store, userStore: userStore}
}

func (h *CartHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/addtocart", auth.JWTauth(h.AddToCartHandler, h.userStore)).Methods("POST")
	router.HandleFunc("/getcart/{orderid}", auth.JWTauth(h.GetCartItemsHandler, h.userStore)).Methods("GET")
}

func (h *CartHandler) AddToCartHandler(w http.ResponseWriter, r *http.Request) {

    var item types.CartItem
    if err := utils.ParseJSON(r, &item); err != nil {
        utils.WriteError(w, http.StatusBadRequest, err)
        return
    }

    err := h.store.AddToCart(item)
    if err != nil {
		err2 := h.store.UpdateCartItemQuantity(item)
        if err2!=nil {
			utils.WriteError(w, http.StatusBadRequest, err2)
			return
		}
    }

    utils.WriteJSON(w, http.StatusCreated, nil)
}

func (h *CartHandler) GetCartItemsHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
    orderID, _ := strconv.Atoi(vars["orderid"])
   
    items, err := h.store.GetCartItems(int(orderID))
    if err != nil {
        utils.WriteError(w, http.StatusInternalServerError, err)
        return
    }

	utils.WriteJSON(w, http.StatusAccepted, items)
}

