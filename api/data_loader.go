package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type APIConfig struct {
	API string `json:"api"`
}

type CombinedData struct {
	Agents  APIConfig `json:"agents"`
	Weapons APIConfig `json:"weapons"`
	Maps    APIConfig `json:"maps"`
	Roles   APIConfig `json:"roles"`
	Skins   APIConfig `json:"skins"`
}

func LoadCombinedData(filePath string) (*CombinedData, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("Erreur lors de la lecture du fichier JSON : %v", err)
	}

	var combined CombinedData
	err = json.Unmarshal(data, &combined)
	if err != nil {
		return nil, fmt.Errorf("Erreur lors du parsing JSON : %v", err)
	}

	return &combined, nil
}

func FetchDataFromAPI(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Erreur lors de la requête API : %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Erreur HTTP %d : %s", resp.StatusCode, resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Erreur lors de la lecture de la réponse : %v", err)
	}

	return body, nil
}
