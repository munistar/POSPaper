package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(propertyHandler *PropertyHandler) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api/v1/properties", func(rt chi.Router) {
		rt.Mount("/", propertyHandler.Routes())
	})

	return r
}
