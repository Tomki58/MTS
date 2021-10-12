package basicauth

import (
	"bytes"
	"net/http"
	"os"

	"MTS/auth/httpserver/cookies"

	"github.com/golang-jwt/jwt"
)

type BasicAuthorizator struct {
	credentials map[string]string
}

type MyClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyClaims{
		Username: username,
	})

	cookies, err := cookies.CreateCookies(*token)
	if err != nil {
		w.Write([]byte("Failed"))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, cookies[0])
	http.SetCookie(w, cookies[1])
	w.WriteHeader(http.StatusAccepted)
}

func (b *BasicAuthorizator) CheckCredentials(username, password string) bool {
	if _, ok := b.credentials[username]; ok {
		return true
	}

	return false
}
