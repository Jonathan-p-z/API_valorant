package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"sync"
)

var (
	favoritesFile = "data/favorites.json"
	mu            sync.Mutex
)

type Favorite struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
	Type  string `json:"type"`
}

func loadFavorites() ([]Favorite, error) {
	file, err := os.Open(favoritesFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Favorite{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var favorites []Favorite
	if err := json.NewDecoder(file).Decode(&favorites); err != nil {
		return nil, err
	}
	return favorites, nil
}

func saveFavorites(favorites []Favorite) error {
	file, err := os.Create(favoritesFile)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(favorites)
}

func AddFavorite(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var favorite Favorite
	if err := json.NewDecoder(r.Body).Decode(&favorite); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	favorites, err := loadFavorites()
	if err != nil {
		http.Error(w, "Failed to load favorites", http.StatusInternalServerError)
		return
	}

	favorites = append(favorites, favorite)
	if err := saveFavorites(favorites); err != nil {
		http.Error(w, "Failed to save favorite", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func AddFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	favorite := Favorite{
		ID:    r.FormValue("id"),
		Name:  r.FormValue("name"),
		Image: r.FormValue("image"),
		Type:  r.FormValue("type"),
	}

	favorites, err := loadFavorites()
	if err != nil {
		http.Error(w, "Failed to load favorites", http.StatusInternalServerError)
		return
	}

	favorites = append(favorites, favorite)
	if err := saveFavorites(favorites); err != nil {
		http.Error(w, "Failed to save favorite", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/fav", http.StatusSeeOther)
}

func RemoveFavorite(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing favorite ID", http.StatusBadRequest)
		return
	}

	favorites, err := loadFavorites()
	if err != nil {
		http.Error(w, "Failed to load favorites", http.StatusInternalServerError)
		return
	}

	for i, favorite := range favorites {
		if favorite.ID == id {
			favorites = append(favorites[:i], favorites[i+1:]...)
			break
		}
	}

	if err := saveFavorites(favorites); err != nil {
		http.Error(w, "Failed to save favorites", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func ListFavorites(w http.ResponseWriter, r *http.Request) {
	favorites, err := loadFavorites()
	if err != nil {
		http.Error(w, "Failed to load favorites", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/fav.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, favorites)
}
