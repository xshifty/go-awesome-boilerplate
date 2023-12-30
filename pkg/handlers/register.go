package handlers

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type HandleOpts struct {
	useLoggerMiddleware bool
	staticDir           string
	staticPrefix        string
	ctx                 context.Context
}

type OptsFunc func(opts *HandleOpts)

func Register(optsFunc ...OptsFunc) *chi.Mux {
	opts := HandleOpts{
		useLoggerMiddleware: false,
		staticDir:           "./static",
		staticPrefix:        "/static/",
		ctx:                 context.Background(),
	}

	for _, fn := range optsFunc {
		fn(&opts)
	}

	r := chi.NewRouter()
	fs := http.FileServer(http.Dir(opts.staticDir))

	if opts.useLoggerMiddleware {
		r.Use(middleware.Logger)
	}

	r.Handle(fmt.Sprintf("%s*", opts.staticPrefix), http.StripPrefix(opts.staticPrefix, fs))

	registerPages(r, opts.ctx)
	registerComponents(r, opts.ctx)

	return r
}

func WithLoggerEnabled(enabled bool) OptsFunc {
	return func(opts *HandleOpts) {
		opts.useLoggerMiddleware = enabled
	}
}

func WithStaticDir(dir string) OptsFunc {
	return func(opts *HandleOpts) {
		opts.staticDir = dir
	}
}

func WithStaticPrefix(endpoint string) OptsFunc {
	return func(opts *HandleOpts) {
		opts.staticPrefix = endpoint
	}
}

func WithContext(ctx context.Context) OptsFunc {
	return func(opts *HandleOpts) {
		opts.ctx = ctx
	}
}
