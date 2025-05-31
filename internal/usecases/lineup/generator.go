package lineup

import (
	"errors"
	"math/rand"
	"slices"
	"time"
    "sort"
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
var outfieldPositions = []Position{"LF", "LCF", "CF", "RCF", "RF"}
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
		if outfieldCount <= 5 {
            for _, pos := range outfieldPositions {
                // Rules for assigning outfield positions: 
                // 1. Position must not be used already
                // 2. Player must not have played the same position more than twice in the current game
                // 3. Player's last position must not be in the outfield
				if !usedPositions[pos] && positionCounts[pos] < 2 && !slices.Contains(outfieldPositions, lastPosition) {
					assigned = pos
					outfieldCount++
					break
				}
			}
		}

		if assigned == "" && infieldCount <= 7 {
			for _, pos := range infieldPositions {
				if !usedPositions[pos] && positionCounts[pos] < 2 {
					assigned = pos
					infieldCount++
					break
				}
			}
		}

		if assigned == "" && lastPosition != benchPosition && len(usedPositions) > 12 { 
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
    
    sort.SliceStable(assignments, func(i, j int) bool {
        return positionPriority[assignments[i].Position] < positionPriority[assignments[j].Position]
    })

	return assignments, nil
}

var positionPriority = map[Position]int{
    "RF": 1,
    "RCF": 2,
    "CF": 3,
    "LCF": 4,
    "LF": 5,
    "SF": 6,
    "3B": 7,
    "SS": 8,
    "2B": 9,
    "1B": 10,
    "P": 11,
    "C": 12,
}

// func sortDefenseAssignments(assignments []Assignment) []Assignment{
//     sort.SliceStable(assignments, func(i, j int) bool {
//         return positionPriority[assignments[i].Position] < positionPriority[assignments[j].Position]
//     })
// }

