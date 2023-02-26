package database

import (
	"api_wonderful_indonesia/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB
var err error

func Init() *gorm.DB {
	conf := config.GetConfig()

	dsn := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME + "?charset=utf8mb4&parseTime&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("Connection failed", err)
	} else {
		log.Println("Connection established")
	}
	return db
}
