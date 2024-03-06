package utils

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func RespondOK(c echo.Context, message string, data any) error {
	return c.JSON(http.StatusOK, EchoResponse(message, data))
}

func BadRequest(c echo.Context) error {
	return c.JSON(http.StatusBadRequest, EchoResponse("bad request", nil))
}

func InternalServerErr(c echo.Context) error {
	return c.JSON(http.StatusInternalServerError, EchoResponse("internal server error", nil))
}
