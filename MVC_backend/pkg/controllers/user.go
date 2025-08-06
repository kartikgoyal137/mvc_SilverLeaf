package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kartikgoyal137/MVC/pkg/middleware"
	"github.com/kartikgoyal137/MVC/pkg/utils"
	"github.com/kartikgoyal137/MVC/pkg/types"
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
	middleware.CreateToken("kartik")
}

func (h *Handler) handleSignup(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUser
	if err:=utils.ParseJSON(r, &payload); err!=nil {
	utils.WriteError(w, http.StatusBadRequest, err)
	} 
}
