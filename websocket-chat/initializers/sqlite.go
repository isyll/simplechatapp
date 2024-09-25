package initializers

import (
	"log"
	"websocket-chat/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const DB_PATH string = "./database/data.db"
var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open(sqlite.Open(DB_PATH), &gorm.Config{})

	if err != nil {
	  	panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Group{})
	db.AutoMigrate(&models.Message{})

	DB = db

	log.Println("Database successfully initialized.")
}
