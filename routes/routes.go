package routes

import (
	"net/http"

	c "hectorzurga.com/feature-flag-manager/controllers"
	m "hectorzurga.com/feature-flag-manager/middlewares"
)

func Setup() {
	http.Handle("/", m.AuthMiddleware(http.HandlerFunc(c.HomeController)))

	http.HandleFunc("POST /counter", c.IncreaseCounterController)

	http.HandleFunc("POST /restart-counter", c.RestartCounterController)
}
