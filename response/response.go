package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rammm2005/shared-go/errors"
)

type StandardResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

// standard JSON response for Echo
func JSON(c echo.Context, data interface{}, err error) error {
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			return c.JSON(appErr.Code, StandardResponse{
				Success: false,
				Error: echo.Map{
					"message": appErr.Message,
					"details": appErr.Details,
				},
			})
		}
		return c.JSON(http.StatusInternalServerError, StandardResponse{
			Success: false,
			Error:   echo.Map{"message": err.Error()},
		})
	}
	return c.JSON(http.StatusOK, StandardResponse{
		Success: true,
		Data:    data,
	})
}
