package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
)

type ContextKey string

const ContextUserKey ContextKey = "user"

type User struct {
	Username string
}

type MyClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

// HandleJWT unmarshal JWT-token in Request
func HandleJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := new(User)
		refreshCookie, err := r.Cookie("refresh")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tokenString := refreshCookie.Value
		username, err := validateJWT(tokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		user.Username = username.(string)

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "user", username)))
	})
}

// validateJWTN validates JWT-token and returns the token
func validateJWT(tokenString string) (interface{}, error) {

	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("MyBaseKey"), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims.Username, nil
	} else {
		return nil, fmt.Errorf("Invalid JWT-token")
	}
}
