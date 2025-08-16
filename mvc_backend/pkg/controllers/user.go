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
	Store types.UserStore
}

func NewUserHandler(store types.UserStore) *UserHandler {
	return &UserHandler{Store: store}
}


func (h *UserHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var user types.LoginUser
	if err := utils.Marshal(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	u, err := h.Store.GetUserByEmail(user.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
		return
	}
	if !auth.ComparePasswords(u.PasswordHash, []byte(user.Password)) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}

	token, err := auth.CreateJWT(u.UserID, u.Role)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	userID := strconv.Itoa(u.UserID)

	utils.UnMarshal(w, http.StatusOK, map[string]string{"token": token, "user_id" : userID, "name" : u.FirstName})
}

func (h *UserHandler) HandleSignup(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUser
	if err := utils.Marshal(r, &payload); err != nil {
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

	_, err := h.Store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email already exists"))
		return
	}

	hashed, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.Store.CreateNewUser(types.User{
		FirstName:    payload.FirstName,
		LastName:     payload.LastName,
		Contact:      payload.Contact,
		Email:        payload.Email,
		PasswordHash: hashed,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.UnMarshal(w, http.StatusCreated, map[string]string{"message": "Created account successfully"})
}

func (h *UserHandler) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := ctx.Value(auth.UserKey).(int)

	user, err := h.Store.GetUserById(userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.UnMarshal(w, http.StatusOK, user)
}

func (h *UserHandler) HandleGetAllUsers(w http.ResponseWriter, r *http.Request) {
	user, err := h.Store.GetAllUsers()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.UnMarshal(w, http.StatusOK, user)
}

func (h *UserHandler) ChangeUserStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	role := vars["role"]
	userid := vars["user"]
	user, err := strconv.Atoi(userid)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err2 := h.Store.ChangeUserStatus(user, role)
	if err2 != nil {
		utils.WriteError(w, http.StatusInternalServerError, err2)
		return
	}

	utils.UnMarshal(w, http.StatusOK, map[string]string{"message": "Status updated successfully"})
}
