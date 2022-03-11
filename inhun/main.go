package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/inhun/GoropBox/config"
	"github.com/inhun/GoropBox/endpoints"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func main() {
	cfg, _ := config.LoadConfig("config.json")

	DBUrl := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d TimeZone=%s", cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.DBName, cfg.DB.Port, cfg.DB.TimeZone)

	db, err := gorm.Open(postgres.Open(DBUrl), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, err := db.DB()
	defer sqlDB.Close()

	ep := endpoints.Endpoints{DB: db}

	// db.AutoMigrate(&models.User{})
	// db.Create(&models.User{Email: "inhun321@khu.ac.kr", Password: "test"})

	router := httprouter.New()

	router.GET("/api/user", ep.GetUserList)

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+"8000", handler))

}
