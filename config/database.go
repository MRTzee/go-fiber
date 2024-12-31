package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		ENV.DB_HOST,
		ENV.DB_USER,
		ENV.DB_PASSWORD,
		ENV.DB_DATABASE,
		ENV.DB_PORT,
	)

	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	DB = db
	log.Println("Database connected")
}
