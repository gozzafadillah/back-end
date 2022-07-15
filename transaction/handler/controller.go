package handler_transaction

import (
	domain_products "ppob/products/domain"
	domain_transaction "ppob/transaction/domain"
	domain_users "ppob/users/domain"

	"github.com/go-playground/validator/v10"
)

type TransactionHandler struct {
	TransactionUsecase domain_transaction.Service
	ProductUsecase     domain_products.Service
	UserUsecase        domain_users.Service
	Validation         *validator.Validate
}

func NewTransactionHandler(tc domain_transaction.Service, pc domain_products.Service, uc domain_users.Service) TransactionHandler {
	return TransactionHandler{
		TransactionUsecase: tc,
		ProductUsecase:     pc,
		UserUsecase:        uc,
		Validation:         validator.New(),
	}
}
