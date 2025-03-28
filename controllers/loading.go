package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"sync"
	"time"
)

var (
	dataLoaded bool
	mu         sync.Mutex
)

type Data struct {
	Message string `json:"message"`
}

func InitData() {
	mu.Lock()
	defer mu.Unlock()

	if !dataLoaded {
		time.Sleep(3 * time.Second) // Simulate data loading
		dataLoaded = true
	}
}

func LoadData(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	if !dataLoaded {
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
	tmpl := template.Must(template.ParseFiles("templates/loading.html"))
	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, nil)

	go func() {
		InitData()
	}()

	for {
		if DataLoaded() {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
}
