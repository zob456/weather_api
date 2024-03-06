package handlers

import "github.com/labstack/echo/v4"

func Routes(e *echo.Echo) {
	api := e.Group("api")
	api.GET("/health", func(c echo.Context) error {
		return c.JSON(200, "success")
	})
}
