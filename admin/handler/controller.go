package handler_admin

import (
	"net/http"
	err_conv "ppob/helper/err"
	domain_products "ppob/products/domain"
	domain_transaction "ppob/transaction/domain"
	domain_users "ppob/users/domain"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	TransactionUsecase domain_transaction.Service
	ProductUsecase     domain_products.Service
	UsersUsecase       domain_users.Service
}

func NewAdminHandler(ts domain_transaction.Service, ps domain_products.Service, us domain_users.Service) AdminHandler {
	return AdminHandler{
		TransactionUsecase: ts,
		ProductUsecase:     ps,
		UsersUsecase:       us,
	}
}

func (ah *AdminHandler) GetAllTransaction(ctx echo.Context) error {
	transactions := ah.TransactionUsecase.GetTransactionAll()
	if len(transactions) == 0 {
		return ctx.JSON(200, map[string]interface{}{
			"message": "success",
			"rescode": 200,
			"result":  "empty",
		})
	}
	sliceData := map[int]interface{}{}
	for i := 0; i <= len(transactions)-1; i++ {
		payment := ah.TransactionUsecase.GetPayment(transactions[i].Payment_Id)
		userSession, err := ah.UsersUsecase.GetUserPhone(transactions[i].Phone)
		if err != nil {
			return err_conv.Conversion(err, ctx)
		}

		data := map[string]interface{}{
			"user":   userSession.Name,
			"paid":   payment.Paid_at,
			"amount": transactions[i].Amount,
			"status": transactions[i].Status,
		}
		sliceData[i] = data
	}
	return ctx.JSON(200, map[string]interface{}{
		"message": "success",
		"rescode": 200,
		"result":  sliceData,
	})
}

func (ah *AdminHandler) CountAllItems(ctx echo.Context) error {
	countUsers := ah.UsersUsecase.CountUsersCustomer()
	countProducts := ah.ProductUsecase.CountProducts()
	countTransactions := ah.TransactionUsecase.CountTransaction()

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"result": map[string]interface{}{
			"users":        countUsers,
			"products":     countProducts,
			"transactions": countTransactions,
		},
	})
}
