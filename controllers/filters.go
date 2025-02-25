package controllers

import (
	"html/template"
	"net/http"
	"strings"
	"api_valorant/api"
)

func HandleFilteredCharacters(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	agents, err := api.FetchCharacters()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des agents.", http.StatusInternalServerError)
		return
	}
	filteredCharacters := []api.Character{}
	for _, character := range agents {
		if strings.Contains(strings.ToLower(character.DisplayName), strings.ToLower(query)) {
			filteredCharacters = append(filteredCharacters, character)
		}
	}

	tmpl, err := template.ParseFiles("templates/characters.html")
	if err != nil {
		http.Error(w, "Erreur lors du chargement du template.", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, filteredCharacters)
}
