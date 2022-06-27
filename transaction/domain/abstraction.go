package domain_transaction

type Service interface {
	AddDetailTransaction(productCode string, domain Detail_Transaction) (Detail_Transaction, error)
}

type Repository interface {
	StoreDetailTransaction(productCode string, domain Detail_Transaction) error
}
