package handler_users

import (
	"net/http"

	err_conv "ppob/helper/err"
	domain_users "ppob/users/domain"
	"ppob/users/handler/request"
	"ppob/users/handler/response"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UsersHandler struct {
	usecase    domain_users.Service
	Validation *validator.Validate
}

func NewUsersHandler(uc domain_users.Service) UsersHandler {
	return UsersHandler{
		usecase:    uc,
		Validation: validator.New(),
	}
}

// Implementation get user by phone for admin (web)
func (uh *UsersHandler) GetUserForAdmin(ctx echo.Context) error {
	phone := ctx.Param("phone")
	// get user by phone
	user, err := uh.usecase.GetUserPhone(phone)
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}
	// get user account (get saldo)
	account := uh.usecase.GetUserAccount(phone)

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get users",
		"rescode": http.StatusOK,
		"result": map[string]interface{}{
			"user":    response.FromDomainUsers(user),
			"account": response.FromDomainAccount(account),
		},
	})
}

func (uh *UsersHandler) VerifUser(ctx echo.Context) error {
	req := request.RequestJSONVerif{}
	ctx.Bind(&req)
	if err := uh.Validation.Struct(req); err != nil {
		stringerr := []string{}
		for _, errval := range err.(validator.ValidationErrors) {
			stringerr = append(stringerr, errval.Field()+" is not "+errval.Tag())
		}
		return ctx.JSON(http.StatusBadRequest, stringerr)
	}

	// get response verification otp
	err := uh.usecase.Verif(req.Code)
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success validate user",
		"rescode": http.StatusOK,
	})
}

// implementation for filter user role by jwt
func (uh *UsersHandler) UserRole(phone string) (string, bool) {
	var role string
	var status bool
	user, err := uh.usecase.GetUserPhone(phone)
	if err == nil {
		role = user.Role
		status = user.Status
	}
	return role, status
}
