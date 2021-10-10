package httpserver

import (
	"net/http"

	login "auth/httpserver/log"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// New creates new server with parameters
func New(versionFile string) (*http.Server, error) {
	router := chi.NewRouter()
	// router.Use(middleware.Logger)
	router.Use(middleware.BasicAuth())

	// login endpoint
	router.Post("/login", login.Login)
	// logout endpoint
	router.Get("/logout", login.Logout)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	return &server, nil
}
