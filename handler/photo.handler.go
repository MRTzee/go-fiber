package handler

import (
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/mrtzee/go-fiber/model"
)

func PhotoHandlerCreate(ctx *fiber.Ctx) error {
	photo := new(model.PhotoCreateRequest)
	if err := ctx.BodyParser(photo); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(photo)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	// Validation Required Image
	var filenameString string
	filenames := ctx.Locals("filenames")
	if filenames == nil {
		return ctx.Status(422).JSON(fiber.Map{
			"message": "image is required",
		})
	} else {
		filenameString = fmt.Sprintf("%v", filenames)
	}

	log.Println(filenameString)
	// newPhoto := entity.Photo{
	// 	Image:      filenameString,
	// 	CategoryId: 1,
	// }

	// errCreatePhoto := config.DB.Create(&newPhoto).Error
	// if errCreatePhoto != nil {
	// 	return ctx.Status(500).JSON(fiber.Map{
	// 		"message": "Failed to store data",
	// 	})
	// }

	return ctx.JSON(fiber.Map{
		"message": "Success",
	})
}
