package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"phobyjun/model"
)

const dsn = "host=localhost user=phobyjun password=password dbname=goropbox port=5432 TimeZone=Asia/Seoul"

var Session *gorm.DB
var err error

func init() {
	Session, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("DB Connection Failed")
	}

	err = Session.AutoMigrate(
		&model.User{},
		&model.File{},
	)
	if err != nil {
		log.Panic("DB Migration Failed")
	}
}
