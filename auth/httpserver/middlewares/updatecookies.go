package middlewares

import (
	"MTS/auth/httpserver/cookies"
	"net/http"

	"github.com/golang-jwt/jwt"
)

func UpdateCookies(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if access is expired then the length of r.Cookies == 1
		if len(r.Cookies()) == 1 && r.Cookies()[0].Name == "refresh" {
			// get cookie value
			refreshCookie, err := r.Cookie("refresh")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			tokenString := refreshCookie.Value
			token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
				return []byte("MyBaseKey"), nil
			})
			newCookies, err := cookies.CreateCookies(*token)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			r.AddCookie(newCookies[0])
			r.AddCookie(newCookies[1])
		}
		next.ServeHTTP(w, r)
	})
}
