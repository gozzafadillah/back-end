package handler_users

import (
	"net/http"
	"ppob/helper/encryption"
	domain_users "ppob/users/domain"
	"ppob/users/handler/request"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UsersHandler struct {
	usecase    domain_users.Service
	validation *validator.Validate
}

func NewUsersHandler(uc domain_users.Service) UsersHandler {
	return UsersHandler{
		usecase:    uc,
		validation: validator.New(),
	}
}

func (uh *UsersHandler) Authorization(ctx echo.Context) error {
	req := request.RequestJSONLogin{}
	ctx.Bind(&req)
	if err := uh.validation.Struct(req); err != nil {
		stringerr := []string{}
		for _, errval := range err.(validator.ValidationErrors) {
			stringerr = append(stringerr, errval.Field()+" is not "+errval.Tag())
		}
		return ctx.JSON(http.StatusBadRequest, stringerr)
	}

	res, err := uh.usecase.Login(req.Email, req.Password)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"rescode": http.StatusBadRequest,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "user success login",
		"rescode": http.StatusOK,
		"data": map[string]interface{}{
			"token": res,
		},
	})
}

func (uh *UsersHandler) Register(ctx echo.Context) error {
	req := request.RequestJSON{}
	ctx.Bind(&req)
	if err := uh.validation.Struct(req); err != nil {
		stringerr := []string{}
		for _, errval := range err.(validator.ValidationErrors) {
			stringerr = append(stringerr, errval.Field()+" is not "+errval.Tag())
		}
		return ctx.JSON(http.StatusBadRequest, stringerr)
	}
	encrypt, err := encryption.HashPassword(req.Password)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
			"rescode": http.StatusInternalServerError,
		})
	}
	req.Password = encrypt
	responseData, err := uh.usecase.Register(request.ToDomain(req))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"rescode": http.StatusBadRequest,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    responseData,
	})

}
