package controllers

import (
	"api/database"
	. "api/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func GetAllProvinsi(c *fiber.Ctx) error {
	var AllProvinsi []Provinsi

	database.DB.Find(&AllProvinsi)

	return c.JSON(fiber.Map{"provinsi": AllProvinsi})

}

func CreateProvinsi(c *fiber.Ctx) error {
	v := validator.New()

	var provinsi Provinsi
	var res Response

	if err := c.BodyParser(&provinsi); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	namaProvinsi := &provinsi.NamaProvinsi

	provinsiForm := Provinsi{NamaProvinsi: *namaProvinsi}

	err := v.Struct(provinsiForm)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := database.DB.Where("nama_provinsi = ?", provinsiForm.NamaProvinsi).First(&provinsi).Error; err != nil {
		if err := database.DB.Create(&provinsiForm).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
		}
	} else {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Provinsi sudah ada di database",
		})
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = fiber.Map{
		"id":            provinsiForm.Id,
		"nama_provinsi": *namaProvinsi,
	}

	return c.JSON(fiber.Map{
		"status":  res.Status,
		"message": res.Message,
		"data":    res.Data,
	})

}
