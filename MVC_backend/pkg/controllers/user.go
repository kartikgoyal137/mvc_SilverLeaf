package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	auth "github.com/kartikgoyal137/MVC/pkg/middleware"
	"github.com/kartikgoyal137/MVC/pkg/types"
	"github.com/kartikgoyal137/MVC/pkg/utils"
)

type UserHandler struct {
	store types.UserStore
}

func NewUserHandler(store types.UserStore) *UserHandler {
	return &UserHandler{store: store}
}

func (h *UserHandler) RegisterRoutes(router *mux.Router) {
	adminHandler1 := auth.AdminAuth(h.HandleGetAllUsers, h.store)
	jwtAdminHandler1 := auth.JWTauth(adminHandler1, h.store)

	adminHandler2 := auth.AdminAuth(h.ChangeUserStatus, h.store)
	jwtAdminHandler2 := auth.JWTauth(adminHandler2, h.store)

	router.HandleFunc("/admin/allusers", jwtAdminHandler1).Methods("GET")
	router.HandleFunc("/admin/userstatus/{role}/{user}", jwtAdminHandler2).Methods("POST")
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/signup", h.handleSignup).Methods("POST")
	router.HandleFunc("/userinfo", auth.JWTauth(h.HandleGetUser , h.store)).Methods("GET")
}



func (h *UserHandler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var user types.LoginUser
	if err:=utils.ParseJSON(r, &user); err!=nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	u, err := h.store.GetUserByEmail(user.Email)
	if err!=nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
		return
	}
	if !auth.ComparePasswords(u.PasswordHash, []byte(user.Password)) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}

	token, err := auth.CreateJWT(u.UserID)
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
	ctx := r.Context()
	userID := ctx.Value(auth.UserKey).(int)

	user, err := h.store.GetUserById(userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, user)
}

func (h *UserHandler) HandleGetAllUsers(w http.ResponseWriter, r *http.Request) {
	user, err := h.store.GetAllUsers()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, user)
}

func (h *UserHandler) ChangeUserStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	role := vars["role"]
	userid := vars["user"]
	user, err := strconv.Atoi(userid)
	if err!=nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err2 := h.store.ChangeUserStatus(user, role)
	if err2!=nil {
		utils.WriteError(w, http.StatusBadRequest, err2)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Status updated successfully"})
}


