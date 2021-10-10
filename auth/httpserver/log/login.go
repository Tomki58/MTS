package log

import (
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "incorret request body", http.StatusBadRequest)
	}

	w.Write(reqBody)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Successfully logout!"))
}
