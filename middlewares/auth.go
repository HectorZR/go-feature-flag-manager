package middlewares

import (
	"log"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Auth middleware executed")
		next.ServeHTTP(w, r)
	})
}
