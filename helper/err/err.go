package err_conv

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func Conversion(err error, ctx echo.Context) error {
	var errNew error
	if strings.Contains(err.Error(), BadRequest) ||
		strings.Contains(err.Error(), UpdateFailed) ||
		strings.Contains(err.Error(), DeleteFailed) ||
		strings.Contains(err.Error(), Err) {
		errNew = ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusBadRequest,
		})
	}
	if strings.Contains(err.Error(), InternalServerErr) {
		errNew = ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusInternalServerError,
		})
	}
	if strings.Contains(err.Error(), MissMatchEmail) {
		errNew = ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusBadRequest,
		})
	}
	if strings.Contains(err.Error(), UnauthorizedUser) {
		errNew = ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusUnauthorized,
		})
	}
	return errNew
}
