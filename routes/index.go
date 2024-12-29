package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrtzee/go-fiber/handler"
)

func RoutesIndex(r *fiber.App) {
	r.Get("/users", handler.UserHandlerGetAll)
	r.Get("/users/:id", handler.UserHandlerGetById)
	r.Post("/users", handler.UserHandlerCreate)
}
