package services

import (
	"log"

	"github.com/yangwawa0323/gin-gorm-jwt/models"
	"github.com/yangwawa0323/gin-gorm-jwt/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Here, we are defining an instance of the database.
// This variable will be used across the entire application to communicate with the database.
// var Instance *gorm.DB
var dbError error

type dbService struct {
	DB *gorm.DB
}

type DBService interface {
	Connect(string) (*dbService, error)
	Migrate() error
}

func NewDBService() *dbService {
	dbsvcInstance := &dbService{}
	dbsvc, err := dbsvcInstance.Connect(utils.GetConnectURI())
	if err != nil {
		panic("cannot connect to database, make sure the database service is started")
	}
	return dbsvc
}

func (dbsvc *dbService) Connect(connectionString string) (*dbService, error) {
	dbsvc.DB, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		return nil, dbError
	}
	log.Println("Connected to database")
	return dbsvc, nil
}

func (dbsvc *dbService) Migrate() error {
	if err := dbsvc.DB.AutoMigrate(
		// account models
		&models.User{},
		// page models
		&models.Page{},
		// course models
		&models.Tag{},
		&models.CourseCategory{},
		&models.Course{},
		&models.Class{},
	); err != nil {
		return err
	}
	log.Println("Database Migration completed!")
	return nil
}
