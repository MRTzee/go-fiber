package handler

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/mrtzee/go-fiber/config"
	"github.com/mrtzee/go-fiber/entity"
	"github.com/mrtzee/go-fiber/model"
)

func CategoryHandlerGetAll(ctx *fiber.Ctx) error {
	var categories []entity.Category
	result := config.DB.Find(&categories)
	if result != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(categories)
}

func CategoryHandlerCreate(ctx *fiber.Ctx) error {
	category := new(model.CategoryCreateRequest)
	if err := ctx.BodyParser(category); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(category)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	newCategory := entity.Category{
		Name: category.Name,
	}

	errCreateCategory := config.DB.Create(&newCategory).Error
	if errCreateCategory != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Failed to store data",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Success",
		"data":    newCategory,
	})
}
