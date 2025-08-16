package auth

import (
	
	"github.com/kartikgoyal137/MVC/pkg/types"
	
	"net/http"
)

func AdminAuth(handlerFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		RoleAuth(handlerFunc, store, "administrator")(w, r)

	}
}
