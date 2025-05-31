package store

type Assignment struct {
    PlayerId uint `json:"player_id"`
    Position string `json:"position"`
    Count int `json:"count"`
}

func GetPlayersAssignments(player_id int) ([]Assignment, error){
    var assignments []Assignment
	result := DB.Where("player_id = ?", player_id).Find(&assignments)
	return assignments, result.Error
}
