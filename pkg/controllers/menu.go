package controller

import (
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/kartikgoyal137/MVC/pkg/types"
	"github.com/kartikgoyal137/MVC/pkg/utils"
	auth "github.com/kartikgoyal137/MVC/pkg/middleware"
)

type MenuHandler struct {
	store types.MenuStore
	userStore types.UserStore
}

func NewMenuHandler(store types.MenuStore, userStore types.UserStore) *MenuHandler {
	return &MenuHandler{store: store, userStore: userStore}
}

func (h *MenuHandler) RegisterRoutes(router *mux.Router) {
	adminHandler1 := auth.AdminAuth(h.HandleAddMenuItem, h.userStore)
	jwtAdminHandler1 := auth.JWTauth(adminHandler1, h.userStore)

	adminHandler2 := auth.AdminAuth(h.HandleRemoveMenuItem, h.userStore)
	jwtAdminHandler2 := auth.JWTauth(adminHandler2, h.userStore)
	
	router.HandleFunc("/menu/cat/all", auth.JWTauth(h.AllCategories, h.userStore)).Methods("GET")
	router.HandleFunc("/menu/cat/{id}", auth.JWTauth(h.MenuByCategory, h.userStore)).Methods("GET")
	router.HandleFunc("/menu/add", jwtAdminHandler1).Methods("PATCH")
	router.HandleFunc("/menu/remove/{product_id}", jwtAdminHandler2).Methods("DELETE")
}


func (h *MenuHandler) AllCategories(w http.ResponseWriter, r *http.Request) {
	cat, err := h.store.ListOfCategory()

	if err!=nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, cat)
}

func (h *MenuHandler) MenuByCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	productID, _ := strconv.Atoi(id)

	cat, err := h.store.GetMenuByCategoryId(productID)
	if err!=nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, cat)
}

func (h *MenuHandler) HandleAddMenuItem(w http.ResponseWriter, r *http.Request) {
    var item types.MenuItem
    if err:=utils.ParseJSON(r, &item); err != nil {
        utils.WriteError(w, http.StatusBadRequest, err)
        return
    }

    err := h.store.AddMenuItem(&item)
    if err != nil {
        utils.WriteError(w, http.StatusInternalServerError, err)
        return
    }

    utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "Menu item added successfully"})
}

func (h *MenuHandler) HandleRemoveMenuItem(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    idStr := vars["product_id"]
    productID, err := strconv.Atoi(idStr)
    if err != nil {
        utils.WriteError(w, http.StatusBadRequest, err)
        return
    }

    err = h.store.RemoveMenuItem(productID)
    if err != nil {
        utils.WriteError(w, http.StatusInternalServerError, err)
        return
    }

    utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Menu item removed successfully"})
}