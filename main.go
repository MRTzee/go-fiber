package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/mrtzee/go-fiber/config"
	"github.com/mrtzee/go-fiber/routes"
)

func main() {
	// Database
	config.LoadConfig()
	config.ConnectDB()
	config.RunMigration()

	// Routes
	app := fiber.New()
	routes.RoutesIndex(app)
	errListen := app.Listen(":8080")
	if errListen != nil {
		log.Println("Fail to listen server")
		os.Exit(1)
	}
}
