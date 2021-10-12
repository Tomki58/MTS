package httpserver

import (
	"MTS/auth/httpserver/api/auth"
	"MTS/auth/httpserver/api/auth/basicauth"
	"MTS/auth/httpserver/api/i"
	"MTS/auth/httpserver/middlewares"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Server struct {
	S    http.Server
	Auth auth.Authorizator
}

func New() (*Server, error) {
	authenticator, err := basicauth.New("./auth/credentials.txt")
	if err != nil {
		return nil, err
	}

	router := chi.NewRouter()

	// TODO: add middleware to refresh cookies
	// append middlewares
	router.Use(middleware.DefaultLogger)
	// login endpoint
	router.Group(func(r chi.Router) {
		r.Get("/login", authenticator.Login)
	})

	router.Group(func(r chi.Router) {
		r.Use(middlewares.HandleJWT)
		r.Use(middlewares.UpdateCookies)
		r.Get("/i", i.Identification)
		r.Get("/me", me.)
	})

	httpServer := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	return &Server{
		S:    httpServer,
		Auth: authenticator,
	}, nil

}
