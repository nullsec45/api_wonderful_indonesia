package controllers

import (
	"api/database"
	"api/helpers"
	. "api/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func Register(c *fiber.Ctx) error {
	v := validator.New()
	var user User
	var res Response
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	username := &user.Username
	password, _ := helpers.HashPassword(&user.Password)

	userForm := User{Username: *username, Password: password}
	err := v.Struct(userForm)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := database.DB.Create(&userForm).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = fiber.Map{
		"id":       userForm.Id,
		"username": *username,
	}

	return c.JSON(fiber.Map{
		"status":  res.Status,
		"message": res.Message,
		"data":    res.Data,
	})
}
