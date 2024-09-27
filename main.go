package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Data struct {
	Name string
}

func handler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	t.Execute(w, Data{Name: "John Doe"})
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Server is running on port 8000")

	log.Fatal(http.ListenAndServe(":8000", nil))
}
