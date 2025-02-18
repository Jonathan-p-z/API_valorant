package controllers

import (
    "encoding/json"
    "net/http"
    "sync"
)

type User struct {
    Username    string `json:"username"`
    Email       string `json:"email"`
    PhoneNumber string `json:"phone_number"`
    Password    string `json:"password"`
}

var (
    users     = make(map[string]User)
    usersLock sync.Mutex
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var newUser User
    err := json.NewDecoder(r.Body).Decode(&newUser)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    usersLock.Lock()
    defer usersLock.Unlock()

    if _, exists := users[newUser.Email]; exists {
        http.Error(w, "Email already registered", http.StatusConflict)
        return
    }

    users[newUser.Email] = newUser
    w.WriteHeader(http.StatusCreated)
    http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var credentials struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    err := json.NewDecoder(r.Body).Decode(&credentials)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    usersLock.Lock()
    defer usersLock.Unlock()

    user, exists := users[credentials.Email]
    if !exists || user.Password != credentials.Password {
        http.Error(w, "Invalid email or password", http.StatusUnauthorized)
        return
    }

    http.Redirect(w, r, "/home", http.StatusSeeOther)
}