package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"sync"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var (
	users     = make(map[string]User)
	usersLock sync.Mutex
)

const usersFile = "data/users.json"

func init() {
	err := loadUsers()
	if err != nil {
		log.Printf("Error loading users during initialization: %v", err)
	}
}

func loadUsers() error {
	log.Println("Attempting to load users from JSON...")

	data, err := os.ReadFile(usersFile)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("users.json not found, creating a new one.")
			users = make(map[string]User)
			return saveUsers()
		}
		log.Printf("Error reading users.json: %v", err)
		return fmt.Errorf("error reading users file: %w", err)
	}

	tempUsers := make(map[string]User)
	if err := json.Unmarshal(data, &tempUsers); err != nil {
		log.Printf("Error decoding users.json: %v", err)
		users = make(map[string]User)
		return saveUsers()
	}

	users = tempUsers
	log.Println("Users successfully loaded from users.json")
	return nil
}

func saveUsers() error {
	dir := "data"
	if err := os.MkdirAll(dir, 0666); err != nil {
		log.Printf("Error creating directory %s: %v", dir, err)
		return fmt.Errorf("error creating directory: %w", err)
	}
	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		log.Printf("Failed to encode users: %v", err)
		return fmt.Errorf("error encoding users to JSON: %w", err)
	}
	if err := os.WriteFile(usersFile, data, 0666); err != nil {
		log.Printf("Failed to write users file: %v", err)
		return fmt.Errorf("error writing users file: %w", err)
	}
	log.Println("Users successfully saved to users.json")
	return nil
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("templates/login.html", "templates/header.html", "templates/footer.html")
		if err != nil {
			http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		log.Printf("Error parsing form: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	newUser := User{
		Username: r.FormValue("username"),
		Email:    r.FormValue("email"),
		Password: string(hashedPassword),
	}

	log.Printf("Signup attempt: username=%s, email=%s", newUser.Username, newUser.Email)

	usersLock.Lock()
	defer usersLock.Unlock()

	if err := loadUsers(); err != nil {
		log.Printf("Error in SignupHandler when loading users: %v", err)
		http.Error(w, "Error loading users", http.StatusInternalServerError)
		return
	}

	if _, exists := users[newUser.Email]; exists {
		log.Printf("Email already registered: %s", newUser.Email)
		http.Error(w, "Email already registered", http.StatusConflict)
		return
	}

	users[newUser.Email] = newUser

	if err := saveUsers(); err != nil {
		log.Printf("Error in SignupHandler when saving users: %v", err)
		http.Error(w, "Error saving users", http.StatusInternalServerError)
		return
	}

	log.Println("User registered successfully")
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("templates/login.html", "templates/header.html", "templates/footer.html")
		if err != nil {
			http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	credentials := struct {
		Email    string
		Password string
	}{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	usersLock.Lock()
	defer usersLock.Unlock()

	if err := loadUsers(); err != nil {
		log.Printf("Error in LoginHandler when loading users: %v", err)
		http.Error(w, "Error loading users", http.StatusInternalServerError)
		return
	}

	user, exists := users[credentials.Email]
	if !exists || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)) != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	usersLock.Lock()
	defer usersLock.Unlock()

	if err := loadUsers(); err != nil {
		log.Printf("Error in GetUsersHandler when loading users: %v", err)
		http.Error(w, "Error loading users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		log.Printf("Error in GetUsersHandler when encoding users: %v", err)
		http.Error(w, "Error encoding users", http.StatusInternalServerError)
	}
}
