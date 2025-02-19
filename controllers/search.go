package controllers

import (
	"api_valorant/api"
	"html/template"
	"log"
	"net/http"
	"sort"
)

type SearchResult struct {
	Name       string
	Image      string
	Link       string
	Percentage float64
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
	query := r.URL.Query().Get("search")
	log.Printf("Search query: %s", query)
	agents, err := api.FetchCharacters()
	if err != nil {
		log.Printf("Error fetching characters: %v", err)
		http.Error(w, "Erreur lors de la récupération des agents.", http.StatusInternalServerError)
		return
	}
	searchResults := searchAgents(query, agents)
	log.Printf("Search results: %v", searchResults)

	tmpl, err := template.ParseFiles("templates/search.html")
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

func searchAgents(query string, agents []api.Agent) []SearchResult {
	var PercentageAcceptable float64 = 60.0
	var temp SearchResult
	var ListOfPercentage []float64
	var searchResults []SearchResult = []SearchResult{}

	for _, agent := range agents {
		ListOfPercentage = []float64{}
		temp = SearchResult{
			Name:  agent.Name,
			Image: agent.Image,
			Link:  "/characters/details?id=" + agent.Name,
		}
		ListOfPercentage = append(ListOfPercentage, similarityPercentage(query, agent.Name))
		ListOfPercentage = append(ListOfPercentage, similarityPercentage(query, agent.Description))
		sort.Slice(ListOfPercentage, func(i, j int) bool {
			return ListOfPercentage[i] > ListOfPercentage[j]
		})
		if ListOfPercentage[0] > PercentageAcceptable {
			temp.Percentage = ListOfPercentage[0]
			searchResults = append(searchResults, temp)
		}
	}
	return searchResults
}
