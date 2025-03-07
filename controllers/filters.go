package controllers

import (
	"api_valorant/api"
	"html/template"
	"net/http"
	"strings"
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

func HandleFilteredWeapons(w http.ResponseWriter, r *http.Request) {
	// Récupération des filtres depuis l'URL
	weaponType := r.URL.Query().Get("type")
	fireRateFilter := r.URL.Query().Get("fireRate")

	// Récupération de toutes les armes
	weapons, err := api.GetWeapons()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des armes", http.StatusInternalServerError)
		return
	}

	// Filtrage des armes
	var filteredWeapons []api.Weapon
	for _, weapon := range weapons {
		if weaponType != "Tous" && weapon.Type != weaponType {
			continue
		}

		// Filtrage de la cadence de tir
		fireRate := weapon.FireRate
		if fireRateFilter == "above" && fireRate <= 8 {
			continue
		}
		if fireRateFilter == "below" && fireRate > 8 {
			continue
		}

		filteredWeapons = append(filteredWeapons, weapon)
	}

	// Chargement du template et affichage des armes filtrées
	tmpl, err := template.ParseFiles("templates/filtered_weapons.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		http.Error(w, "Erreur lors du chargement du template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, map[string]interface{}{
		"Weapons": filteredWeapons,
	})
}
