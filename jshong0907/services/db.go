package services

import (
	"fmt"
	"log"
	"os"

	"gorop-box/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	if db, err = connection(); err != nil {
		log.Printf("failed to connect database, got error %v", err)
		os.Exit(1)
	} else {
		runMigrations()
	}
}

func connection() (*gorm.DB, error) {
	dbDSN := fmt.Sprintf(
		"host=%s user=%s dbname=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
	)
	database, err := gorm.Open(postgres.Open(dbDSN), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return database, err
}

func runMigrations() {
	allModels := []interface{}{&models.User{}, &models.File{}}

	for _, model := range allModels {
		db.AutoMigrate(model)
	}
}
