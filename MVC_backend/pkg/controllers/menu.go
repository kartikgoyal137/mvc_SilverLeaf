package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kartikgoyal137/MVC/pkg/types"
	"github.com/kartikgoyal137/MVC/pkg/utils"
	"github.com/kartikgoyal137/MVC/pkg/middleware/auth"
)

type MenuHandler struct {
	store types.MenuStore
	userStore types.UserStore
}

func NewMenuHandler(store types.MenuStore, userStore types.UserStore) *MenuHandler {
	return &MenuHandler{store: store, userStore: userStore}
}

func (h *MenuHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/api/category", auth.JWTauth(h.AllCategories, h.userStore)).Methods("GET")
	router.HandleFunc("/api/menu", auth.JWTauth(h.MenuByCategory, h.userStore)).Methods("GET")
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
	var id int
	if err:=utils.ParseJSON(r, &id); err!=nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	cat, err := h.store.GetMenuByCategoryId(id)
	if err!=nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, cat)
}
