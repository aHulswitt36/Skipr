package lineup

import (
	"errors"
	"fmt"
)

func ValidateLineup(lineup LineUp, players []Player) error {
	playerMap := map[string]Player{}
	for _, p := range players {
		playerMap[p.Id] = p
	}

	playerPositionCount := map[string]map[string]int{}
	lastPositions := map[string]string{}
	consecOutOrBench := map[string]int{}
	infieldPlayed := map[string]bool{}

	for inning := 1; inning <= 6; inning++ {
		infielders := 0
		outfielders := 0

		assignments := lineup.Lineup[inning]
		for _, a := range assignments {
			playerId := a.PlayerId
			pos := a.Position

			if _, ok := playerPositionCount[playerId]; !ok {
				playerPositionCount[playerId] = map[string]int{}
			}
		}
	}
}
