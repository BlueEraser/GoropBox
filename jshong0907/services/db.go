package services

import (
	"gorop-box/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	if DB, err = Connection(); err != nil {
		log.Printf("failed to connect database, got error %v", err)
		os.Exit(1)
	} else {
		RunMigrations()
	}
}

func Connection() (*gorm.DB, error) {
	dbDSN := "host=127.0.0.1 user=hong dbname=GoropBox"
	db, err := gorm.Open(postgres.Open(dbDSN), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db, err
}

func RunMigrations() {
	allModels := []interface{}{&models.User{}, &models.File{}}

	for _, model := range allModels {
		DB.AutoMigrate(model)
	}
}
