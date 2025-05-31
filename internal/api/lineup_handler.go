package api

import (
	"encoding/json"
	"net/http"

	"skipr/internal/usecases/lineup"
)

var players []lineup.Player
var playerId int = 1

func PlayerHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(players)
	case http.MethodPost:
		var p lineup.Player
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
        p.Id = playerId
        playerId++
		players = append(players, p)
		w.WriteHeader(http.StatusCreated)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func PlayersHandler(w http.ResponseWriter, r *http.Request) {
    var p lineup.Player
    p.Id = playerId
    p.Name = "Donny"
    playerId++
    players = append(players, p)

    p.Id = playerId
    p.Name = "Veevan"
    playerId++
    players = append(players, p)

    p.Id = playerId
    p.Name = "Jonah"
    playerId++
    players = append(players, p)

    p.Id = playerId
    p.Name = "Caleb"
    playerId++
    players = append(players, p)

    p.Id = playerId
    p.Name = "Calen"
    playerId++
    players = append(players, p)

    p.Id = playerId
    p.Name = "Jaxx"
    playerId++
    players = append(players, p)

    p.Id = playerId
    p.Name = "Dez"
    playerId++
    players = append(players, p)

    p.Id = playerId
    p.Name = "Nate"
    playerId++
    players = append(players, p)

    p.Id = playerId
    p.Name = "Luke"
    playerId++
    players = append(players, p)

    p.Id = playerId
    p.Name = "Andrew"
    playerId++
    players = append(players, p)

    p.Id = playerId
    p.Name = "Lucas"
    playerId++
    players = append(players, p)

    p.Id = playerId
    p.Name = "Lorenzo"
    playerId++
    players = append(players, p)

    w.WriteHeader(http.StatusCreated)
}

func LineupHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	plan, err := lineup.GenerateLineup(players, 6)	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	json.NewEncoder(w).Encode(plan)
}
