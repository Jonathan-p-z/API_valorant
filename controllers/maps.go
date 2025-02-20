package controllers

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Map struct {
	Name        string `json:"displayName"`
	Description string `json:"description"`
	Image       string `json:"splash"`
}

type PageData struct {
	Title string
	Maps  []Map
}

type APIResponse struct {
	Data []Map `json:"data"`
}

func fetchAPI(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error making HTTP request:", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Non-OK HTTP status:", resp.StatusCode)
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading HTTP response body:", err)
		return err
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
	}
	return err
}

func HandleMaps(w http.ResponseWriter, r *http.Request) {
	var apiResponse APIResponse

	err := fetchAPI("https://valorant-api.com/v1/maps", &apiResponse)
	if err != nil {
		log.Println("Error fetching maps:", err)
		http.Error(w, "Failed to fetch maps", http.StatusInternalServerError)
		return
	}

	data := PageData{
		Title: "Cartes de Valorant",
		Maps:  apiResponse.Data,
	}

	tmpl := template.Must(template.ParseFiles("templates/cartes.html", "templates/header.html"))
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Failed to render page", http.StatusInternalServerError)
	}
}

func HandleMapDetails(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("id")
	var apiResponse APIResponse
	err := fetchAPI("https://valorant-api.com/v1/maps", &apiResponse)
	if err != nil {
		http.Error(w, "Error fetching maps: "+err.Error(), http.StatusInternalServerError)
		return
	}
	maps := apiResponse.Data
	var mapDetail Map
	for _, m := range maps {
		if strings.EqualFold(m.Name, name) {
			mapDetail = m
			break
		}
	}
	if mapDetail.Name == "" {
		http.NotFound(w, r)
		return
	}
	tmpl, err := template.ParseFiles("templates/map_details.html", "templates/header.html")
	if err != nil {
		http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, mapDetail)
	if err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
	}
}
