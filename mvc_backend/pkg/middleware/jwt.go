package auth

import (
	"context"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"

	"github.com/kartikgoyal137/MVC/pkg/types"
	"github.com/kartikgoyal137/MVC/pkg/utils"
	"time"
)

type contextKey string

const UserKey contextKey = "userID"
const RoleKey contextKey = "role"

func JWTauth(handlerFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := utils.GetTokenFromRequest(r)
		token, err := VerifyJWT(tokenString)
		if err != nil {
			utils.WriteError(w, http.StatusUnauthorized, err)
			return
		}

		if !token.Valid {
			utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("unauthorised access"))
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userIDfloat := claims["userID"].(float64)
		userID := int(userIDfloat)

		u, err := store.GetUserById(userID)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, u.UserID)
		ctx = context.WithValue(ctx, RoleKey, u.Role)
		r = r.WithContext(ctx)

		handlerFunc(w, r)
	}
}

func CreateJWT(userID int, role string) (string, error) {
	var secretKey = os.Getenv("TOKENKEY")
	if secretKey == "" {
		return "", fmt.Errorf("TOKENKEY environment variable not set")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userID": userID,
			"role": role,
			"expiresAt": time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, err
}

func VerifyJWT(tokenString string) (*jwt.Token, error) {
	var secretKey = os.Getenv("TOKENKEY")
	if secretKey == "" {
		return nil, fmt.Errorf("TOKENKEY environment variable not set")
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}
