package service

import (
	"rssReader/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DatabaseServiceInstance = DatabaseService{}

type DatabaseService struct {
	Connection *gorm.DB
}

func (databaseService *DatabaseService) Connect() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if(err != nil) {
		panic(err)
	}

	databaseService.Connection = db
}

func (datebaseService *DatabaseService) Migrate() {
	datebaseService.Connection.AutoMigrate(&models.Provider{})
}
