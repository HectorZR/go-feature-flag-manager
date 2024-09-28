package main

import (
	"log"
	"net/http"

	"hectorzurga.com/feature-flag-manager/routes"
)

func main() {
	routes.Setup()

	log.Println("Server is running on port 8000")

	log.Fatal(http.ListenAndServe(":8000", nil))
}
