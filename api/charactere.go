package api

import (
	"encoding/json"
	"fmt"
)

// FetchCharacters récupère les agents depuis l'API
func FetchCharacters() ([]Agent, error) {
	apiURL := "https://valorant-api.com/v1/agents"
	body, err := FetchData(apiURL)
	if err != nil {
		return nil, err
	}

	var data struct {
		Data []Agent `json:"data"`
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("Erreur parsing JSON : %v", err)
	}

	return data.Data, nil
}
