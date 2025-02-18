package controllers

import (
	"html/template"
	"net/http"
	"api_valorant/api"
)

func HandleWeapons(w http.ResponseWriter, r *http.Request) {
	weapons, err := api.FetchWeapons()
	if err != nil {
		http.Error(w, "Erreur chargement armes.", http.StatusInternalServerError)
		return
	}
	tmpl, err := template.ParseFiles("templates/weapons.html")
	if err != nil {
		http.Error(w, "Erreur chargement template.", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, weapons)
}
