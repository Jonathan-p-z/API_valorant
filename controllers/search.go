package controllers

import (
	"api_valorant/api"
	"html/template"
	"log"
	"net/http"
	"sort"
	"strings"
)

type SearchResult struct {
	Name        string
	Image       string
	Link        string
	Description string
	Percentage  float64
}

func similarityPercentage(str1, str2 string) float64 {
	lenStr1 := len(str1)
	lenStr2 := len(str2)
	if lenStr1 == 0 {
		return float64(lenStr2)
	}
	if lenStr2 == 0 {
		return float64(lenStr1)
	}
	matrix := make([][]int, lenStr1+1)
	for i := range matrix {
		matrix[i] = make([]int, lenStr2+1)
	}
	for i := 0; i <= lenStr1; i++ {
		matrix[i][0] = i
	}
	for j := 0; j <= lenStr2; j++ {
		matrix[0][j] = j
	}
	for i := 1; i <= lenStr1; i++ {
		for j := 1; j <= lenStr2; j++ {
			cost := 0
			if str1[i-1] != str2[j-1] {
				cost = 1
			}
			matrix[i][j] = min(matrix[i-1][j]+1, min(matrix[i][j-1]+1, matrix[i-1][j-1]+cost))
		}
	}
	distance := matrix[lenStr1][lenStr2]
	maxLen := max(lenStr1, lenStr2)
	return (1.0 - float64(distance)/float64(maxLen)) * 100
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleSearch called")
	query := strings.ToLower(r.URL.Query().Get("search"))
	log.Printf("Search query: %s", query)

	agents, err := api.FetchCharacters()
	if err != nil {
		log.Printf("Error fetching characters: %v", err)
		http.Error(w, "Erreur lors de la récupération des agents.", http.StatusInternalServerError)
		return
	}

	weapons, err := api.FetchWeapons()
	if err != nil {
		log.Printf("Error fetching weapons: %v", err)
		http.Error(w, "Erreur lors de la récupération des armes.", http.StatusInternalServerError)
		return
	}

	maps, err := api.FetchMaps()
	if err != nil {
		log.Printf("Error fetching maps: %v", err)
		http.Error(w, "Erreur lors de la récupération des cartes.", http.StatusInternalServerError)
		return
	}

	searchResults := searchAll(query, agents, weapons, maps)
	log.Printf("Search results: %v", searchResults)

	tmpl, err := template.ParseFiles("templates/search.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		log.Printf("Error loading template: %v", err)
		http.Error(w, "Erreur lors du chargement du template.", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, searchResults)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Erreur lors de l'exécution du template.", http.StatusInternalServerError)
	}
}

func searchAll(query string, agents []api.Character, weapons []api.Weapon, maps []api.Map) []SearchResult {
	var PercentageAcceptable float64 = 60.0
	var searchResults []SearchResult

	for _, agent := range agents {
		percentage := similarityPercentage(query, agent.DisplayName)
		if percentage > PercentageAcceptable {
			searchResults = append(searchResults, SearchResult{
				Name:        agent.DisplayName,
				Image:       agent.DisplayIcon,
				Description: agent.Description,
				Link:        "/characters/details?id=" + agent.UUID,
				Percentage:  percentage,
			})
		}
	}

	for _, weapon := range weapons {
		percentage := similarityPercentage(query, weapon.DisplayName)
		if percentage > PercentageAcceptable {
			searchResults = append(searchResults, SearchResult{
				Name:        weapon.DisplayName,
				Image:       weapon.DisplayIcon,
				Description: weapon.Description,
				Link:        "/weapons/details?id=" + weapon.UUID,
				Percentage:  percentage,
			})
		}
	}

	for _, m := range maps {
		percentage := similarityPercentage(query, m.DisplayName)
		if percentage > PercentageAcceptable {
			searchResults = append(searchResults, SearchResult{
				Name:        m.DisplayName,
				Image:       m.Splash,
				Description: m.Description,
				Link:        "/maps/details?id=" + m.UUID,
				Percentage:  percentage,
			})
		}
	}

	// Sort results by similarity percentage in descending order
	sort.Slice(searchResults, func(i, j int) bool {
		return searchResults[i].Percentage > searchResults[j].Percentage
	})

	return searchResults
}
