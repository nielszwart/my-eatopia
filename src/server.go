package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

var templateData TemplateData

func initServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/logout", logoutHandler)
	mux.HandleFunc("/dashboard", dashboardHandler)
	mux.HandleFunc("/recepten", recipesHandler)
	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if isLoggedIn() {
		http.Redirect(w, r, "/dashboard", 302)
	}

	if r.Method == "POST" {
		r.ParseForm()
		email := r.FormValue("email")
		password := r.FormValue("password")

		if email != "" && password != "" {
			if loginUser(email, password) {
				http.Redirect(w, r, "/dashboard", 302)
				return
			}
		}
	}

	render(w, "templates/index.html", nil)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	user = User{}
	http.Redirect(w, r, "/", 302)
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "templates/dashboard.html", templateData)
}

func recipesHandler(w http.ResponseWriter, r *http.Request) {
	templateData.Recipes = getRecipesForUser(user.UserID)
	render(w, "templates/recipes.html", templateData)
}

var mainTemplates = template.Must(template.ParseGlob("templates/main/*.html"))

func render(w http.ResponseWriter, filepath string, data interface{}) {
	tmpl, err := template.Must(mainTemplates.Clone()).ParseFiles(filepath)
	if err != nil {
		log.Println(err)
		http.Error(w, "Something went wrong parsing the template", http.StatusInternalServerError)
	}

	if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
		log.Println(err)
		http.Error(w, "Something went wrong executing the template", http.StatusInternalServerError)
	}
}

func isLoggedIn() bool {
	return (User{}) != user
}
