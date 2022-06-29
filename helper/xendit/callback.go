package helper_xendit

import (
	"encoding/json"
	"errors"
	"fmt"
	domain_transaction "ppob/transaction/domain"

	"github.com/labstack/echo/v4"
)

func GetCallback(ctx echo.Context) (interface{}, error) {
	fmt.Println("otp ", ctx.Request().Header.Get("x-callback-token"))

	decoder := json.NewDecoder(ctx.Request().Body)
	callbackData := domain_transaction.Callback_Invoice{}

	err := decoder.Decode(&callbackData)
	if err != nil {
		return "empty", errors.New("internal status error")
	}

	defer ctx.Request().Body.Close()

	callback, _ := json.Marshal(callbackData)

	ctx.Response().Header().Set("Content-Type", "application/json")

	ctx.Response().WriteHeader(200)
	return callback, nil
}

// BjVVRO8eKgceve38jmqm6twtK9YLjtAfk7CbJLxfiToTilHX
