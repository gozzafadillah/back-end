package handler_transaction

import (
	"net/http"
	err_conv "ppob/helper/err"

	"github.com/labstack/echo/v4"
)

func (th *TransactionHandler) GetTransactionByPaymentSuccess(ctx echo.Context) error {
	paymentID := ctx.Param("payment_id")
	res, err := th.TransactionUsecase.GetTransactionByPaymentId(paymentID)
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"result":  res.Amount,
	})
}
