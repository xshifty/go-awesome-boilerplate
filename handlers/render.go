package handlers

import (
	"context"
	"github.com/a-h/templ"
	"net/http"
)

func NewComponentHandler(ctx context.Context, c templ.Component) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := c.Render(ctx, w); err != nil {
			errorUnableToLoadResource(w, r, err)
		}
	}
}
