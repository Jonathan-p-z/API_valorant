package api

// Agent représente un agent dans Valorant
type Agent struct {
    Name        string `json:"name"`
    Description string `json:"description"`
    Image       string `json:"image"`
}

// Response représente la réponse de l'API contenant une liste d'agents
type Response struct {
    Agents []Agent `json:"agents"`
}