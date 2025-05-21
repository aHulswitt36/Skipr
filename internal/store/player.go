package store

import "gorm.io/gorm"

type Player struct {
    gorm.Model
    Name string `json:"name"`
    TeamID uint `json:"team_id"`
}

func GetAllPlayers() ([]Player, error){
	var players []Player
	result := DB.Find(&players)
	return players, result.Error
}
