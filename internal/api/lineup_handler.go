package api

import (
	"encoding/json"
	"fmt"
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

	if error := lineup.ValidateLineup(plan, players); error != nil {
		http.Error(w, "Validation failed: " + error.Error(), http.StatusBadRequest)
		return
	}
	
	json.NewEncoder(w).Encode(plan)
}
