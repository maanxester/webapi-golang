package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB

func RunDB() {
	str := "host=localhost port=5432 user=anakin dbname=users sslmode=disable password=anakin"

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})

	if err != nil {
		log.Fatal("Error database:", err)
	}

	db = database

	config, _ := db.DB()


	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxIdleTime(time.Hour)
}


func getDatabase() *gorm.DB {
	return db
}