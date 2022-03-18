package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"phobyjun/config"
	"phobyjun/model"
)

var Session *gorm.DB
var err error

func init() {
	if connect() != nil {
		log.Panic("DB Connection Failed")
	}
	if migrate() != nil {
		log.Panic("DB Migration Failed")
	}
}

func connect() error {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s TimeZone=Asia/Seoul",
		config.Cfg.Database.HOST,
		config.Cfg.Database.USER,
		config.Cfg.Database.PASSWORD,
		config.Cfg.Database.DBNAME,
		config.Cfg.Database.PORT,
	)
	Session, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return err
}

func migrate() error {
	return Session.AutoMigrate(
		&model.User{},
		&model.File{},
	)
}
