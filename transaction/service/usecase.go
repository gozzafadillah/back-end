package service_transaction

import (
	"errors"
	"ppob/helper/slug"
	domain_transaction "ppob/transaction/domain"

	"github.com/pborman/uuid"
	"github.com/xendit/xendit-go"
)

type TransactionService struct {
	Repository domain_transaction.Repository
}

// GetDetailTransaction implements domain_transaction.Service
func (ts TransactionService) GetDetailTransaction(transaction_code string) (domain_transaction.Detail_Transaction, error) {
	data, err := ts.Repository.GetDetailTransaction(transaction_code)
	if err != nil {
		return domain_transaction.Detail_Transaction{}, errors.New("bad request")
	}
	return data, nil
}

// GetFavoritesByPhone implements domain_transaction.Service
func (ts TransactionService) GetFavoritesByPhone(cat, phone string) domain_transaction.Transaction {
	transactions := ts.Repository.GetFavorite(cat, phone)
	// fmt.Println("transactions ", transactions)
	var a, b int
	var paymenta, paymentb string
	if len(transactions) == 1 {
		data, _ := ts.Repository.GetTransactionByPaymentId(transactions[0].Payment_Id)
		return data
	}
	for i, _ := range transactions {
		paymenta, a = ts.Repository.Count(transactions[i].Category_Slug, transactions[i].Phone, transactions[i].ID_Customer, transactions[i].Detail_Product_Slug)
		// fmt.Println("payment A :", paymenta)
		if a > b {
			b = a
			paymentb = paymenta
			// fmt.Println("payment A :", paymentb)
		}
	}
	data, _ := ts.Repository.GetTransactionByPaymentId(paymentb)

	// fmt.Println("data :", data)

	return data
}

func NewTransactionService(repository domain_transaction.Repository) domain_transaction.Service {
	return TransactionService{
		Repository: repository,
	}
}

// GetTransactionsByPhone implements domain_transaction.Service
func (ts TransactionService) GetTransactionsByPhone(phone string) []domain_transaction.Transaction {
	TransactionSlice := ts.Repository.GetTransactionByPhone(phone)
	return TransactionSlice
}

// AddDetailTransaction implements domain_transaction.Service
func (ts TransactionService) AddDetailTransaction(productCode string, domain domain_transaction.Detail_Transaction) (domain_transaction.Detail_Transaction, error) {
	code := uuid.NewRandom().String()
	domain.Transaction_Code = "transaction-" + code
	domain.Product_Detail_code = productCode

	err := ts.Repository.StoreDetailTransaction(domain)
	if err != nil {
		return domain_transaction.Detail_Transaction{}, errors.New("internal server error")
	}
	return domain, nil
}

// AddTransaction implements domain_transaction.Service
func (ts TransactionService) AddTransaction(data *xendit.Invoice, detail domain_transaction.Detail_Transaction) error {
	transaction := domain_transaction.Transaction{
		Transaction_Code:    detail.Transaction_Code,
		ID_Customer:         detail.ID_Customer,
		Phone:               data.Customer.MobileNumber,
		Amount:              int(data.Amount),
		Category_Slug:       slug.GenerateSlug(data.Items[0].Category),
		Detail_Product_Slug: slug.GenerateSlug(data.Items[0].Name),
		Payment_Id:          data.ID,
		Status:              data.Status,
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

// GetPayment implements domain_transaction.Service
func (ts TransactionService) GetPayment(id string) domain_transaction.Payment {
	data := ts.Repository.GetPayment(id)
	return data
}
