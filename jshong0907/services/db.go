package services

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorop-box/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	if DB, err = Connection(); err != nil {
		log.Printf("failed to connect database, got error %v", err)
		os.Exit(1)
	} else {
		RunMigrations()
	}
}

func Connection() (*gorm.DB, error) {
	dbDSN := fmt.Sprintf(
		"host=%s user=%s dbname=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
	)
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
