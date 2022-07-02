package mysql_transaction

import (
	"errors"
	"fmt"
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
	err := tr.DB.Create(&domain).Error
	return err
}

// StoreTransaction implements domain_transaction.Repository
func (tr TransactionRepo) StoreTransaction(domain domain_transaction.Transaction) error {
	err := tr.DB.Create(&domain).Error
	return err
}

// GetTransactionByPhone implements domain_transaction.Repository
func (tr TransactionRepo) GetTransactionByPhone(phone string) (transaction []domain_transaction.Transaction) {
	rec := []Transaction{}
	tr.DB.Where("phone = ?", phone).Find(&rec)
	fmt.Println("transaction ", rec, tr.DB.Where("phone = ?", phone).Find(&rec).Error)
	for _, value := range rec {
		transaction = append(transaction, ToDomainTransaction(value))
	}
	return transaction
}

// GetTransactionByPaymentId implements domain_transaction.Repository
func (tr TransactionRepo) GetTransactionByPaymentId(id string) (domain_transaction.Transaction, error) {
	rec := Transaction{}
	err := tr.DB.Where("payment_id", id).First(&rec).Error
	return ToDomainTransaction(rec), err
}

// UpdateTransaction implements domain_transaction.Repository
func (tr TransactionRepo) UpdateTransaction(domain domain_transaction.Transaction) error {
	update := tr.DB.Model(&domain).Where("payment_id = ?", domain.Payment_Id).Updates(domain).RowsAffected
	var err error
	if update == 0 {
		err = errors.New("update failed")
	}
	return err
}

// StorePayment implements domain_transaction.Repository
func (tr TransactionRepo) StorePayment(domain domain_transaction.Payment) error {
	err := tr.DB.Create(&domain).Error
	return err
}

// GetPayment implements domain_transaction.Repository
func (tr TransactionRepo) GetPayment(id string) domain_transaction.Payment {
	rec := Payment{}
	tr.DB.Where("payment_id = ?", id).First(&rec)
	return ToDomainPayment(rec)
}
