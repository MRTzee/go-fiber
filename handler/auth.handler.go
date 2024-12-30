package handler

import (
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mrtzee/go-fiber/config"
	"github.com/mrtzee/go-fiber/entity"
	"github.com/mrtzee/go-fiber/model"
	"github.com/mrtzee/go-fiber/utils"
)

func LoginHandler(ctx *fiber.Ctx) error {
	loginRequest := new(model.Auth)
	if err := ctx.BodyParser(loginRequest); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	var user entity.User
	err := config.DB.First(&user, "email = ?", loginRequest.Email).Error
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Wrong Credential",
		})
	}

	isValid := utils.CheckPasswordHas(loginRequest.Password, user.Password)
	if !isValid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Wrong Credential",
		})
	}

	// JWT
	claims := jwt.MapClaims{}
	claims["names"] = user.Name
	claims["email"] = user.Email
	claims["address"] = user.Address
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()
	if user.Email == "popon8345@mail.com" {
		claims["role"] = "admin"
	} else {
		claims["role"] = "user"
	}

	token, errGenerateToken := utils.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Wrong Credential",
		})
	}

	return ctx.JSON(fiber.Map{
		"token": token,
	})
}
