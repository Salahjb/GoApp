package database

import (
	"gin-backend/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	err = database.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate users table:", err)
	}

	DB = database
	log.Println("Connected to database successfully")
}
