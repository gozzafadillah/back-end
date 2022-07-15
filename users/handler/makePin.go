package handler_users

import (
	"net/http"
	"ppob/app/middlewares"
	"ppob/helper/encryption"
	err_conv "ppob/helper/err"
	"ppob/users/handler/request"
	"ppob/users/handler/response"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// implementation store/save pin data users
func (uh *UsersHandler) MakePin(ctx echo.Context) error {
	req := request.RequestJSONAccount{}
	ctx.Bind(&req)
	if err := uh.Validation.Struct(req); err != nil {
		stringerr := []string{}
		for _, errval := range err.(validator.ValidationErrors) {
			stringerr = append(stringerr, errval.Field()+" is not "+errval.Tag())
		}
		return ctx.JSON(http.StatusBadRequest, stringerr)
	}
	// Encryption
	encrypt, err := encryption.HashPassword(req.Pin)
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}
	req.Pin = encrypt
	// get data from jwt
	dataUser := middlewares.GetUser(ctx)
	req.Phone = dataUser.Phone
	// input from request to usecase layer
	res, err := uh.usecase.InsertAccount(request.ToDomainAccount(req))
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}
	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success create account",
		"rescode": http.StatusCreated,
		"data":    res,
	})
}

// implementation get all data
func (uh *UsersHandler) GetUsers(ctx echo.Context) error {
	sliceResponse := []interface{}{}
	res, err := uh.usecase.GetUsers()
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}
	for _, value := range res {
		sliceResponse = append(sliceResponse, response.FromDomainUsers(value))
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get users",
		"rescode": http.StatusOK,
		"result":  sliceResponse,
	})
}
