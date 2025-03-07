package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Weapon struct {
	UUID         string  `json:"uuid"`
	DisplayName  string  `json:"displayName"`
	Description  string  `json:"description"`
	DisplayIcon  string  `json:"displayIcon"`
	Type         string  `json:"category"`
	FireRate     float64 `json:"fireRate"`
	MagazineSize int     `json:"magazineSize"`
	ReloadTime   float64 `json:"reloadTimeSeconds"`
}

func FetchWeapons() ([]Weapon, error) {
	apiURL := "https://valorant-api.com/v1/weapons"
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Data []struct {
			UUID        string `json:"uuid"`
			DisplayName string `json:"displayName"`
			Description string `json:"description"`
			DisplayIcon string `json:"displayIcon"`
			Category    string `json:"category"`
			WeaponStats struct {
				FireRate     float64 `json:"fireRate"`
				MagazineSize int     `json:"magazineSize"`
				ReloadTime   float64 `json:"reloadTimeSeconds"`
			} `json:"weaponStats"`
		} `json:"data"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("Erreur parsing JSON: %v", err)
	}

	var weapons []Weapon
	for _, w := range result.Data {
		weapons = append(weapons, Weapon{
			UUID:         w.UUID,
			DisplayName:  w.DisplayName,
			Description:  w.Description,
			DisplayIcon:  w.DisplayIcon,
			Type:         w.Category,
			FireRate:     w.WeaponStats.FireRate,
			MagazineSize: w.WeaponStats.MagazineSize,
			ReloadTime:   w.WeaponStats.ReloadTime,
		})
	}

	return weapons, nil
}

func GetWeapons() ([]Weapon, error) {
	resp, err := http.Get("https://valorant-api.com/v1/weapons")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Data []Weapon `json:"data"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

func HandleWeapons(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Weapons handler"))
}
