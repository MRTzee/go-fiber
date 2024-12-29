package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrtzee/go-fiber/config"
	"github.com/mrtzee/go-fiber/handler"
	"github.com/mrtzee/go-fiber/middleware"
)

func RoutesIndex(r *fiber.App) {
	r.Static("/public", config.ProjectRootPath+"/public/asset")
	r.Get("/users", middleware.Auth, handler.UserHandlerGetAll)
	r.Get("/users/:id", handler.UserHandlerGetById)
	r.Post("/users", handler.UserHandlerCreate)
	r.Put("/users/:id", handler.UserHandlerUpdate)
	r.Put("/users/:id/update-email", handler.UserHandlerUpdateEmail)
	r.Delete("/users/:id", handler.UserHandlerDelete)
}
