package auth

import (
	"fmt"
	"net/http"
	"github.com/kartikgoyal137/MVC/pkg/types"
	"github.com/kartikgoyal137/MVC/pkg/utils"
)

func ChefAuth(handlerFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
   return func(w http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value(UserKey).(int)
	if !ok {
			utils.WriteError(w, http.StatusNotFound, fmt.Errorf("user not found"))
			return
	}

	user, err := store.GetUserById(userID)
	if err != nil {
			utils.WriteError(w, http.StatusNotFound, fmt.Errorf("user not found"))
			return
	}

	if user.Role != "chef" {
			utils.WriteError(w, http.StatusForbidden, fmt.Errorf("user not authorized"))
			return
		}

    handlerFunc(w, r)
   }
}