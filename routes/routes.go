package routes

import (
	"net/http"

	"hectorzurga.com/feature-flag-manager/controllers"
)

func Setup() {
	http.HandleFunc("/", controllers.HomeController)
}
