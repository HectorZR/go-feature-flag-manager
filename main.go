package main

import (
	"fmt"
	"log"
	"net/http"

	"hectorzurga.com/feature-flag-manager/routes"
)

func main() {
	routes.Setup()

	port := 8000
	log.Println("Server is running on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
