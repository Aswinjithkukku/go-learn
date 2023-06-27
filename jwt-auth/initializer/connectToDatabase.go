package initializer

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("jwt.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to db")
	}

	log.Println("Connected to database successfully")
}
