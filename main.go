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

    http.HandleFunc("/home", homeHandler)
    http.HandleFunc("/auth", authHandler)
    http.HandleFunc("/login", loginHandler)
    http.HandleFunc("/character/", controllers.HandleCharacterDetails) // Nouvelle route pour les détails du personnage

    // Servir les fichiers statiques correctement
    fs := http.FileServer(http.Dir("./Static"))
    http.Handle("/Static/", http.StripPrefix("/Static/", fs))

    http.ListenAndServe("localhost:8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("templates/accueil.html", "templates/header.html")
    if err != nil {
        http.Error(w, "Error loading template", http.StatusInternalServerError)
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