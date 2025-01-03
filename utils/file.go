package utils

import (
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

const DefaultPathAssetImage = "./public/images/"

func HandleSingleFile(ctx *fiber.Ctx) error {
	file, errFile := ctx.FormFile("cover")
	if errFile != nil {
		log.Println("Error file =", errFile)
	}

	var filename *string
	if file != nil {
		errCheckContentType := CheckContentType(file, "image/jpg", "image/png")
		if errCheckContentType != nil {
			return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"message": errCheckContentType.Error(),
			})
		}
		filename = &file.Filename
		extensionFile := filepath.Ext(*filename)
		newFilename := fmt.Sprintf("gambar-%s", extensionFile)
		errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/images/%s", newFilename))
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
			extensionFile := filepath.Ext(file.Filename)
			filename = fmt.Sprintf("%d-%s%s", i, "gambar", extensionFile)
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

func HandleRemoveFile(filename string, path ...string) error {
	if len(path) > 0 {
		err := os.Remove(path[0] + filename)
		if err != nil {
			log.Println("Failed to remove file")
			return err
		}
		return nil
	} else {
		err := os.Remove(DefaultPathAssetImage + filename)
		if err != nil {
			log.Println("Failed to remove file")
			return err
		}
		return nil
	}

}

func CheckContentType(file *multipart.FileHeader, contentTypes ...string) error {
	if len(contentTypes) > 0 {
		for _, contentType := range contentTypes {
			contentTypeFile := file.Header.Get("Content-Type")
			if contentTypeFile == contentType {
				return nil
			}
		}
		return errors.New("not allow file extension")
	} else {
		return errors.New("not found file")
	}
}
