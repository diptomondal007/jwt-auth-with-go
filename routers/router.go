package routers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"jwtauth/api"
)

func InitRoutes(handler api.AuthHandler) *chi.Mux{
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/register",handler.RegisterPost)
	r.Post("/login",handler.LoginPost)
	r = SetHelloRoutes(r)
	return r
}
