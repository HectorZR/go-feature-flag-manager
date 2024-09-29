package controllers

import (
	"html/template"
	"log"
	"net/http"

	"hectorzurga.com/feature-flag-manager/constants"
)

func LoginView(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/login.html")

	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)

	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusNotFound)
	}
}

func LoginController(w http.ResponseWriter, r *http.Request) {
	log.Print("Validate user")

	cookie := http.Cookie{
		Name:     constants.AUTH_COOKIE_NAME,
		Value:    "true", // use token here
		Path:     "/",
		MaxAge:   3600,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)

	t, err := template.ParseFiles("templates/dashboard.html")

	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)

	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusNotFound)
	}
}

func LogoutController(w http.ResponseWriter, r *http.Request) {
	log.Print("Logout user")

	http.SetCookie(w, &http.Cookie{
		Name:  constants.AUTH_COOKIE_NAME,
		Value: "false",
	})

	t, err := template.ParseFiles("templates/login.html")

	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)

	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusNotFound)
	}
}
