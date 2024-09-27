package main

import (
	"fmt"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello world!</h1>")
}

func main() {
	http.HandleFunc("/", indexHandler)

	fmt.Println("Server is running on port 8000")

	log.Fatal(http.ListenAndServe(":8000", nil))
}
