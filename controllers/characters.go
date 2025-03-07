package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

const itemsPerPage = 10

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
	Category    string    `json:"category"`
}

type APIAgent struct {
	DisplayName  string `json:"displayName"`
	Description  string `json:"description"`
	FullPortrait string `json:"fullPortrait"`
	Role         struct {
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
		if apiAgent.FullPortrait == "" {
			continue // Skip agents with missing images
		}
		agent := Agent{
			Name:        apiAgent.DisplayName,
			Description: apiAgent.Description,
			Image:       apiAgent.FullPortrait,
			RoleIcon:    apiAgent.Role.DisplayIcon,
			Category:    "Agent", // Set the category to "Agent"
		}
		for _, ability := range apiAgent.Abilities {
			videoURL := ""
			isVideo := false
			agentName := strings.ToLower(agent.Name)
			switch agentName {
			case "sage":
				switch strings.ToLower(ability.DisplayName) {
				case "barrier orb":
					videoURL = "/static/vidéo/sage/barrier_orb.mp4"
					isVideo = true
				case "slow orb":
					videoURL = "/static/vidéo/sage/slow_orb.mp4"
					isVideo = true
				case "healing orb":
					videoURL = "/static/vidéo/sage/healing_orb.mp4"
					isVideo = true
				case "resurrection":
					videoURL = "/static/vidéo/sage/resurrection.mp4"
					isVideo = true
				}
			case "phoenix":
				switch strings.ToLower(ability.DisplayName) {
				case "curveball":
					videoURL = "/static/vidéo/phoenix/curveball.mp4"
					isVideo = true
				case "hot hands":
					videoURL = "/static/vidéo/phoenix/hot_hands.mp4"
					isVideo = true
				case "blaze":
					videoURL = "/static/vidéo/phoenix/blaze.mp4"
					isVideo = true
				case "run it back":
					videoURL = "/static/vidéo/phoenix/run_it_back.mp4"
					isVideo = true
				}
			case "jett":
				switch strings.ToLower(ability.DisplayName) {
				case "cloudburst":
					videoURL = "/static/vidéo/jett/cloudburst.mp4"
					isVideo = true
				case "updraft":
					videoURL = "/static/vidéo/jett/updraft.mp4"
					isVideo = true
				case "tailwind":
					videoURL = "/static/vidéo/jett/tailwind.mp4"
					isVideo = true
				case "blade storm":
					videoURL = "/static/vidéo/jett/blade_storm.mp4"
					isVideo = true
				}
			case "viper":
				switch strings.ToLower(ability.DisplayName) {
				case "poison cloud":
					videoURL = "/static/vidéo/viper/poison_cloud.mp4"
					isVideo = true
				case "toxic screen":
					videoURL = "/static/vidéo/viper/toxic_screen.mp4"
					isVideo = true
				case "snake bite":
					videoURL = "/static/vidéo/viper/snake_bite.mp4"
					isVideo = true
				case "viper's pit":
					videoURL = "/static/vidéo/viper/vipers_pit.mp4"
					isVideo = true
				}
			case "omen":
				switch strings.ToLower(ability.DisplayName) {
				case "shrouded step":
					videoURL = "/static/vidéo/omen/shrouded_step.mp4"
					isVideo = true
				case "paranoia":
					videoURL = "/static/vidéo/omen/paranoia.mp4"
					isVideo = true
				case "dark cover":
					videoURL = "/static/vidéo/omen/dark_cover.mp4"
					isVideo = true
				case "from the shadows":
					videoURL = "/static/vidéo/omen/from_the_shadows.mp4"
					isVideo = true
				}
			case "brimstone":
				switch strings.ToLower(ability.DisplayName) {
				case "incendiary":
					videoURL = "/static/vidéo/brimstone/incendiary.mp4"
					isVideo = true
				case "stim beacon":
					videoURL = "/static/vidéo/brimstone/stim_beacon.mp4"
					isVideo = true
				case "sky smoke":
					videoURL = "/static/vidéo/brimstone/sky_smoke.mp4"
					isVideo = true
				case "orbital strike":
					videoURL = "/static/vidéo/brimstone/orbital_strike.mp4"
					isVideo = true
				}
			case "cypher":
				switch strings.ToLower(ability.DisplayName) {
				case "cyber cage":
					videoURL = "/static/vidéo/cypher/cyber_cage.mp4"
					isVideo = true
				case "spycam":
					videoURL = "/static/vidéo/cypher/spycam.mp4"
					isVideo = true
				case "trapwire":
					videoURL = "/static/vidéo/cypher/trapwire.mp4"
					isVideo = true
				case "neural theft":
					videoURL = "/static/vidéo/cypher/neural_theft.mp4"
					isVideo = true
				}
			case "sova":
				switch strings.ToLower(ability.DisplayName) {
				case "shock bolt":
					videoURL = "/static/vidéo/sova/shock_bolt.mp4"
					isVideo = true
				case "owl drone":
					videoURL = "/static/vidéo/sova/owl_drone.mp4"
					isVideo = true
				case "recon bolt":
					videoURL = "/static/vidéo/sova/recon_bolt.mp4"
					isVideo = true
				case "hunter's fury":
					videoURL = "/static/vidéo/sova/hunters_fury.mp4"
					isVideo = true
				}
			case "reyna":
				switch strings.ToLower(ability.DisplayName) {
				case "leer":
					videoURL = "/static/vidéo/reyna/leer.mp4"
					isVideo = true
				case "devour":
					videoURL = "/static/vidéo/reyna/devour.mp4"
					isVideo = true
				case "dismiss":
					videoURL = "/static/vidéo/reyna/dismiss.mp4"
					isVideo = true
				case "empress":
					videoURL = "/static/vidéo/reyna/empress.mp4"
					isVideo = true
				}
			case "killjoy":
				switch strings.ToLower(ability.DisplayName) {
				case "nanoswarm":
					videoURL = "/static/vidéo/killjoy/nanoswarm.mp4"
					isVideo = true
				case "alarmbot":
					videoURL = "/static/vidéo/killjoy/alarmbot.mp4"
					isVideo = true
				case "turret":
					videoURL = "/static/vidéo/killjoy/turret.mp4"
					isVideo = true
				case "lockdown":
					videoURL = "/static/vidéo/killjoy/lockdown.mp4"
					isVideo = true
				}
			case "skye":
				switch strings.ToLower(ability.DisplayName) {
				case "regrowth":
					videoURL = "/static/vidéo/skye/regrowth.mp4"
					isVideo = true
				case "trailblazer":
					videoURL = "/static/vidéo/skye/trailblazer.mp4"
					isVideo = true
				case "guiding light":
					videoURL = "/static/vidéo/skye/guiding_light.mp4"
					isVideo = true
				case "seekers":
					videoURL = "/static/vidéo/skye/seekers.mp4"
					isVideo = true
				}
			case "yoru":
				switch strings.ToLower(ability.DisplayName) {
				case "fakeout":
					videoURL = "/static/vidéo/yoru/fakeout.mp4"
					isVideo = true
				case "blindside":
					videoURL = "/static/vidéo/yoru/blindside.mp4"
					isVideo = true
				case "gatecrash":
					videoURL = "/static/vidéo/yoru/gatecrash.mp4"
					isVideo = true
				case "dimensional drift":
					videoURL = "/static/vidéo/yoru/dimensional_drift.mp4"
					isVideo = true
				}
			case "astra":
				switch strings.ToLower(ability.DisplayName) {
				case "gravity well":
					videoURL = "/static/vidéo/astra/gravity_well.mp4"
					isVideo = true
				case "nova pulse":
					videoURL = "/static/vidéo/astra/nova_pulse.mp4"
					isVideo = true
				case "nebula":
					videoURL = "/static/vidéo/astra/nebula.mp4"
					isVideo = true
				case "dissipate":
					videoURL = "/static/vidéo/astra/dissipate.mp4"
					isVideo = true
				case "astral form":
					videoURL = "/static/vidéo/astra/astral_form.mp4"
					isVideo = true
				}
			case "kayo":
				switch strings.ToLower(ability.DisplayName) {
				case "frag/ment":
					videoURL = "/static/vidéo/kayo/frag_ment.mp4"
					isVideo = true
				case "flash/drive":
					videoURL = "/static/vidéo/kayo/flash_drive.mp4"
					isVideo = true
				case "zero/point":
					videoURL = "/static/vidéo/kayo/zero_point.mp4"
					isVideo = true
				case "null/cmd":
					videoURL = "/static/vidéo/kayo/null_cmd.mp4"
					isVideo = true
				}
			case "breach":
				switch strings.ToLower(ability.DisplayName) {
				case "aftershock":
					videoURL = "/static/vidéo/breach/aftershock.mp4"
					isVideo = true
				case "flashpoint":
					videoURL = "/static/vidéo/breach/flashpoint.mp4"
					isVideo = true
				case "fault line":
					videoURL = "/static/vidéo/breach/fault_line.mp4"
					isVideo = true
				case "rolling thunder":
					videoURL = "/static/vidéo/breach/rolling_thunder.mp4"
					isVideo = true
				}
			case "gekko":
				switch strings.ToLower(ability.DisplayName) {
				case "wingman":
					videoURL = "/static/vidéo/gekko/wingman.mp4"
					isVideo = true
				case "dizzy":
					videoURL = "/static/vidéo/gekko/dizzy.mp4"
					isVideo = true
				case "mosh pit":
					videoURL = "/static/vidéo/gekko/mosh_pit.mp4"
					isVideo = true
				case "thrash":
					videoURL = "/static/vidéo/gekko/thrash.mp4"
					isVideo = true
				}
			case "iso":
				switch strings.ToLower(ability.DisplayName) {
				case "undercut":
					videoURL = "/static/vidéo/iso/undercut.mp4"
					isVideo = true
				case "double tap":
					videoURL = "/static/vidéo/iso/double_tap.mp4"
					isVideo = true
				case "contingency":
					videoURL = "/static/vidéo/iso/contingency.mp4"
					isVideo = true
				case "kill contract":
					videoURL = "/static/vidéo/iso/kill_contract.mp4"
					isVideo = true
				}
			case "fade":
				switch strings.ToLower(ability.DisplayName) {
				case "prowler":
					videoURL = "/static/vidéo/fade/prowler.mp4"
					isVideo = true
				case "seize":
					videoURL = "/static/vidéo/fade/seize.mp4"
					isVideo = true
				case "haunt":
					videoURL = "/static/vidéo/fade/haunt.mp4"
					isVideo = true
				case "nightfall":
					videoURL = "/static/vidéo/fade/nightfall.mp4"
					isVideo = true
				}
			case "harbor":
				switch strings.ToLower(ability.DisplayName) {
				case "cascade":
					videoURL = "/static/vidéo/harbor/cascade.mp4"
					isVideo = true
				case "cove":
					videoURL = "/static/vidéo/harbor/cove.mp4"
					isVideo = true
				case "high tide":
					videoURL = "/static/vidéo/harbor/high_tide.mp4"
					isVideo = true
				case "reckoning":
					videoURL = "/static/vidéo/harbor/reckoning.mp4"
					isVideo = true
				}
			case "deadlock":
				switch strings.ToLower(ability.DisplayName) {
				case "barrier mesh":
					videoURL = "/static/vidéo/deadlock/barrier_mesh.mp4"
					isVideo = true
				case "gravnet":
					videoURL = "/static/vidéo/deadlock/gravnet.mp4"
					isVideo = true
				case "sonic sensor":
					videoURL = "/static/vidéo/deadlock/sonic_sensor.mp4"
					isVideo = true
				case "anihilation":
					videoURL = "/static/vidéo/deadlock/anihilation.mp4"
					isVideo = true
				}
			case "neon":
				switch strings.ToLower(ability.DisplayName) {
				case "relay_bolt":
					videoURL = "/static/vidéo/neon/relay_bolt.mp4"
					isVideo = true
				case "fast lane":
					videoURL = "/static/vidéo/neon/fast_lane.mp4"
					isVideo = true
				case "overdrive":
					videoURL = "/static/vidéo/neon/overdrive.mp4"
					isVideo = true
				case "high gear":
					videoURL = "/static/vidéo/neon/high_gear.mp4"
					isVideo = true
				}
			case "clove":
				switch strings.ToLower(ability.DisplayName) {
				case "blinding light":
					videoURL = "/static/vidéo/clove/blinding_light.mp4"
					isVideo = true
				case "smoke screen":
					videoURL = "/static/vidéo/clove/smoke_screen.mp4"
					isVideo = true
				case "shadow step":
					videoURL = "/static/vidéo/clove/shadow_step.mp4"
					isVideo = true
				case "shadow form":
					videoURL = "/static/vidéo/clove/shadow_form.mp4"
				}
			case "chamber":
				switch strings.ToLower(ability.DisplayName) {
				case "rendezvous":
					videoURL = "/static/vidéo/chamber/rendezvous.mp4"
					isVideo = true
				case "headhunter":
					videoURL = "/static/vidéo/chamber/headhunter.mp4"
					isVideo = true
				case "trademark":
					videoURL = "/static/vidéo/chamber/trademark.mp4"
					isVideo = true
				case "tour_de_force":
					videoURL = "/static/vidéo/chamber/tour_de_force.mp4"
					isVideo = true
				}
			case "vyse":
				switch strings.ToLower(ability.DisplayName) {
				case "razor vine":
					videoURL = "/static/vidéo/vyse/razor_vine.mp4"
					isVideo = true
				case "shear":
					videoURL = "/static/vidéo/vyse/shear.mp4"
					isVideo = true
				case "arc rose":
					videoURL = "/static/vidéo/vyse/arc_rose.mp4"
					isVideo = true
				case "steel garden":
					videoURL = "/static/vidéo/vyse/steel_garden.mp4"
					isVideo = true
				}
			case "waylay":
				switch strings.ToLower(ability.DisplayName) {
				case "saturation":
					videoURL = "/static/vidéo/waylay/saturation.mp4"
					isVideo = true
				case "vitesse lumière":
					videoURL = "/static/vidéo/waylay/vitesse_lumiere.mp4"
					isVideo = true
				case "réfraction":
					videoURL = "/static/vidéo/waylay/refraction.mp4"
					isVideo = true
				case "croisée des chemins":
					videoURL = "/static/vidéo/waylay/croisee_des_chemins.mp4"
					isVideo = true
				}
				// Add more agents and their abilities here
			default:
				videoURL = "/static/vidéo/" + agentName + "/default.mp4"
				isVideo = true
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

type PaginatedCharacters struct {
	Characters []Agent
	Page       int
	TotalPages int
}

func HandleCharacters(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	agents, err := FetchAgentsFromAPI()
	if err != nil {
		http.Error(w, "Error fetching agents: "+err.Error(), http.StatusInternalServerError)
		return
	}

	totalPages := (len(agents) + itemsPerPage - 1) / itemsPerPage
	start := (page - 1) * itemsPerPage
	end := start + itemsPerPage
	if end > len(agents) {
		end = len(agents)
	}

	paginatedCharacters := PaginatedCharacters{
		Characters: agents[start:end],
		Page:       page,
		TotalPages: totalPages,
	}

	funcMap := template.FuncMap{
		"sub": func(a, b int) int { return a - b },
		"add": func(a, b int) int { return a + b },
	}

	tmpl, err := template.New("characters.html").Funcs(funcMap).ParseFiles("templates/characters.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, paginatedCharacters)
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
	tmpl, err := template.ParseFiles("templates/characters_details.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, agent)
	if err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
	}
}

// Duplicate AddFavoriteHandler function removed
