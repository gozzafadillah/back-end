package domain_transaction

import "github.com/xendit/xendit-go"

type Service interface {
	AddDetailTransaction(productCode string, domain Detail_Transaction) (Detail_Transaction, error)
	AddTransaction(data *xendit.Invoice, detail Detail_Transaction) error
}

type Repository interface {
	StoreDetailTransaction(productCode string, domain Detail_Transaction) error
	StoreTransaction(domain Transaction) error
}
