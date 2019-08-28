package handler

import (
	"fmt"
	"net/http"
	"strings"
)

var hostMap = map[string]string{
	"heartbeat.haneru.dev": "hanerubeat.app",
}

func normalizeHost(host string) string {
	if x := hostMap[host]; x != "" {
		return x
	}

	if strings.HasPrefix(host, "www.") {
		return strings.TrimPrefix(host, "www.")
	}

	return host
}

// Handler replies to the request
func Handler(w http.ResponseWriter, r *http.Request) {
	newHost := normalizeHost(r.Host)

	if newHost != r.Host {
		u := fmt.Sprintf("https://%s%s", newHost, r.URL.RequestURI())

		http.Redirect(w, r, u, http.StatusMovedPermanently)
	} else {
		http.NotFound(w, r)
	}
}
