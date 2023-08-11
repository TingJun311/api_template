package ojswsj 

import (
	"fmt"
	"time"
	"github.com/dgrijalva/jwt-go"
)


var jwtSecret = []byte("your-secret-key")

func generateToken(userID string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Second * 10).Unix(), // Token expiration time (1 day)
    })
    return token.SignedString(jwtSecret)
}

func validateToken(tokenString string) (jwt.MapClaims, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })

    if err != nil || !token.Valid {
        return nil, fmt.Errorf("Invalid token")
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return nil, fmt.Errorf("Invalid claims format")
    }

    return claims, nil
}

func secosystem() {
    userID := "123"
    token, err := generateToken(userID)
    if err != nil {
        fmt.Println("Error generating token:", err)
        return
    }
    fmt.Println("Generated token:", token)


	tokens := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTE1NjU5NjYsInVzZXJfaWQiOiIxMjMifQ.pKe_ouml5p6AiLntHcDjvJL1lTo5CnvlmsSpwI2Tcgo"
    claims, err := validateToken(tokens)
    if err != nil {
        fmt.Println("Token validation error:", err)
        return
    }
    userIDs := claims["user_id"].(string)
    fmt.Println("Valid token for user:", userIDs)
}
