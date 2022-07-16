package handler_transaction

import (
	"fmt"
	"ppob/app/middlewares"
	err_conv "ppob/helper/err"

	"github.com/labstack/echo/v4"
)

func (th *TransactionHandler) FavoriteUser(ctx echo.Context) error {
	claims := middlewares.GetUser(ctx)
	data := map[string]interface{}{}
	categories, err := th.ProductUsecase.GetCategories()
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}
	fmt.Println("cat", categories)
	for i, _ := range categories {
		// fmt.Println("cat =====> ", categories[i].Category_Slug)
		transaction := th.TransactionUsecase.GetFavoritesByPhone(categories[i].Category_Slug, claims.Phone)

		// detail transaction
		detailTransaction, _ := th.TransactionUsecase.GetDetailTransaction(transaction.Transaction_Code)

		// detail product
		detailproduct, _ := th.ProductUsecase.GetDetail(transaction.Detail_Product_Slug)

		// get product
		product, _ := th.ProductUsecase.GetProductTransaction(detailproduct.Product_Slug)

		// get Category product

		data[categories[i].Category_Slug] = map[string]interface{}{
			"customer_name": detailTransaction.Customer_Name,
			"payment_id":    transaction.Payment_Id,
			"id_customer":   transaction.ID_Customer,
			"product_image": product.Image,
		}
	}

	return ctx.JSON(200, map[string]interface{}{
		"message": "success",
		"rescode": 200,
		"result":  data,
	})
}
