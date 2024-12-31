package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrtzee/go-fiber/config"
	"github.com/mrtzee/go-fiber/handler"
	"github.com/mrtzee/go-fiber/middleware"
	"github.com/mrtzee/go-fiber/utils"
)

func RoutesIndex(r *fiber.App) {
	r.Static("/public", config.ProjectRootPath+"/public")
	r.Static("/public", "./public")

	r.Post("/login", handler.LoginHandler)

	r.Get("/users", middleware.Auth, handler.UserHandlerGetAll)
	r.Get("/users/:id", handler.UserHandlerGetById)
	r.Post("/users", handler.UserHandlerCreate)
	r.Put("/users/:id", handler.UserHandlerUpdate)
	r.Put("/users/:id/update-email", handler.UserHandlerUpdateEmail)
	r.Delete("/users/:id", handler.UserHandlerDelete)

	r.Get("/books", handler.BookHandlerGetAll)
	r.Post("/books", utils.HandleSingleFile, handler.BookHandlerCreate)

	r.Get("/category", handler.CategoryHandlerGetAll)
	r.Post("/category", handler.CategoryHandlerCreate)

	r.Get("/gallery", handler.PhotoHandlerGetAll)
	r.Post("/gallery", utils.HandleMultipleFile, handler.PhotoHandlerCreate)
	r.Delete("/gallery/:id", handler.PhotoHandlerDelete)
}
