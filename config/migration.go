package config

import (
	"fmt"
	"log"

	"github.com/mrtzee/go-fiber/entity"
)

func RunMigration() {
	err := DB.AutoMigrate(&entity.User{})
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Database migrated")
}
