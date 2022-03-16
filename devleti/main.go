package main

import (
	"github.com/devleti/goropbox/devleti/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	dsn := "host=localhost user=leti password=password dbname=goropbox port=5432 TimeZone=Asia/Seoul"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Failed to connect DB")
	}

	err = db.AutoMigrate(
		&model.User{},
		&model.File{},
	)
	if err != nil {
		panic("Failed to migrate")
	}
}
