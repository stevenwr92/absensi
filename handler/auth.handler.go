package handler

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stevenwr92/absensi/handler/dto"
	"github.com/stevenwr92/absensi/models"
	"github.com/stevenwr92/absensi/utils"
)

var validate = validator.New()

func Register(c *fiber.Ctx) error {
	user := models.User{}

	if err := c.BodyParser(&user); err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Please check your input")
	}

	if err := validate.Struct(user); err != nil {
		var errorMessage string
		for _, validationErr := range err.(validator.ValidationErrors) {
			errorMessage += fmt.Sprintf("%s is %s; ", validationErr.Field(), validationErr.Tag())
		}

		return utils.SendErrorResponse(c, fiber.StatusBadRequest, errorMessage)
	}

	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Failed to hash password")
	}
	user.Password = hash

	if err := utils.DB.Create(&user).Error; err != nil {
		return utils.SendErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}

func Login(c *fiber.Ctx) error {
	user := models.User{}
	login := dto.Login{}

	if err := c.BodyParser(&login); err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Please check your input")
	}

	err := validate.Struct(login)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("%s is %s", err.Field(), err.Tag())
			return utils.SendErrorResponse(c, fiber.StatusBadRequest, errorMessage)
		}
	}

	checkUser := utils.DB.Where("email = ?", login.Email).First(&user)

	if checkUser.Error != nil {
		return utils.SendErrorResponse(c, fiber.StatusUnauthorized, "Invalid credentials")
	}

	if !utils.CheckPasswordHash(login.Password, user.Password) {
		return utils.SendErrorResponse(c, fiber.StatusUnauthorized, "Invalid credentials")
	}

	claims := jwt.MapClaims{
		"Id":       user.ID,
		"UserName": user.UserName,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func Accessible(c *fiber.Ctx) error {
	return c.SendString("Accessible")
}

func Restricted(c *fiber.Ctx) error {

	claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	fmt.Println(claims)
	return c.SendString("Welcome ")
}
