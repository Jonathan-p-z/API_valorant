package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"sync"
)

var (
	dataLoaded = false
	mu         sync.Mutex
)

type Data struct {
	Message string `json:"message"`
}

func InitData() {
	mu.Lock()
	defer mu.Unlock()

	if !dataLoaded {
		// Simulate data loading
		dataLoaded = true
	}
}

func LoadData(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	if !dataLoaded {
		// Simulate data loading
		dataLoaded = true
	}

	data := Data{
		Message: "Data loaded successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func DataLoaded() bool {
	mu.Lock()
	defer mu.Unlock()
	return dataLoaded
}

func LoadingHandler(w http.ResponseWriter, r *http.Request) {
	if DataLoaded() {
		http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
		return
	}

	// Simulate data loading
	InitData()

	tmpl := template.Must(template.ParseFiles("templates/loading.html"))
	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, nil)
}
