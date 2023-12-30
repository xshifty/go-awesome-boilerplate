package handlers

import (
	"context"
	"net/http"

	"github.com/xshifty/go-awesome-boilerplate/pkg/components"
	"github.com/xshifty/go-awesome-boilerplate/pkg/layouts"
	"github.com/xshifty/go-awesome-boilerplate/pkg/pages"
)

var ctx = context.Background()

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	Render(ctx, layouts.Base(pages.Home()), w)
}

func homeContentsHandler(w http.ResponseWriter, _ *http.Request) {
	Render(ctx, pages.Home(), w)
}

func aboutUsHandler(w http.ResponseWriter, _ *http.Request) {
	Render(ctx, layouts.Base(pages.AboutUs()), w)
}

func aboutUsContentsHandler(w http.ResponseWriter, _ *http.Request) {
	Render(ctx, pages.AboutUs(), w)
}

func navigationComponentHandler(w http.ResponseWriter, _ *http.Request) {
	Render(ctx, components.Navigation(), w)
}
