package auth

import (
	"fmt"
	"net/http"
	"github.com/kartikgoyal137/MVC/pkg/types"
	"github.com/kartikgoyal137/MVC/pkg/utils"
)

func RoleAuth(handlerFunc http.HandlerFunc, store types.UserStore, role string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		userRole, ok := ctx.Value(RoleKey).(string)
		if !ok {
			utils.WriteError(w, http.StatusNotFound, fmt.Errorf("user not found"))
			return
		}

		if userRole != role {
			utils.WriteError(w, http.StatusForbidden, fmt.Errorf("user not authorized"))
			return
		}

		handlerFunc(w, r)
	}
}
