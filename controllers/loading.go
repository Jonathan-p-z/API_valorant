package controllers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type Data struct {
	Message string `json:"message"`
}

func LoadData(w http.ResponseWriter, r *http.Request) {
	data := Data{
		Message: "Data loaded successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	// Chargement du template HTML
	tmpl := template.Must(template.ParseFiles("templates/"))

	// Gestionnaire de la page d'accueil
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		tmpl.Execute(w, nil)
	})

	// Gestionnaire de l'API de chargement des données
	http.HandleFunc("/api/load-data", LoadData)

	// Servir les fichiers statiques (CSS et JS)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Démarrage du serveur
	log.Println("Serveur démarré sur http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
