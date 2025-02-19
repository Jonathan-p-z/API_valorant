package main

import (
	"api_valorant/controllers"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/characters", controllers.HandleCharacters)
	http.HandleFunc("/characters/search", controllers.HandleSearch)
	http.HandleFunc("/characters/filter", controllers.HandleFilteredCharacters)
	http.HandleFunc("/weapons", controllers.HandleWeapons)
	http.HandleFunc("/maps", controllers.HandleMaps)     // New route for maps
	http.HandleFunc("/search", controllers.HandleSearch) // New route for search

	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/auth", authHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/loading", loadingHandler)                        // New route for loading page
	http.HandleFunc("/character/", controllers.HandleCharacterDetails) // New route for character details

	// Serve static files correctly
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe("localhost:8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/accueil.html", "templates/header.html")
	if err != nil {
		http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/auth.html", "templates/header.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/login.html", "templates/header.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func loadingHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/loading.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
