package api

type Agent struct {
    Name        string `json:"displayName"`
    Image       string `json:"fullPortrait"`
    Description string `json:"description"`
}

type Response struct {
    Agents []Agent `json:"agents"`
}