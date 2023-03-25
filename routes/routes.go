package routes

import (
	. "api/controllers"
	"github.com/gofiber/fiber/v2"
)

func Init() {
	app := fiber.New()

	api := app.Group("/api")
	provinsi := api.Group("/provinsi")

	api.Post("/register", Register)
	api.Post("/login", Login)

	provinsi.Get("/", GetAllProvinsi)
	provinsi.Post("/", CreateProvinsi)
	app.Listen(":8000")
}
