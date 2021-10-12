package cookies

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte("MyBaseKey")

type MyClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

// CreateCookie generates new Cookie for user containing JWT-token of username
func CreateCookies(userToken jwt.Token) ([]*http.Cookie, error) {
	tokenString, err := userToken.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	accessCookie := http.Cookie{
		Name:     "access",
		Value:    tokenString,
		MaxAge:   int(time.Minute / 1e9),
		HttpOnly: true,
	}

	refreshCookie := http.Cookie{
		Name:     "refresh",
		Value:    tokenString,
		MaxAge:   int(time.Hour / 1e9),
		HttpOnly: true,
	}

	return []*http.Cookie{&accessCookie, &refreshCookie}, nil
}
