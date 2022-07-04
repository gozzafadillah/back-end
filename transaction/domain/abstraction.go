package domain_transaction

import "github.com/xendit/xendit-go"

type Service interface {
	// detail transaction
	AddDetailTransaction(productCode string, domain Detail_Transaction) (Detail_Transaction, error)

	// transaction
	GetTransactionsByPhone(phone string) []Transaction
	AddTransaction(data *xendit.Invoice, detail Detail_Transaction) error
	EditTransaction(data Callback_Invoice) error

	// payment
	GetPayment(id string) Payment
}

type Repository interface {
	// detail transaction / checkout
	StoreDetailTransaction(domain Detail_Transaction) error
	// transaction
	StoreTransaction(domain Transaction) error
	GetTransactionByPhone(phone string) (transaction []Transaction)
	GetTransactionByPaymentId(id string) (Transaction, error)
	UpdateTransaction(domain Transaction) error
	// payment
	StorePayment(domain Payment) error
	GetPayment(id string) Payment
}
