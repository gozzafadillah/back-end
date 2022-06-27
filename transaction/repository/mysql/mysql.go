package mysql_transaction

import (
	domain_transaction "ppob/transaction/domain"

	"gorm.io/gorm"
)

type TransactionRepo struct {
	DB *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) domain_transaction.Repository {
	return TransactionRepo{
		DB: db,
	}
}

// StoreDetailTransaction implements domain_transaction.Repository
func (tr TransactionRepo) StoreDetailTransaction(productCode string, domain domain_transaction.Detail_Transaction) error {
	err := tr.DB.Save(&domain).Error
	return err
}
