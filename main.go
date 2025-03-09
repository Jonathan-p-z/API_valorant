package main

import (
	"api_valorant/controllers"
	"html/template"
	"log"
	"net/http"
)

func main() {
	controllers.InitData()

	http.Handle("/characters", CheckDataLoaded(http.HandlerFunc(controllers.HandleCharacters)))
	http.Handle("/characters/search", CheckDataLoaded(http.HandlerFunc(controllers.HandleSearch)))
	http.Handle("/characters/filter", CheckDataLoaded(http.HandlerFunc(controllers.HandleFilteredCharacters)))
	http.Handle("/weapons", CheckDataLoaded(http.HandlerFunc(controllers.HandleWeapons)))
	http.HandleFunc("/filtered_weapons", controllers.HandleFilteredWeapons)
	http.Handle("/maps", CheckDataLoaded(http.HandlerFunc(controllers.HandleMaps)))
	http.Handle("/maps/details", CheckDataLoaded(http.HandlerFunc(controllers.HandleMapDetails)))
	http.Handle("/search", CheckDataLoaded(http.HandlerFunc(controllers.HandleSearch)))
	http.Handle("/fav", CheckDataLoaded(http.HandlerFunc(controllers.ListFavorites)))

	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/auth", authHandler)
	http.HandleFunc("/login", controllers.LoginHandler)
	http.HandleFunc("/signup", controllers.SignupHandler)
	http.HandleFunc("/loading", controllers.LoadingHandler)
	http.HandleFunc("/load-data", controllers.LoadData)
	http.HandleFunc("/api/favorites", controllers.AddFavorite)
	http.HandleFunc("/api/remove-favorite", controllers.RemoveFavorite)
	http.HandleFunc("/api/users", controllers.GetUsersHandler)
	http.Handle("/character/", CheckDataLoaded(http.HandlerFunc(controllers.HandleCharacterDetails)))
	http.HandleFunc("/add-favorite", controllers.AddFavoriteHandler)
	http.HandleFunc("/about", aboutHandler)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/400", controllers.ErrorHandler(400))
	http.HandleFunc("/401", controllers.ErrorHandler(401))
	http.HandleFunc("/403", controllers.ErrorHandler(403))
	http.HandleFunc("/404", controllers.ErrorHandler(404))
	http.HandleFunc("/405", controllers.ErrorHandler(405))
	http.HandleFunc("/408", controllers.ErrorHandler(408))
	http.HandleFunc("/429", controllers.ErrorHandler(429))
	http.HandleFunc("/500", controllers.ErrorHandler(500))
	http.HandleFunc("/502", controllers.ErrorHandler(502))
	http.HandleFunc("/503", controllers.ErrorHandler(503))
	http.HandleFunc("/504", controllers.ErrorHandler(504))

	http.HandleFunc("/", controllers.ErrorHandler(404))

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

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/about.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func CheckDataLoaded(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !controllers.DataLoaded() {
			http.Redirect(w, r, "/loading", http.StatusTemporaryRedirect)
			return
		}
		next.ServeHTTP(w, r)
	})
}