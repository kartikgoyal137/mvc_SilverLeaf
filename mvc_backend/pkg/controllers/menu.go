package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/kartikgoyal137/MVC/pkg/models"
	"github.com/gorilla/mux"
	"github.com/kartikgoyal137/MVC/pkg/types"
	"github.com/kartikgoyal137/MVC/pkg/utils"
)

type MenuHandler struct {
	store     types.MenuStore
	UserStore types.UserStore
}

func NewMenuHandler(store types.MenuStore, userStore types.UserStore) *MenuHandler {
	return &MenuHandler{store: store, UserStore: userStore}
}

func (h *MenuHandler) AllCategories(w http.ResponseWriter, r *http.Request) {

	models.CacheMutex.RLock()
	cachedData := models.CategoryCacheString
	models.CacheMutex.RUnlock()

	if cachedData != "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(models.CategoryCacheString))
		return
	}

	cat, err := h.store.ListOfCategory()

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.UnMarshal(w, http.StatusOK, cat)
}

func (h *MenuHandler) MenuByCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	productID, err := strconv.Atoi(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to convert productID into integer"))
		return
	}

	models.CacheMutex.RLock()
	menuJSON, found := models.MenuCache[productID]
	models.CacheMutex.RUnlock()

	if found {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(menuJSON))
		return
	}

	cat, err := h.store.GetMenuByCategoryId(productID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.UnMarshal(w, http.StatusOK, cat)
}

func (h *MenuHandler) HandleAddMenuItem(w http.ResponseWriter, r *http.Request) {
	var item types.MenuItem
	if err := utils.Marshal(r, &item); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err := h.store.AddMenuItem(&item)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.UnMarshal(w, http.StatusCreated, map[string]string{"message": "Menu item added successfully"})
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

	utils.UnMarshal(w, http.StatusOK, map[string]string{"message": "Menu item removed successfully"})
}
