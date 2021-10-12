package i

import (
	"net/http"
)

type ContextKey string

const ContextUserKey ContextKey = "user"

func Identification(w http.ResponseWriter, r *http.Request) {
	if name, ok := r.Context().Value("user").(string); ok {
		w.Write([]byte(name))
	} else {
		http.Error(w, "Cannot get username", http.StatusInternalServerError)
		return
	}
}
