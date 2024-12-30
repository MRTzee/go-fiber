package utils

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func HandleSingleFile(ctx *fiber.Ctx) error {
	file, errFile := ctx.FormFile("cover")
	if errFile != nil {
		log.Println("Error file =", errFile)
	}

	var filename *string
	if file != nil {
		filename = &file.Filename
		errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/images/%s", *filename))
		if errSaveFile != nil {
			log.Println("Fail to store file into public/images")
		}
	} else {
		log.Println("Nothing file upload")
	}

	if filename != nil {
		ctx.Locals("filename", *filename)
	} else {
		ctx.Locals("filename", nil)

	}

	return ctx.Next()
}

func HandleMultipleFile(ctx *fiber.Ctx) error {
	form, errForm := ctx.MultipartForm()
	if errForm != nil {
		log.Println("Error read multipart form request, Error =", errForm)
	}

	files := form.File["photos"]
	var filenames []string
	for i, file := range files {
		var filename string
		if file != nil {
			filename = fmt.Sprintf("%d-%s", i, file.Filename)
			errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/images/%s", filename))
			if errSaveFile != nil {
				log.Println("Fail to store file into public/images")
			}
		} else {
			log.Println("Nothing file upload")
		}

		if filename != "" {
			filenames = append(filenames, filename)
		}
	}

	ctx.Locals("filenames", filenames)
	return ctx.Next()
}
