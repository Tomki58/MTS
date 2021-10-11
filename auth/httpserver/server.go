package httpserver

import (
	"MTS/auth/httpserver/api/auth"
	"MTS/auth/httpserver/api/auth/basicauth"
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
	// append middlewares
	router.Use(middleware.Logger)
	// login endpoint
	router.Get("/login", authenticator.Login)

	httpServer := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	return &Server{
		S:    httpServer,
		Auth: authenticator,
	}, nil

}
