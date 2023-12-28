package handlers

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/xshifty/go-awesome-boilerplate/views/components"
)

func registerComponents(r *chi.Mux, ctx context.Context) {
	r.Route("/components", func(r chi.Router) {
		r.Get("/navigation", NewComponentHandler(ctx, components.Navigation()))
	})
}
