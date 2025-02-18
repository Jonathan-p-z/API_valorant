package controllers

import (
    "encoding/json"
    "html/template"
    "net/http"
)

type Weapon struct {
    Name        string  `json:"displayName"`
    Description string  `json:"description"`
    Image       string  `json:"displayIcon"`
    Type        string  `json:"category"`
    FireRate    float64 `json:"fireRate"`
    MagazineSize int    `json:"magazineSize"`
    ReloadTime  float64 `json:"reloadTime"`
}

type APIWeapon struct {
    DisplayName  string `json:"displayName"`
    Description  string `json:"description"`
    DisplayIcon  string `json:"displayIcon"`
    Category     string `json:"category"`
    WeaponStats struct {
        FireRate     float64 `json:"fireRate"`
        MagazineSize int     `json:"magazineSize"`
        ReloadTime   float64 `json:"reloadTimeSeconds"`
    } `json:"weaponStats"`
}

func FetchWeaponsFromAPI() ([]Weapon, error) {
    resp, err := http.Get("https://valorant-api.com/v1/weapons")
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var result struct {
        Data []APIWeapon `json:"data"`
    }
    err = json.NewDecoder(resp.Body).Decode(&result)
    if err != nil {
        return nil, err
    }

    var weapons []Weapon
    for _, apiWeapon := range result.Data {
        weapon := Weapon{
            Name:        apiWeapon.DisplayName,
            Description: apiWeapon.Description,
            Image:       apiWeapon.DisplayIcon,
            Type:        apiWeapon.Category,
            FireRate:    apiWeapon.WeaponStats.FireRate,
            MagazineSize: apiWeapon.WeaponStats.MagazineSize,
            ReloadTime:  apiWeapon.WeaponStats.ReloadTime,
        }
        weapons = append(weapons, weapon)
    }
    return weapons, nil
}

func HandleWeapons(w http.ResponseWriter, r *http.Request) {
    weapons, err := FetchWeaponsFromAPI()
    if err != nil {
        http.Error(w, "Error fetching weapons: "+err.Error(), http.StatusInternalServerError)
        return
    }

    tmpl, err := template.ParseFiles("templates/weapons.html", "templates/header.html")
    if err != nil {
        http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
        return
    }
    data := struct {
        Weapons []Weapon
    }{
        Weapons: weapons,
    }
    err = tmpl.Execute(w, data)
    if err != nil {
        http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
    }
}