package routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Selamat Datang di API Wonderful Indonesia")
	})
	return e
}
