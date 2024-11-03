package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/Molina-Q/svelte-wiki/backend/internal/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)

	r.Route("/private", func(router chi.Router) {

		router.Use(middleware.Authorization)
	})

	router.Get("/status", getServerStatus)
}