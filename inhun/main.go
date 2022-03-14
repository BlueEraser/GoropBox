package main

import (
	"log"

	"github.com/inhun/GoropBox/models"
	"github.com/inhun/GoropBox/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg, _ := utils.LoadConfig("config.json")

	db, err := gorm.Open(postgres.Open(cfg.Url), &gorm.Config{})
	if err != nil {
		log.Println("failed")
	}

	db.AutoMigrate(&models.User{})
	db.Create(&models.User{Email: "inhun321@khu.ac.kr", Password: "test"})

}
