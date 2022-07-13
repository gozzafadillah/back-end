package domain_transaction

import "github.com/xendit/xendit-go"

type Service interface {
	// detail transaction
	AddDetailTransaction(productCode string, domain Detail_Transaction) (Detail_Transaction, error)
	GetDetailTransaction(transaction_code string) (Detail_Transaction, error)
	GetTransactionAll() []Transaction

	// transaction
	GetTransactionsByPhone(phone string) []Transaction
	GetFavoritesByPhone(cat, phone string) Transaction
	AddTransaction(data *xendit.Invoice, detail Detail_Transaction) error
	EditTransaction(data Callback_Invoice) error

	// payment
	GetPayment(id string) Payment

	// Admin-Dashboard
	CountTransaction() int
}

type Repository interface {
	// detail transaction / checkout
	StoreDetailTransaction(domain Detail_Transaction) error
	GetDetailTransaction(transaction_code string) (Detail_Transaction, error)

	// transaction
	StoreTransaction(domain Transaction) error
	GetFavorite(cat, phone string) []Transaction
	Count(cat, phone, id_customer, detail_product string) (string, int)
	GetTransactions() []Transaction
	GetTransactionByPhone(phone string) (transaction []Transaction)
	GetTransactionByPaymentId(payment_id string) (Transaction, error)
	UpdateTransaction(domain Transaction) error

	// payment
	StorePayment(domain Payment) error
	GetPayment(id string) Payment

	//
	Counts() int
}
