package lineup

import (
    "errors"
    "math/rand"
    "time"
)

func GenerateLineup(players []Player, innings int) (Lineup, error){
    // Not sure if we will keep this as we can't control how many kids make it
    if len(players) < 9 {
        return Lineup{}, errors.New("not enough players")
    }

    // Initialize the game plan
    lineup := Lineup{
        Innings: innings,
        Players: players,
        Defense: make(map[int][]Assignment),
        BattingOrder: generateBattingOrder(players),
    }

    playerAssignments := make(map[int][]Assignment)

    // Assign field positions for each inning
    for inning := 1; inning <= innings; inning++ {
        assignments, err := assignPositionsForInning(players, inning, playerAssignments)
        if err != nil {
            return Lineup{}, err
        }
        lineup.Defense[inning] = assignments

        for _, a := range assignments {
            playerAssignments[a.PlayerId] = append(playerAssignments[a.PlayerId], a)
        }
    }

    // Validate game plan against rules
    if err := ValidateLineup(lineup, players); err != nil {
        return Lineup{}, err
    }

    return lineup, nil
}

func generateBattingOrder(players []Player) []Player {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
	shuffled := make([]Player, len(players))
	copy(shuffled, players)

    r.Shuffle(len(shuffled), func(i, j int) {
        shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
    })

    return shuffled
}

var infieldPositions = []Position{"P", "C", "1B", "2B", "SS", "3B", "SF"}
var outfieldPositions = []Position{"LF", "LCF", "RCF", "RF"}
var benchPosition Position = Bench

func assignPositionsForInning(players []Player, inning int, history map[int][]Assignment) ([]Assignment, error){
	assignments := []Assignment{}
	usedPositions := map[Position]bool{}
	infieldCount := 0
	outfieldCount := 0

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	shuffled := make([]Player, len(players))
	copy(shuffled, players)
	r.Shuffle(len(shuffled), func(i, j int){
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	for _, player := range shuffled {
		// get players past assignments
		past := history[player.Id]
		positionCounts := make(map[Position]int)
		var lastPosition Position

		for _, a := range past {
			positionCounts[a.Position]++
			if a.Inning == inning-1 {
				lastPosition = a.Position
			}
		}

		var assigned Position

		if infieldCount < 7 {
			for _, pos := range infieldPositions {
				if !usedPositions[pos] && positionCounts[pos] < 2 {
					assigned = pos
					infieldCount++
					break
				}
			}
		}

		if assigned == "" && outfieldCount < 5 {
			for _, pos := range outfieldPositions {
				if !usedPositions[pos] && positionCounts[pos] < 2 && lastPosition != pos {
					assigned = pos
					outfieldCount++
					break
				}
			}
		}

		if assigned == "" && lastPosition != benchPosition {
			assigned = benchPosition
		}

		if assigned == "" {
			assigned = infieldPositions[infieldCount%len(infieldPositions)]
		}

		usedPositions[assigned] = true

		assignments = append(assignments, Assignment{
			PlayerId: player.Id,
            PlayerName: player.Name,
			Inning: inning,
			Position: assigned,
		})
	}

	return assignments, nil
}

