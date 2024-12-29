package main

import (
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
	app.Listen(":8080")
}
