package api

import (
	"encoding/json"
	"fmt"
)

type Weapon struct {
	Name  string `json:"displayName"`
	Image string `json:"displayIcon"`
}

// FetchWeapons récupère les armes depuis l'API
func FetchWeapons() ([]Weapon, error) {
	apiURL := "https://valorant-api.com/v1/weapons"
	body, err := FetchData(apiURL)
	if err != nil {
		return nil, err
	}

	var data struct {
		Data []Weapon `json:"data"`
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("Erreur parsing JSON : %v", err)
	}

	return data.Data, nil
}
