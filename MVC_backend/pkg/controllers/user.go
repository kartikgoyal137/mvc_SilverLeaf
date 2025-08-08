package controller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/kartikgoyal137/MVC/pkg/middleware/auth"
	"github.com/kartikgoyal137/MVC/pkg/types"
	"github.com/kartikgoyal137/MVC/pkg/utils"
	"strconv"
)

type UserHandler struct {
	store types.UserStore
}

func NewUserHandler(store types.UserStore) *UserHandler {
	return &UserHandler{store: store}
}

func (h *UserHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/signup", h.handleSignup).Methods("POST")
}



func (h *UserHandler) handleLogin(w http.ResponseWriter, r *http.Request) {
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
	if !auth.ComparePasswords([]byte(u.PasswordHash), []byte(user.Password)) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}

	secret := os.Getenv("TOKENKEY")
	token, err := auth.CreateJWT(secret, u.UserID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})	
}

func (h *UserHandler) handleSignup(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUser
	if err:=utils.ParseJSON(r, &payload); err!=nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	} 
	//validate here
	_, err := h.store.GetUserByEmail(payload.Email)
	if err==nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email already exists"))
		return
	}

	hashed, err := auth.HashPassword(payload.Password)
	if err!=nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.CreateNewUser(types.User{
		FirstName: payload.FirstName,
		LastName: payload.LastName,
		Contact: payload.Contact,
		Email: payload.Email,
		PasswordHash: hashed,
	})

	if err!=nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}

func (h *UserHandler) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["userID"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing user ID"))
		return
	}

	userID, err := strconv.Atoi(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid user ID"))
		return
	}

	user, err := h.store.GetUserById(userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, user)
}
