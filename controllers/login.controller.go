package controllers

import (
	. "api/database"
	"api/helpers"
	. "api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func Login(c *fiber.Ctx) error {
	var userInput User
	var user User
	var res Response

	if err := c.BodyParser(&userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	username := &userInput.Username
	password := &userInput.Password

	// ambil data user berdasarkan username
	if err := DB.Where("username = ?", username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			res.Message = "Gagal login!. Silahkan masukkan username atau password lagi."
			return c.JSON(fiber.Map{
				"message": res.Message,
			})
		default:
			res.Message = err.Error()
			return c.JSON(fiber.Map{
				"message": res.Message,
			})
		}
	}

	match, _ := helpers.CheckPasswordHash(*password, user.Password)
	if !match {
		res.Message = "Gagal login!. Silahkan masukkan username atau password lagi."
		return c.JSON(fiber.Map{
			"message": res.Message,
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = username
	claims["level"] = "application"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenJWT, err := token.SignedString([]byte("secret"))

	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = err.Error()
		return c.JSON(fiber.Map{
			"status":  res.Status,
			"message": res.Message,
		})
	}
	res.Status = http.StatusOK
	res.Message = "Berhasil login!"
	return c.JSON(fiber.Map{
		"status":  res.Status,
		"message": res.Message,
		"token":   tokenJWT,
	})
}
