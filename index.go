package handler

import (
	"fmt"
	"net/http"
	"strings"
)

// Handler replies to the request
func Handler(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.Host, "www.") {
		host := strings.TrimPrefix(r.Host, "www.")
		u := fmt.Sprintf("https://%s%s", host, r.URL.RequestURI())

		http.Redirect(w, r, u, http.StatusMovedPermanently)
	} else {
		http.NotFound(w, r)
	}
}
