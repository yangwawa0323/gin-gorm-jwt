package database

import (
	"log"

	"github.com/yangwawa0323/gin-gorm-jwt/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Here, we are defining an instance of the database.
// This variable will be used across the entire application to communicate with the database.
var Instance *gorm.DB
var dbError error

func Connect(connectionString string) {
	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to db")
	}
	log.Println("Connected to database")
}

func Migrate() {
	Instance.AutoMigrate(&models.User{})
	log.Println("Database Migration completed!")
}
