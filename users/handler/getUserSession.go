package handler_users

import (
	"net/http"
	"ppob/app/middlewares"
	err_conv "ppob/helper/err"
	regexPhone "ppob/helper/phone"
	"ppob/users/handler/response"

	"github.com/labstack/echo/v4"
)

// Implementation get user session
func (uh *UsersHandler) GetUserSession(ctx echo.Context) error {
	jwtClaims := middlewares.GetUser(ctx)
	phone := jwtClaims.Phone
	// get user phone
	user, err := uh.usecase.GetUserPhone(phone)
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}
	// get user account
	account := uh.usecase.GetUserAccount(phone)

	// generate to old phone
	oldPhone := regexPhone.GenerateToOld(account.Phone)
	user.Phone = oldPhone

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get customer",
		"rescode": http.StatusOK,
		"result": map[string]interface{}{
			"user":    response.FromDomainUsers(user),
			"account": response.FromDomainAccount(account),
		},
	})
}
