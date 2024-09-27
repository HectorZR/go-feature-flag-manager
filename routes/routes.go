package routes

import (
	"net/http"

	"hectorzurga.com/feature-flag-manager/controllers"
)

func Setup() {
	http.HandleFunc("/", controllers.HomeController)

	http.HandleFunc("POST /counter", controllers.IncreaseCounterController)

	http.HandleFunc("POST /restart-counter", controllers.RestartCounterController)
}
