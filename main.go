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
	http.HandleFunc("/login", controllers.LoginHandler)
	http.HandleFunc("/signup", controllers.SignupHandler)
	http.HandleFunc("/loading", controllers.LoadingHandler)
	http.HandleFunc("/api/favorites", controllers.AddFavorite)
	http.HandleFunc("/api/remove-favorite", controllers.RemoveFavorite)
	http.HandleFunc("/api/users", controllers.GetUsersHandler) // New route for getting users
	http.Handle("/character/", CheckDataLoaded(http.HandlerFunc(controllers.HandleCharacterDetails)))

	// Serve static files correctly
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Custom error handlers
	http.HandleFunc("/400", errorHandler(400))
	http.HandleFunc("/401", errorHandler(401))
	http.HandleFunc("/403", errorHandler(403))
	http.HandleFunc("/404", errorHandler(404))
	http.HandleFunc("/405", errorHandler(405))
	http.HandleFunc("/408", errorHandler(408))
	http.HandleFunc("/429", errorHandler(429))
	http.HandleFunc("/500", errorHandler(500))
	http.HandleFunc("/502", errorHandler(502))
	http.HandleFunc("/503", errorHandler(503))
	http.HandleFunc("/504", errorHandler(504))

	// Log the exact URL when the server starts
	serverAddress := "http://localhost:8080/loading"
	log.Printf("Server started on %s", serverAddress)
	log.Println("You can access the application at:", serverAddress)

	http.ListenAndServe("localhost:8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/accueil.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/auth.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/login.html", "templates/header.html", "templates/footer.html")
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

func errorHandler(status int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		tmpl, err := template.ParseFiles(
			"templates/header.html",
			"templates/footer.html",
			"templates/error.html",
		)
		if err != nil {
			http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.ExecuteTemplate(w, "error.html", map[string]interface{}{
			"Status": status,
		})
	}
}
