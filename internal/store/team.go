package store

import "gorm.io/gorm"

type Team struct {
    gorm.Model
    Name string `json:"name"`
    Players []Player `json:"players"`
}

func GetTeamPlayers(teamId int) ([]Player, error){
	var players []Player
	result := DB.Where("Id = ?", teamId).Find(&players)
	return players, result.Error
}

func CreateTheMets(db *gorm.DB) error {
    team := Team{
        Name: "The Mets",
        Players: []Player{
            {Name: "Donny"},
            {Name: "Jaxx"},
            {Name: "Dez"},
            {Name: "Luke"},
            {Name: "Nathan"},
            {Name: "Lucas"},
            {Name: "Andrew"},
            {Name: "Lorenzo"},
            {Name: "Veevaan"},    
            {Name: "Calen"},
            {Name: "Caleb"},
            {Name: "Jonah"},
        },
    }

    return db.Create(&team).Error
}
