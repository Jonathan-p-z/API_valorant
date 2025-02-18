package controllers

import (
    "encoding/json"
    "html/template"
    "net/http"
    "strings"
)

type Ability struct {
    Name        string `json:"displayName"`
    Description string `json:"description"`
    Image       string `json:"displayIcon"`
    VideoURL    string `json:"videoURL"`
}

type Agent struct {
    Name        string    `json:"displayName"`
    Description string    `json:"description"`
    Image       string    `json:"fullPortrait"`
    RoleIcon    string    `json:"roleIcon"`
    Images      []string  `json:"images"`
    Passive     string    `json:"passive"`
    Abilities   []Ability `json:"abilities"`
}

type APIAgent struct {
    DisplayName string `json:"displayName"`
    Description string `json:"description"`
    FullPortrait string `json:"fullPortrait"`
    Role struct {
        DisplayIcon string `json:"displayIcon"`
    } `json:"role"`
    Abilities []struct {
        DisplayName string `json:"displayName"`
        Description string `json:"description"`
        DisplayIcon string `json:"displayIcon"`
    } `json:"abilities"`
}

func FetchAgentsFromAPI() ([]Agent, error) {
    resp, err := http.Get("https://valorant-api.com/v1/agents")
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var result struct {
        Data []APIAgent `json:"data"`
    }
    err = json.NewDecoder(resp.Body).Decode(&result)
    if err != nil {
        return nil, err
    }

    var agents []Agent
    for _, apiAgent := range result.Data {
        agent := Agent{
            Name:        apiAgent.DisplayName,
            Description: apiAgent.Description,
            Image:       apiAgent.FullPortrait,
            RoleIcon:    apiAgent.Role.DisplayIcon,
        }
        for _, ability := range apiAgent.Abilities {
            agent.Abilities = append(agent.Abilities, Ability{
                Name:        ability.DisplayName,
                Description: ability.Description,
                Image:       ability.DisplayIcon,
                VideoURL:    "https://www.youtube.com/embed/" + getYouTubeVideoID(ability.DisplayName), // Remplacez par l'URL de la vidéo appropriée
            })
        }
        agents = append(agents, agent)
    }
    return agents, nil
}

func getYouTubeVideoID(abilityName string) string {
    // Remplacez cette fonction par une logique pour obtenir l'ID de la vidéo YouTube appropriée pour chaque compétence
    switch abilityName {
    case "Cloudburst":
        return "dQw4w9WgXcQ" // Remplacez par l'ID de la vidéo YouTube appropriée
    case "Updraft":
        return "dQw4w9WgXcQ" // Remplacez par l'ID de la vidéo YouTube appropriée
    case "Tailwind":
        return "dQw4w9WgXcQ" // Remplacez par l'ID de la vidéo YouTube appropriée
    case "Blade Storm":
        return "dQw4w9WgXcQ" // Remplacez par l'ID de la vidéo YouTube appropriée
    default:
        return ""
    }
}

func HandleCharacters(w http.ResponseWriter, r *http.Request) {
    agents, err := FetchAgentsFromAPI()
    if err != nil {
        http.Error(w, "Error fetching agents: "+err.Error(), http.StatusInternalServerError)
        return
    }

    tmpl, err := template.ParseFiles("templates/characters.html", "templates/header.html")
    if err != nil {
        http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
        return
    }
    data := struct {
        Agents []Agent
    }{
        Agents: agents,
    }
    err = tmpl.Execute(w, data)
    if err != nil {
        http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
    }
}

func HandleCharacterDetails(w http.ResponseWriter, r *http.Request) {
    name := strings.TrimPrefix(r.URL.Path, "/character/")
    agents, err := FetchAgentsFromAPI()
    if err != nil {
        http.Error(w, "Error fetching agents: "+err.Error(), http.StatusInternalServerError)
        return
    }

    var agent Agent
    for _, a := range agents {
        if strings.ToLower(a.Name) == strings.ToLower(name) {
            agent = a
            break
        }
    }
    if agent.Name == "" {
        http.NotFound(w, r)
        return
    }

    tmpl, err := template.ParseFiles("templates/characters_details.html", "templates/header.html")
    if err != nil {
        http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
        return
    }
    err = tmpl.Execute(w, agent)
    if err != nil {
        http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
    }
}