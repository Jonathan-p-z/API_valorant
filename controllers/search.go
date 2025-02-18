package controllers

import (
	"html/template"
	"net/http"
	"strings"
	"api_valorant/api"
)

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("search")
	agents, err := api.FetchCharacters()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des agents.", http.StatusInternalServerError)
		return
	}
	searchResults := []api.Agent{}
	for _, agent := range agents {
		if strings.Contains(strings.ToLower(agent.Name), strings.ToLower(query)) {
			searchResults = append(searchResults, agent)
		}
	}

	tmpl, err := template.ParseFiles("templates/search.html")
	if err != nil {
		http.Error(w, "Erreur lors du chargement du template.", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, searchResults)
}
