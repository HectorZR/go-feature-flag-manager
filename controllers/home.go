package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")

	if err != nil {
		fmt.Println(err)
		return
	}

	t.Execute(w, nil)
}
