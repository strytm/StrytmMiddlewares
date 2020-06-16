package middlewares

import (
	"net/http"
)

func XssProtectMiddleware() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			w.Header().Set("X-XSS-Protection", "1; mode=block")
			w.Header().Set("X-Frame-Options", "deny")

			h.ServeHTTP(w, r)
		})
	}
}
