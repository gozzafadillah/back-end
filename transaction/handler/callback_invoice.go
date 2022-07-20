package handler_transaction

import (
	"fmt"
	err_conv "ppob/helper/err"
	helper_xendit "ppob/helper/xendit"
	"ppob/transaction/handler/request"

	"github.com/labstack/echo/v4"
)

func (th *TransactionHandler) Callback_Invoice(ctx echo.Context) error {
	data, dataByte, err := helper_xendit.GetCallBack(ctx)
	fmt.Println(dataByte)
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}

	err = th.TransactionUsecase.EditTransaction(request.ToDomainCallBack(data))
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}
	_, err = fmt.Fprintf(ctx.Response().Writer, "%s", "ok")
	return err
}
