package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	db *gorm.DB
)

func Connect() {
	// Updated DSN format for GORM v2
	dsn := "snorpiii:0000@tcp(localhost:3306)/book?charset=utf8&parseTime=True&loc=Local"

	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {

		log.Panic("Failed to connect to database:", err)

	}

	db = d
}

func GetDB() *gorm.DB {
	
	return db
}
