package api

import (
    "encoding/json"
    "net/http"
)

func FetchCharacters() ([]Agent, error) {
    resp, err := http.Get("https://valorant-api.com/v1/agents")
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var result Response
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, err
    }

    return result.Agents, nil
}