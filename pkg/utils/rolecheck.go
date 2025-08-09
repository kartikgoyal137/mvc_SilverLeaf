package utils

import (
	"net/http"
	"github.com/kartikgoyal137/MVC/pkg/types"
	"fmt"
)

func RoleAuth(handlerFunc http.HandlerFunc, store types.UserStore, role string, userID int) http.HandlerFunc {
   return func(w http.ResponseWriter, r *http.Request) {

	user, err := store.GetUserById(userID)
	if err != nil {
			WriteError(w, http.StatusNotFound, fmt.Errorf("user not found"))
			return
	}

	if user.Role != role {
			WriteError(w, http.StatusForbidden, fmt.Errorf("user not authorized"))
			return
		}

    handlerFunc(w, r)
   }
}