package helper_xendit

import (
	"encoding/json"
	"errors"
	"fmt"
	"ppob/transaction/handler/request"

	"github.com/labstack/echo/v4"
)

func GetCallBack(ctx echo.Context) (request.Callback_Invoice, error) {
	fmt.Println("otp ", ctx.Request().Header.Get("x-callback-token"))

	decoder := json.NewDecoder(ctx.Request().Body)
	callbackData := request.Callback_Invoice{}

	err := decoder.Decode(&callbackData)
	if err != nil {
		return request.Callback_Invoice{}, errors.New("internal server error")
	}

	defer ctx.Request().Body.Close()

	ctx.Response().Header().Set("Content-Type", "application/json")

	ctx.Response().WriteHeader(200)

	return callbackData, nil
}
