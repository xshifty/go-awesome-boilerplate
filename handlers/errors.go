package handlers

import (
	"fmt"
	"net/http"
)

func errorUnableToLoadResource(w http.ResponseWriter, r *http.Request, e error) {
	http.Error(w, fmt.Sprintf("Unable to load %s: %s", r.URL, e), 500)
}
