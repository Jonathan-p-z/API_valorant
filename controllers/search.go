package controllers

import (
	"api_valorant/api"
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type SearchResult struct {
	Name        string
	Image       string
	Link        string
	Description string
	Percentage  float64
}

type Config struct {
	PercentageAcceptable float64
	CacheDuration        time.Duration
	MinimumPercentage    float64
}

var (
	config        Config
	cachedAgents  []api.Character
	cachedWeapons []api.Weapon
	cachedMaps    []api.Map
	lastFetchTime time.Time
)

func init() {
	config.PercentageAcceptable = getEnvAsFloat("PERCENTAGE_ACCEPTABLE", 60.0)
	config.CacheDuration = getEnvAsDuration("CACHE_DURATION", 5*time.Minute)
}

func getEnvAsFloat(key string, defaultValue float64) float64 {
	if value, exists := os.LookupEnv(key); exists {
		if floatValue, err := strconv.ParseFloat(value, 64); err == nil {
			return floatValue
		}
	}
	return defaultValue
}

func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	if value, exists := os.LookupEnv(key); exists {
		if durationValue, err := time.ParseDuration(value); err == nil {
			return durationValue
		}
	}
	return defaultValue
}

func levenshteinDistance(s1, s2 string) int {
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)
	lenS1 := len(s1)
	lenS2 := len(s2)

	if lenS1 == 0 {
		return lenS2
	}
	if lenS2 == 0 {
		return lenS1
	}

	prevRow := make([]int, lenS2+1)
	for j := 0; j <= lenS2; j++ {
		prevRow[j] = j
	}

	for i := 1; i <= lenS1; i++ {
		currentRow := make([]int, lenS2+1)
		currentRow[0] = i

		for j := 1; j <= lenS2; j++ {
			cost := 0
			if s1[i-1] != s2[j-1] {
				cost = 1
			}
			currentRow[j] = min(currentRow[j-1]+1, min(prevRow[j]+1, prevRow[j-1]+cost))
		}

		prevRow = currentRow
	}

	return prevRow[lenS2]
}

func similarityPercentage(s1, s2 string) float64 {
	distance := levenshteinDistance(s1, s2)
	maxLen := max(len(s1), len(s2))
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
	startTime := time.Now()
	log.Println("HandleSearch called")

	query := strings.TrimSpace(strings.ToLower(r.URL.Query().Get("search")))
	if query == "" {
		http.Error(w, "Le paramètre 'search' est requis.", http.StatusBadRequest)
		return
	}
	log.Printf("Search query: %s", query)
	if time.Since(lastFetchTime) > config.CacheDuration {
		var errAgents, errWeapons, errMaps error
		cachedAgents, errAgents = api.FetchCharacters()
		cachedWeapons, errWeapons = api.FetchWeapons()
		cachedMaps, errMaps = api.FetchMaps()

		if errAgents != nil || errWeapons != nil || errMaps != nil {
			log.Printf("Erreur lors de la récupération des données: agents: %v, weapons: %v, maps: %v", errAgents, errWeapons, errMaps)
			http.Error(w, "Erreur lors de la récupération des données.", http.StatusInternalServerError)
			return
		}
		lastFetchTime = time.Now()
	}

	searchResults := searchAll(query, cachedAgents, cachedWeapons, cachedMaps)
	log.Printf("Search results: %v", searchResults)

	tmpl, err := template.ParseFiles("templates/search.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		log.Printf("Erreur lors du chargement du template: %v", err)
		http.Error(w, "Erreur lors du chargement du template.", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, searchResults); err != nil {
		log.Printf("Erreur lors de l'exécution du template: %v", err)
		http.Error(w, "Erreur lors de l'exécution du template.", http.StatusInternalServerError)
	}

	log.Printf("Temps d'exécution de HandleSearch: %v", time.Since(startTime))
}

func searchAll(query string, agents []api.Character, weapons []api.Weapon, maps []api.Map) []SearchResult {
	var searchResults []SearchResult

	addResult := func(name, image, description, link string, percentage float64) {
		searchResults = append(searchResults, SearchResult{
			Name:        name,
			Image:       image,
			Description: description,
			Link:        link,
			Percentage:  percentage,
		})
	}
	switch query {
	case "agents", "characters":
		for _, agent := range agents {
			addResult(agent.DisplayName, agent.DisplayIcon, agent.Description, "/character/"+agent.DisplayName, 100.0)
		}
	case "weapons":
		for _, weapon := range weapons {
			addResult(weapon.DisplayName, weapon.DisplayIcon, weapon.Description, "/weapons", 100.0)
		}
	case "maps":
		for _, m := range maps {
			addResult(m.DisplayName, m.Splash, m.Description, "/maps/details?id="+m.UUID, 100.0)
		}
	default:
		for _, agent := range agents {
			if percentage := similarityPercentage(query, agent.DisplayName); percentage > 30.0 {
				addResult(agent.DisplayName, agent.DisplayIcon, agent.Description, "/character/"+agent.DisplayName, percentage)
			}
		}
		for _, weapon := range weapons {
			if percentage := similarityPercentage(query, weapon.DisplayName); percentage > 30.0 {
				addResult(weapon.DisplayName, weapon.DisplayIcon, weapon.Description, "/weapons", percentage)
			}
		}
		for _, m := range maps {
			if percentage := similarityPercentage(query, m.DisplayName); percentage > 30.0 {
				addResult(m.DisplayName, m.Splash, m.Description, "/maps/details?id="+m.UUID, percentage)
			}
		}
	}
	sort.Slice(searchResults, func(i, j int) bool {
		return searchResults[i].Percentage > searchResults[j].Percentage
	})

	return searchResults
}

func init() {
	config.PercentageAcceptable = getEnvAsFloat("PERCENTAGE_ACCEPTABLE", 60.0)
	config.CacheDuration = getEnvAsDuration("CACHE_DURATION", 5*time.Minute)
	config.MinimumPercentage = getEnvAsFloat("MINIMUM_PERCENTAGE", 30.0)
}
