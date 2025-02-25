package usecase

import "github.com/labstack/echo/v4"

type Status struct {
	Message string `json:"message"`
}

func ErrorResponse(c echo.Context, statuscode int, message string) error {
	return c.JSON(statuscode, Status{
		Message: message,
	})
}
