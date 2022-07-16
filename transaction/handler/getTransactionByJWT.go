package handler_transaction

import (
	"fmt"
	"net/http"
	"ppob/app/middlewares"
	err_conv "ppob/helper/err"

	"github.com/labstack/echo/v4"
)

func (th *TransactionHandler) GetHistoryTransaction(ctx echo.Context) error {
	dataMap := []interface{}{}
	// get jwt
	claim := middlewares.GetUser(ctx)
	fmt.Println("phone :", claim.Phone)
	// get transactions by phone
	transactions := th.TransactionUsecase.GetTransactionsByPhone(claim.Phone)

	for i := 0; i < len(transactions); i++ {
		// get Detail product
		detailproduct, err := th.ProductUsecase.GetDetail(transactions[i].Detail_Product_Slug)
		if err != nil {
			return err_conv.Conversion(err, ctx)
		}
		// get product
		product, err := th.ProductUsecase.GetProductTransaction(detailproduct.Product_Slug)
		if err != nil {
			return err_conv.Conversion(err, ctx)
		}

		// get Category product
		category, err := th.ProductUsecase.GetCategory(product.Category_Id)
		if err != nil {
			return err_conv.Conversion(err, ctx)
		}

		payment := th.TransactionUsecase.GetPayment(transactions[i].Payment_Id)
		dataMap = append(dataMap, map[string]interface{}{
			"transaction":    transactions[i].Transaction_Code,
			"category":       category.Name,
			"category_image": category.Image,
			"amount":         transactions[i].Amount,
			"payment_paid":   payment.Paid_at,
			"status":         transactions[i].Status,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"result":  dataMap,
	})
}
