package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchData(apiURL string) ([]byte, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("Erreur API : %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Erreur HTTP %d : %s", resp.StatusCode, resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Erreur lecture réponse : %v", err)
	}

	return body, nil
}

func FetchAgents() ([]Agent, error) {
	const url = "https://valorant-api.com/v1/agents"
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la requête API : %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erreur HTTP %d : %s", resp.StatusCode, resp.Status)
	}
	var result struct {
		Data []Agent `json:"data"`
	}
	return result.Data, nil
}
