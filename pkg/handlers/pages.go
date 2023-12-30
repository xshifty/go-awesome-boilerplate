package handlers

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/xshifty/go-awesome-boilerplate/pkg/layouts"
	"github.com/xshifty/go-awesome-boilerplate/pkg/pages"
)

func registerPages(r *chi.Mux, ctx context.Context) {
	r.Get("/", NewComponentHandler(ctx, layouts.Base(pages.Home())))
	r.Get("/about-us", NewComponentHandler(ctx, layouts.Base(pages.AboutUs())))

	r.Route("/pages", func(r chi.Router) {
		r.Get("/home", NewComponentHandler(ctx, pages.Home()))
		r.Get("/about-us", NewComponentHandler(ctx, pages.AboutUs()))
	})
}
