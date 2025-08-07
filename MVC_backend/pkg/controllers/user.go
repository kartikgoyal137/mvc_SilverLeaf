package user

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kartikgoyal137/MVC/pkg/middleware"
	"github.com/kartikgoyal137/MVC/pkg/types"
	"github.com/kartikgoyal137/MVC/pkg/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/signup", h.handleSignup).Methods("POST")
}



func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var user types.LoginUser
	if err:=utils.ParseJSON(r, &user); err!=nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	//validate here

	u, err := h.store.GetUserByEmail(user.Email)
	if err!=nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
		return
	}

	
}

func (h *Handler) handleSignup(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUser
	if err:=utils.ParseJSON(r, &payload); err!=nil {
	utils.WriteError(w, http.StatusBadRequest, err)
	} 
	
	h.store.CreateNewUser(payload) 
}
