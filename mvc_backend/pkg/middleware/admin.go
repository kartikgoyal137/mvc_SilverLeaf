package auth

import (
	"fmt"
	"github.com/kartikgoyal137/MVC/pkg/types"
	"github.com/kartikgoyal137/MVC/pkg/utils"
	"net/http"
)

func AdminAuth(handlerFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userID, ok := r.Context().Value(UserKey).(int)
		if !ok {
			utils.WriteError(w, http.StatusNotFound, fmt.Errorf("user not found"))
			return
		}

		utils.RoleAuth(handlerFunc, store, "administrator", userID)(w, r)

	}
}
