package handler

import (
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/mrtzee/go-fiber/config"
	"github.com/mrtzee/go-fiber/entity"
	"github.com/mrtzee/go-fiber/model"
)

func BookHandlerGetAll(ctx *fiber.Ctx) error {
	var books []entity.Book
	result := config.DB.Find(&books)
	if result != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(books)
}

func BookHandlerCreate(ctx *fiber.Ctx) error {
	book := new(model.BookCreateRequest)
	if err := ctx.BodyParser(book); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(book)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	// Validation Required Image
	var filenameString string
	filename := ctx.Locals("filename")
	if filename == nil {
		return ctx.Status(422).JSON(fiber.Map{
			"message": "image is required",
		})
	} else {
		filenameString = fmt.Sprintf("%v", filename)
	}

	newBook := entity.Book{
		Title:  book.Title,
		Author: book.Author,
		Cover:  filenameString,
	}

	errCreateBook := config.DB.Create(&newBook).Error
	if errCreateBook != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Failed to store data",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Success",
		"data":    newBook,
	})
}
