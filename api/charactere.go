package api

import (
    "encoding/json"
    "fmt"
    
)

type Character struct {
	UUID        string `json:"uuid"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
	DeveloperName string `json:"developerName"`
	CharacterTags []string `json:"characterTags"`
	DisplayIcon string `json:"displayIcon"`
	BustPortrait string `json:"bustPortrait"`
	FullPortrait string `json:"fullPortrait"`
	KillfeedPortrait string `json:"killfeedPortrait"`
	Background string `json:"background"`
	AssetPath string `json:"assetPath"`
}

func FetchCharacters() ([]Character, error) {
	apiURL := "https://valorant-api.com/v1/agents?isPlayableCharacter=true"
	body, err := FetchData(apiURL)
	if err != nil {
		return nil, err
	}

	var data struct {
		Data []Character `json:"data"`
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("Erreur parsing JSON : %v", err)
	}

	return data.Data, nil
}