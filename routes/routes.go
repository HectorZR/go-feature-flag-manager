package routes

import (
	"net/http"

	c "hectorzurga.com/feature-flag-manager/controllers"
	m "hectorzurga.com/feature-flag-manager/middlewares"
)

func Setup() {
	http.HandleFunc("/", c.HomeController)

	http.HandleFunc("GET /login", c.LoginView)
	http.HandleFunc("POST /login", c.LoginController)

	http.Handle("POST /logout", m.AuthMiddleware(http.HandlerFunc(c.LogoutController)))

	http.HandleFunc("POST /counter", c.IncreaseCounterController)

	http.HandleFunc("POST /restart-counter", c.RestartCounterController)
}
