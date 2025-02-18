package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Image struct {
	URL string `json:"url"`
}

func FetchImage(apiURL string) (string, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return "", fmt.Errorf("Erreur lors de la requête à l'API : %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Erreur HTTP %d : %s", resp.StatusCode, resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Erreur lors de la lecture de la réponse : %v", err)
	}

	var image Image
	err = json.Unmarshal(body, &image)
	if err != nil {
		return "", fmt.Errorf("Erreur lors du parsing JSON : %v", err)
	}

	return image.URL, nil
}
