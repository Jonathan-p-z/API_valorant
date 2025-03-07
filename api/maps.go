package api

import (
	"encoding/json"
	"fmt"
)

type Map struct {
	UUID        string `json:"uuid"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
	Splash      string `json:"splash"`
}

func FetchMaps() ([]Map, error) {
	apiURL := "https://valorant-api.com/v1/maps"
	body, err := FetchData(apiURL)
	if err != nil {
		return nil, err
	}

	var data struct {
		Data []Map `json:"data"`
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("erreur parsing JSON: %v", err)
	}

	return data.Data, nil
}