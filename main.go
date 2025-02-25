package main

import (
    "api_valorant/controllers"
    "html/template"
    "log"
    "net/http"
)

func main() {
    // Initialize data at startup
    controllers.InitData()

    http.Handle("/characters", CheckDataLoaded(http.HandlerFunc(controllers.HandleCharacters)))
    http.Handle("/characters/search", CheckDataLoaded(http.HandlerFunc(controllers.HandleSearch)))
    http.Handle("/characters/filter", CheckDataLoaded(http.HandlerFunc(controllers.HandleFilteredCharacters)))
    http.Handle("/weapons", CheckDataLoaded(http.HandlerFunc(controllers.HandleWeapons)))
    http.Handle("/maps", CheckDataLoaded(http.HandlerFunc(controllers.HandleMaps)))
    http.Handle("/maps/details", CheckDataLoaded(http.HandlerFunc(controllers.HandleMapDetails)))
    http.Handle("/search", CheckDataLoaded(http.HandlerFunc(controllers.HandleSearch)))
    http.Handle("/fav", CheckDataLoaded(http.HandlerFunc(controllers.ListFavorites)))

    http.HandleFunc("/home", homeHandler)
    http.HandleFunc("/auth", authHandler)
    http.HandleFunc("/login", loginHandler)
    http.HandleFunc("/loading", controllers.LoadingHandler)
    http.HandleFunc("/api/favorites", controllers.AddFavorite)
    http.HandleFunc("/api/remove-favorite", controllers.RemoveFavorite)
    http.Handle("/character/", CheckDataLoaded(http.HandlerFunc(controllers.HandleCharacterDetails)))

    // Serve static files correctly
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    log.Println("Server started on http://localhost:8080")
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

func CheckDataLoaded(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Check if data is loaded
        if !controllers.DataLoaded() {
            http.Redirect(w, r, "/loading", http.StatusTemporaryRedirect)
            return
        }
        next.ServeHTTP(w, r)
    })
}