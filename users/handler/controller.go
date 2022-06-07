package handler_users

import (
	"net/http"
	"ppob/app/middlewares"
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
	req := request.RequestJSONUser{}
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
	token, err := uh.usecase.Register(request.ToDomainUser(req))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"rescode": http.StatusBadRequest,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data": map[string]interface{}{
			"token": token,
		},
	})
}

func (uh *UsersHandler) InsertAccount(ctx echo.Context) error {
	req := request.RequestJSONAccount{}
	ctx.Bind(&req)
	if err := uh.validation.Struct(req); err != nil {
		stringerr := []string{}
		for _, errval := range err.(validator.ValidationErrors) {
			stringerr = append(stringerr, errval.Field()+" is not "+errval.Tag())
		}
		return ctx.JSON(http.StatusBadRequest, stringerr)
	}
	// Encryption
	encrypt, err := encryption.HashPassword(req.Pin)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
			"rescode": http.StatusInternalServerError,
		})
	}
	req.Pin = encrypt
	// get data from jwt
	dataUser := middlewares.GetUser(ctx)
	req.Phone = dataUser.Phone

	res, err := uh.usecase.InsertAccount(request.ToDomainAccount(req))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"rescode": http.StatusBadRequest,
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"data":    res,
	})
}

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
