package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/zob456/weather_api/internal/handlers/weather_handlers"
)

func Routes(e *echo.Echo) {
	api := e.Group("api")
	// Health check is not really needed for a coding exercise but I like having it out of habit
	api.GET("/health", func(c echo.Context) error {
		return c.JSON(200, "success")
	})

	api.POST("", weather_handlers.GetWeather)

}
