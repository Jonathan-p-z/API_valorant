package api

// Agent représente un agent dans Valorant
type Agent struct {
    Name        string `json:"displayName"`
    Image       string `json:"fullPortrait"`
    Description string `json:"description"`
}

// Response représente la réponse de l'API contenant une liste d'agents
type Response struct {
    Agents []Agent `json:"agents"`
}