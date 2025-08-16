package auth

import (

	"github.com/kartikgoyal137/MVC/pkg/types"

	"net/http"
)

func ChefAuth(handlerFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		RoleAuth(handlerFunc, store, "chef")(w, r)
	}
}
