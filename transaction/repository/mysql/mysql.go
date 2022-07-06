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

// Count implements domain_transaction.Repository
func (tr TransactionRepo) Count(cat string, phone string, id_customer string, detail_product string) (string, int) {
	var rec Transaction
	tr.DB.Raw("SELECT  * FROM `transactions` INNER JOIN `detail_transactions`ON `detail_transactions`.`transaction_code` = `transactions`.`transaction_code` WHERE transactions.category_slug = ? AND transactions.status = 'PAID' AND transactions.phone = ? AND transactions.id_customer = ? AND transactions.detail_product_slug = ?", cat, phone, id_customer, detail_product).Scan(&rec)
	fmt.Println("rec :", rec)
	var counting int
	tr.DB.Raw("SELECT COUNT(*) FROM `transactions` INNER JOIN `detail_transactions`ON `detail_transactions`.`transaction_code` = `transactions`.`transaction_code` WHERE transactions.category_slug = ? AND transactions.status = 'PAID' AND transactions.phone = ? AND transactions.id_customer = ? AND transactions.detail_product_slug = ?", cat, phone, id_customer, detail_product).Scan(&counting)
	fmt.Println("count :", counting)

	return rec.Payment_Id, counting
}

// GetFavorite implements domain_transaction.Repository
func (tr TransactionRepo) GetFavorite(cat, phone string) []domain_transaction.Transaction {
	rec := []Transaction{}
	sliceTransaction := []domain_transaction.Transaction{}
	tr.DB.Raw("SELECT * FROM transactions INNER JOIN detail_transactions ON detail_transactions.transaction_code = transactions.transaction_code	WHERE  transactions.phone = ? AND transactions.category_slug = ? AND status = 'PAID'", phone, cat).Scan(&rec)
	fmt.Println("cek :", rec)

	for _, value := range rec {
		sliceTransaction = append(sliceTransaction, ToDomainTransaction(value))
	}
	return sliceTransaction
}

func NewTransactionRepo(db *gorm.DB) domain_transaction.Repository {
	return TransactionRepo{
		DB: db,
	}
}

// StoreDetailTransaction implements domain_transaction.Repository
func (tr TransactionRepo) StoreDetailTransaction(domain domain_transaction.Detail_Transaction) error {
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
