package store

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() error {
	var err error
	DB, err := gorm.Open(sqlite.Open("planner.db"), &gorm.Config{})

	if err != nil{
		return err	
	}

	return DB.AutoMigrate(&Player{})
}
