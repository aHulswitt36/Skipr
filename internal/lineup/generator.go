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

    playerAssignments := make(map[string][]Assignment)

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
    if err := ValidateLineup(lineup); err != nil {
        return Lineup{}, err
    }

    return lineup, nil
}

func generateBattingOrder(players []Player) []string {
    ids := make([]string, len(players))
    for i, p := range players {
        ids[i] = p.Id
    }

    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    r.Shuffle(len(ids), func(i, j int) {
        ids[i], ids[j] = ids[j], ids[i]
    })

    return ids
}
