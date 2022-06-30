package helper_xendit

import (
	"encoding/json"
	"errors"
	"os"
	"ppob/transaction/handler/request"

	"github.com/labstack/echo/v4"
)

func GetCallBack(ctx echo.Context) (request.Callback_Invoice, []byte, error) {
	tokenBayeueCallback := os.Getenv("OTP_Xendit_Callback")
	xTokenCallback := ctx.Request().Header.Get("x-callback-token")
	if xTokenCallback != tokenBayeueCallback {
		return request.Callback_Invoice{}, []byte{}, errors.New("unauthorized")
	}

	decoder := json.NewDecoder(ctx.Request().Body)
	callbackData := request.Callback_Invoice{}

	err := decoder.Decode(&callbackData)
	if err != nil {
		return request.Callback_Invoice{}, []byte{}, errors.New("internal server error")
	}

	callback, _ := json.Marshal(callbackData)

	defer ctx.Request().Body.Close()

	ctx.Response().Header().Set("Content-Type", "application/json")

	ctx.Response().WriteHeader(200)

	return callbackData, callback, nil
}
