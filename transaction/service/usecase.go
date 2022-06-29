package service

import (
	"errors"
	domain_transaction "ppob/transaction/domain"

	"github.com/pborman/uuid"
	"github.com/xendit/xendit-go"
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

// AddTransaction implements domain_transaction.Service
func (ts TransactionService) AddTransaction(data *xendit.Invoice, detail domain_transaction.Detail_Transaction) error {
	transaction := domain_transaction.Transaction{
		Transaction_Code: detail.Transaction_Code,
		ID_Customer:      detail.ID_Customer,
		Phone:            data.Customer.MobileNumber,
		Amount:           int(data.Amount),
		Payment_Id:       data.ID,
		Status:           data.Status,
	}
	err := ts.Repository.StoreTransaction(transaction)
	if err != nil {
		return errors.New("transaction failed")
	}
	return nil
}

// EditTransaction implements domain_transaction.Service
func (ts TransactionService) EditTransaction(data domain_transaction.Callback_Invoice) error {
	transaction, err := ts.Repository.GetTransactionByPaymentId(data.ID)
	if err != nil {
		return errors.New("transaction not found")
	}
	transaction.Status = data.Status
	err = ts.Repository.UpdateTransaction(transaction)
	if err != nil {
		return errors.New("update failed")
	}
	pay := domain_transaction.Payment{
		Payment_Id: data.ID,
		Method:     data.PaymentMethod,
		Channel:    data.PaymentChannel,
		Paid_at:    data.PaidAt,
	}
	err = ts.Repository.StorePayment(pay)
	if err != nil {
		return errors.New("internal server error")
	}
	return nil
}
