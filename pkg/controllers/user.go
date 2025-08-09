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

	router.HandleFunc("/client/admin/all", jwtAdminHandler1).Methods("GET")
	router.HandleFunc("/client/admin/status/{role}/{user}", jwtAdminHandler2).Methods("PATCH")
	router.HandleFunc("/client/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/client/signup", h.handleSignup).Methods("POST")
	router.HandleFunc("/client/userinfo", auth.JWTauth(h.HandleGetUser , h.store)).Methods("GET")
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

	if payload.FirstName == "" || payload.Email == "" || payload.Password == "" || payload.LastName == "" {
	utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("all fields are required"))
	return
    }

	if len(payload.Password) < 8 {
        utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("password must be at least 8 characters long"))
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

	utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "Created account successfully"})
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
		utils.WriteError(w, http.StatusInternalServerError, err2)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Status updated successfully"})
}


