package authentications

import (
	"fmt"
	"net/http"
	"time"
	Config "pkg/config"
	"github.com/dgrijalva/jwt-go"
)

var TABLE map[string]interface{}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	tempTable := map[string]interface{} {
		"id_1": "21wdejcwin10dj02inoe3IUG1212w",
		"id_2": "12wqswq",
	}
	TABLE = tempTable

    return func(w http.ResponseWriter, r *http.Request) {
        tokenString := r.Header.Get("Authorization")
		requestField := r.Header.Get("user_id")

        if tokenString == "" || requestField == "" {
            http.Error(w, "Unauthorized 1", http.StatusUnauthorized)
            return
        }

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte(fmt.Sprint(TABLE[requestField])), nil
        })
        if err != nil || !token.Valid {
            http.Error(w, "Unauthorized 2", http.StatusUnauthorized)
			fmt.Println(token, err)
            return
        }

        next(w, r)
    }
}

func GenerateToken(userID string) (string, error) {
	jwtSecret := "9182hd19h82_default_key"

	if userID == Config.REQUEST {
		jwtSecret = Config.KEY
	}
	if userID == Config.REQUEST2 {
		jwtSecret = Config.KEY2
	}
	fmt.Println(userID, jwtSecret)

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Minute * 1).Unix(), // Token expiration time (1 day)
    })
    return token.SignedString([]byte(jwtSecret))
}