package controllers

import (
    "html/template"
    "net/http"
)

type Weapon struct {
    Name  string `json:"name"`
    Type  string `json:"type"`
    Image string `json:"image"`
}

func GetWeapons() []Weapon {
    return []Weapon{
        {
            Name:  "Vandal",
            Type:  "Fusil d'assaut",
            Image: "https://media.valorant-api.com/weapons/vandal.png",
        },
        {
            Name:  "Phantom",
            Type:  "Fusil d'assaut",
            Image: "https://media.valorant-api.com/weapons/phantom.png",
        },
        {
            Name:  "Operator",
            Type:  "Fusil de précision",
            Image: "https://media.valorant-api.com/weapons/operator.png",
        },
		{
			Name:  "Judge",
			Type:  "Fusil à pompe",
			Image: "https://media.valorant-api.com/weapons/judge.png",
		},

    }
}

func HandleWeapons(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("templates/weapons.html", "templates/header.html")
    if err != nil {
        http.Error(w, "Error loading template", http.StatusInternalServerError)
        return
    }
    data := GetWeapons()
    tmpl.Execute(w, data)
}