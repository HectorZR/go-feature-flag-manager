package routes

import (
	c "hectorzurga.com/feature-flag-manager/controllers"
	m "hectorzurga.com/feature-flag-manager/middlewares"
)

func Setup() {
	var router Router

	router.Get("/", c.HomeController)
	router.Get("/login", c.LoginView)
	router.Post("/login", c.LoginController)
	router.Post("/logout", m.AuthMiddleware(c.LogoutController))
	router.Post("/counter", c.IncreaseCounterController)
	router.Post("/restart-counter", c.RestartCounterController)
}
