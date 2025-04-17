package store

type Player struct {
	Id int `gorm."primaryKey"`
	Name string
}

func GetAllPlayers() ([]Player, error){
	var players []Player
	result := DB.Find(&players)
	return players, result.Error
}
