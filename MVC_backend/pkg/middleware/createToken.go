package middleware

import (
 "github.com/golang-jwt/jwt/v5"
 "time"
  "os"
  "github.com/joho/godotenv"
  "fmt"
)

func init() {
	godotenv.Load()
}

func CreateToken(username string) (string, error) {
	var secretKey = os.Getenv("TOKENKEY")
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
        jwt.MapClaims{ 
        "username": username, 
        "exp": time.Now().Add(time.Hour * 24).Unix(), 
        })

    tokenString, err := token.SignedString(secretKey)
    if err != nil {
    return "", err
    }

 return tokenString, nil
}

func VerifyToken(tokenString string) error {
   var secretKey = os.Getenv("TOKENKEY")
   token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
      return secretKey, nil
   })
  
   if err != nil {
      return err
   }
  
   if !token.Valid {
      return fmt.Errorf("invalid token")
   }
  
   return nil
}