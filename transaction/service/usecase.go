package service

import (
	"errors"
	domain_transaction "ppob/transaction/domain"

	"github.com/pborman/uuid"
)

type TransactionService struct {
	Repository domain_transaction.Repository
}

func NewTransactionService(repository domain_transaction.Repository) domain_transaction.Service {
	return TransactionService{
		Repository: repository,
	}
}

// AddDetailTransaction implements domain_transaction.Service
func (ts TransactionService) AddDetailTransaction(productCode string, domain domain_transaction.Detail_Transaction) (domain_transaction.Detail_Transaction, error) {
	code := uuid.NewRandom().String()
	domain.Transaction_Code = "transaction-" + code
	domain.Product_Detail_code = productCode

	err := ts.Repository.StoreDetailTransaction(productCode, domain)
	if err != nil {
		return domain_transaction.Detail_Transaction{}, errors.New("internal server error")
	}
	return domain, nil
}
