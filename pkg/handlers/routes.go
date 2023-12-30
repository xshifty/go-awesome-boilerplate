package handlers

import (
	"context"
	"github.com/go-chi/chi/v5"
)

func registerRoutes(r *chi.Mux, ctx context.Context) {
	r.Get("/", homeHandler)
	r.Get("/pages/home", homeContentsHandler)

	r.Get("/about-us", aboutUsHandler)
	r.Get("/pages/about-us", aboutUsContentsHandler)

	// components
	r.Get("/components/navigation", navigationComponentHandler)
}
