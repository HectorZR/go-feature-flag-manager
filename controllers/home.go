package controllers

import (
	"html/template"
	"log"
	"net/http"
)

func HomeController(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")

	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)

	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", 404)
	}
}
