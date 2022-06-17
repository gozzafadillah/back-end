package err_conv

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func Conversion(err error, ctx echo.Context) error {
	var errNew error
	if strings.Contains(err.Error(), BadRequest) {
		errNew = ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": BadRequest,
			"rescode": http.StatusBadRequest,
		})
	}
	if strings.Contains(err.Error(), InternalServerErr) {
		errNew = ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": InternalServerErr,
			"rescode": http.StatusInternalServerError,
		})
	}
	if strings.Contains(err.Error(), MissMatchEmail) {
		errNew = ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": MissMatchEmail,
			"rescode": http.StatusBadRequest,
		})
	}
	if strings.Contains(err.Error(), UnauthorizedUser) {
		errNew = ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": UnauthorizedUser,
			"rescode": http.StatusUnauthorized,
		})
	}
	return errNew
}
