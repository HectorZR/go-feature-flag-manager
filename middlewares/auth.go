package middlewares

import (
	"net/http"

	"hectorzurga.com/feature-flag-manager/constants"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(constants.AUTH_COOKIE_NAME)

		if err != nil || cookie.Value != "true" {
			http.Redirect(w, r, "/", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
