package handler

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/mrtzee/go-fiber/config"
	"github.com/mrtzee/go-fiber/entity"
	"github.com/mrtzee/go-fiber/model"
	"github.com/mrtzee/go-fiber/utils"
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
	filenames := ctx.Locals("filenames")
	if filenames == nil {
		return ctx.Status(422).JSON(fiber.Map{
			"message": "image is required",
		})
	} else {
		filenamesData := filenames.([]string)
		for _, filename := range filenamesData {
			newPhoto := entity.Photo{
				Image:      filename,
				CategoryId: photo.CategoryId,
			}

			errCreatePhoto := config.DB.Create(&newPhoto).Error
			if errCreatePhoto != nil {
				log.Println("Something wrong in ur file")
			}
		}
	}

	return ctx.JSON(fiber.Map{
		"message": "Success",
	})
}

func PhotoHandlerDelete(ctx *fiber.Ctx) error {
	var photo entity.Photo
	photoId := ctx.Params("id")
	err := config.DB.First(&photo, "id = ?", photoId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "image not found",
		})
	}

	// Delete file from public
	errDeleteFile := utils.HandleRemoveFile(photo.Image)
	if errDeleteFile != nil {
		log.Println("Something wrong when delete this file")
	}

	errDelete := config.DB.Delete(&photo).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "image has deleted",
	})
}
