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
    IsVideo     bool   `json:"isVideo"`
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
            videoURL := ""
            isVideo := false
            if strings.ToLower(agent.Name) == "sage" {
                switch strings.ToLower(ability.DisplayName) {
                case "barrier orb":
                    videoURL = "/static/video/barrier_orb.png"
                    isVideo = true
                case "slow orb":
                    videoURL = "/static/vidéo/slow_orb.mp4"
                    isVideo = true
                case "healing orb":
                    videoURL = "/static/img/healing_orb.png"
                case "resurrection":
                    videoURL = "/static/img/resurrection.png"
                }
            }
            agent.Abilities = append(agent.Abilities, Ability{
                Name:        ability.DisplayName,
                Description: ability.Description,
                Image:       ability.DisplayIcon,
                VideoURL:    videoURL,
                IsVideo:     isVideo,
            })
        }
        agents = append(agents, agent)
    }
    return agents, nil
}
    

/*func getYouTubeVideoID(abilityName string) string {
    switch abilityName {
    case "Cloudburst":
        return "dQw4w9WgXcQ"
    case "Updraft":
        return "dQw4w9WgXcQ"
    case "Tailwind":
        return "dQw4w9WgXcQ"
    case "Blade Storm":
        return "dQw4w9WgXcQ" // Remplacez par l'ID de la vidéo YouTube appropriée
    default:
        return ""
    }
}*/

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
        if strings.EqualFold(a.Name, name) {
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