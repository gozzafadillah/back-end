package domain_transaction

import "github.com/xendit/xendit-go"

type Service interface {
	AddDetailTransaction(productCode string, domain Detail_Transaction) (Detail_Transaction, error)
	AddTransaction(data *xendit.Invoice, detail Detail_Transaction) error
	EditTransaction(data Callback_Invoice) error
}

type Repository interface {
	// detail transaction / checkout
	StoreDetailTransaction(productCode string, domain Detail_Transaction) error
	// transaction
	StoreTransaction(domain Transaction) error
	GetTransactionByPaymentId(id string) (Transaction, error)
	UpdateTransaction(domain Transaction) error
	// payment
	StorePayment(domain Payment) error
}
