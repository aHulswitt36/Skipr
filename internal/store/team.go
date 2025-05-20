package store

type Team struct {
	Id int `gorm."primaryKey"`
	Name string
	Players []Player
}

func GetTeamPlayers() ([]Player, error){
	var players []Player
	result := DB.Find(&players)
	return players, result.Error
}
