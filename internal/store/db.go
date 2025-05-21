package store

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"os"
	"fmt"
)

var DB *gorm.DB

func InitDatabase() error {
	var err error
	connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DBNAME"),
		os.Getenv("DB_PORT"),
		)

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})

	if err != nil{
		return err	
	}

	DB = db
	return DB.AutoMigrate(&Team{}, &Player{})
}
