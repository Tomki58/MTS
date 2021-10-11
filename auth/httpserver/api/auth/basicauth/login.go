package basicauth

import (
	"bytes"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type BasicAuthorizator struct {
	credentials map[string]string
}

func New(pathToConfig string) (*BasicAuthorizator, error) {
	data, err := os.ReadFile(pathToConfig)
	if err != nil {
		return nil, err
	}

	creds := make(map[string]string)
	for _, auth := range bytes.Split(data, []byte{'\n'}) {
		cs := bytes.IndexByte(auth, ':')
		if cs < 0 {
			continue
		}
		username, password := auth[:cs], auth[cs+1:]
		creds[string(username)] = string(password)
	}

	return &BasicAuthorizator{
		credentials: creds,
	}, nil
}

func (b *BasicAuthorizator) Login(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if !ok {
		http.Error(w, "Incorrect Authorization Header", http.StatusForbidden)
		return
	}

	if ok := b.CheckCredentials(username, password); !ok {
		http.Error(w, "Invalid Credentials", http.StatusForbidden)
		return
	}

	// TODO: change the response on sending JWT-token to client
	// Generating access cookie
	accessToken := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"username": username,
	})
	accessTokenString, err := accessToken.SigningString()

	if err != nil {
		http.Error(w, "Can't generate JWT-token", http.StatusInternalServerError)
		return
	}
	accessCookie := http.Cookie{
		Name: "access",
		// login as jwt token
		Value:    accessTokenString,
		MaxAge:   int(time.Minute),
		HttpOnly: true,
	}

	// TODO: generate refresh cookie
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"username": username,
	})
	refreshTokenString, err := refreshToken.SigningString()
	if err != nil {
		http.Error(w, "Can't generate JWT-token", http.StatusInternalServerError)
		return
	}

	refreshCookie := http.Cookie{
		Name:     "refresh",
		Value:    refreshTokenString,
		MaxAge:   int(time.Hour),
		HttpOnly: true,
	}
	http.SetCookie(w, &accessCookie)
	http.SetCookie(w, &refreshCookie)
	w.WriteHeader(http.StatusAccepted)
}

func (b *BasicAuthorizator) CheckCredentials(username, password string) bool {
	if _, ok := b.credentials[username]; ok {
		return true
	}

	return false
}
