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
    }

    // Generate a valid batting order
    lineup.BattingOrder = generateBattingOrder(players)

    // Assign field positions for each inning
    for inning := 1; inning <= innings; inning++ {
        assignments := assignPositionsForInning(players, inning, lineup)
        lineup.Defense[inning] = assignments
    }

    // Validate game plan against rules
    if err := ValidateLineup(lineup); err != nil {
        return Lineup{}, err
    }

    return lineup, nil
}
