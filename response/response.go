package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rammm2005/shared-go/errors"
)

func JSON(c echo.Context, data interface{}, err error) error {
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			return c.JSON(appErr.Code, echo.Map{
				"success": false,
				"error":   appErr.Message,
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
		"data":    data,
	})
}
