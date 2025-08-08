package auth

import (
	"context"
	"fmt"
	"log"
   "github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strconv"
	"time"
   "github.com/kartikgoyal137/MVC/pkg/types"
	"github.com/kartikgoyal137/MVC/pkg/utils"
)

type contextKey string
const UserKey contextKey = "userID"

func JWTauth(handlerFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
   return func(w http.ResponseWriter, r *http.Request) {
      tokenString := utils.GetTokenFromRequest(r)
      token, err := VerifyJWT(tokenString)
      if err!=nil {
         utils.WriteError(w, http.StatusUnauthorized, err)
         return
      }

      if !token.Valid  {
         utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("unauthorised access"))
         return
      }

      claims := token.Claims.(jwt.MapClaims)
      str := claims["userID"].(string)

      userID, err := strconv.Atoi(str)
      if err!=nil {
         log.Printf("failed to convert")
         return
      }

      u, err := store.GetUserById(userID)
      if err!=nil {
         utils.WriteError(w, http.StatusBadRequest, err)
         return
      }

      ctx := r.Context()
      ctx = context.WithValue(ctx, UserKey, u.UserID)
      r = r.WithContext(ctx)

      handlerFunc(w, r)
   }
}

func CreateJWT(secret string ,userID int) (string, error) {
	var secretKey = os.Getenv("TOKENKEY")
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
        jwt.MapClaims{ 
        "user_id": userID, 
        "expiresAt": time.Now().Add(time.Hour * 24).Unix(), 
        })

    tokenString, err := token.SignedString(secretKey)
    if err != nil {
    return "", err
    }

 return tokenString, err
}

func VerifyJWT(tokenString string) (*jwt.Token,error) {
   var secretKey = os.Getenv("TOKENKEY")
   token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
      return secretKey, nil
   })
  
   if err != nil {
      return nil, err
   }
  
   if !token.Valid {
      return nil, fmt.Errorf("invalid token")
   }
  
   return token ,nil
}