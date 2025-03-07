package controllers

import (
	"api_valorant/api"
	"encoding/json"
	"html/template"
	"net/http"
)

type Weapon struct {
	Name         string  `json:"displayName"`
	Description  string  `json:"description"`
	Image        string  `json:"displayIcon"`
	Type         string  `json:"category"`
	FireRate     float64 `json:"fireRate"`
	MagazineSize int     `json:"magazineSize"`
	ReloadTime   float64 `json:"reloadTime"`
}

type APIWeapon struct {
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
	DisplayIcon string `json:"displayIcon"`
	Category    string `json:"category"`
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
			Name:         apiWeapon.DisplayName,
			Description:  apiWeapon.Description,
			Image:        apiWeapon.DisplayIcon,
			Type:         apiWeapon.Category,
			FireRate:     apiWeapon.WeaponStats.FireRate,
			MagazineSize: apiWeapon.WeaponStats.MagazineSize,
			ReloadTime:   apiWeapon.WeaponStats.ReloadTime,
		}
		weapons = append(weapons, weapon)
	}
	return weapons, nil
}

func HandleFilteredWeaponsByType(w http.ResponseWriter, r *http.Request) {
	// Récupération des paramètres du formulaire
	weaponType := r.URL.Query().Get("type")
	fireRateFilter := r.URL.Query().Get("fireRate")

	// Récupération de toutes les armes depuis l'API
	weapons, err := api.GetWeapons()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des armes", http.StatusInternalServerError)
		return
	}

	// Filtrage des armes en fonction des critères
	filteredWeapons := []api.Weapon{}
	for _, weapon := range weapons {
		if weaponType != "Tous" && weapon.Type != weaponType {
			continue
		}

		if fireRateFilter == "above" && weapon.FireRate <= 8 {
			continue
		}
		if fireRateFilter == "below" && weapon.FireRate >= 8 {
			continue
		}

		filteredWeapons = append(filteredWeapons, weapon)
	}

	// Chargement du template
	tmpl, err := template.ParseFiles("templates/filtered_weapons.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		http.Error(w, "Erreur de chargement du template", http.StatusInternalServerError)
		return
	}

	// Exécution du template avec les données filtrées
	err = tmpl.Execute(w, struct {
		Weapons []api.Weapon
	}{
		Weapons: filteredWeapons,
	})
	if err != nil {
		http.Error(w, "Erreur lors de l'affichage des armes filtrées", http.StatusInternalServerError)
		return
	}
}

func WeaponErrorHandler(statusCode int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		tmpl, err := template.ParseFiles("templates/error.html", "templates/header.html", "templates/footer.html")
		if err != nil {
			http.Error(w, "Erreur lors du chargement du template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	}
}

func getWeaponTypeName(weaponType string) string {
	switch weaponType {
	case "EEquippableCategory::Sidearm":
		return "Armes de poing"
	case "EEquippableCategory::SMG":
		return "PM"
	case "EEquippableCategory::Shotgun":
		return "Fusils à pompe"
	case "EEquippableCategory::LMG":
		return "Mitrailleuses"
	case "EEquippableCategory::Melee":
		return "Mêlée"
	case "EEquippableCategory::Rifle":
		return "Fusils"
	case "EEquippableCategory::Sniper":
		return "Snipers"
	default:
		return "Autres"
	}
}

func HandleWeapons(w http.ResponseWriter, r *http.Request) {
	weapons, err := api.FetchWeapons()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des armes.", http.StatusInternalServerError)
		return
	}

	// Categorize weapons by type
	categorizedWeapons := map[string][]api.Weapon{
		"EEquippableCategory::Sidearm": {},
		"EEquippableCategory::SMG":     {},
		"EEquippableCategory::Shotgun": {},
		"EEquippableCategory::LMG":     {},
		"EEquippableCategory::Melee":   {},
		"EEquippableCategory::Rifle":   {},
		"EEquippableCategory::Sniper":  {},
	}

	for _, weapon := range weapons {
		categorizedWeapons[weapon.Type] = append(categorizedWeapons[weapon.Type], weapon)
	}

	funcMap := template.FuncMap{
		"getWeaponTypeName": getWeaponTypeName,
	}

	tmpl, err := template.New("weapons.html").Funcs(funcMap).ParseFiles("templates/weapons.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		http.Error(w, "Erreur lors du chargement du template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, map[string]interface{}{
		"CategorizedWeapons": categorizedWeapons,
	})
	if err != nil {
		http.Error(w, "Erreur lors de l'exécution du template: "+err.Error(), http.StatusInternalServerError)
	}
}
