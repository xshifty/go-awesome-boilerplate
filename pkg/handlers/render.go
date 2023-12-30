package handlers

import (
	"context"
	"errors"
	"fmt"
	"github.com/a-h/templ"
	"github.com/xshifty/go-awesome-boilerplate/pkg/layouts"
	"net/http"
)

func Render(ctx context.Context, c templ.Component, w http.ResponseWriter) {
	if err := c.Render(ctx, w); err != nil {
		err = layouts.Error(errors.Join(
			errors.New("internal server error"),
			fmt.Errorf("cannot render component"),
			err,
		), fmt.Sprint(http.StatusInternalServerError)).Render(ctx, w)

		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		}
	}
}

