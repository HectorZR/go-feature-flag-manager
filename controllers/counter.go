package controllers

import (
	"html/template"
	"log"
	"net/http"
)

type count struct {
	Count int
}

func (c *count) increase() {
	c.Count += 1
}

func (c *count) restart() {
	c.Count = 0
}

var counter = count{
	Count: 0,
}

func IncreaseCounterController(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/counter.html")

	if err != nil {
		log.Println(err)
		return
	}

	counter.increase()

	t.Execute(w, counter)
}

func RestartCounterController(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/counter.html")

	if err != nil {
		log.Println(err)
		return
	}

	counter.restart()

	t.Execute(w, counter)
}
