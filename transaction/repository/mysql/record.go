package mysql_transaction

import (
	domain_transaction "ppob/transaction/domain"
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID               int
	Transaction_Code string
	ID_Customer      string
	Phone            string
	Amount           int
	Payment_Id       string
	Status           string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func ToDomainTransaction(rec Transaction) domain_transaction.Transaction {
	return domain_transaction.Transaction{
		ID:               rec.ID,
		Transaction_Code: rec.Transaction_Code,
		ID_Customer:      rec.ID_Customer,
		Phone:            rec.Phone,
		Amount:           rec.Amount,
		Payment_Id:       rec.Payment_Id,
		Status:           rec.Status,
		CreatedAt:        rec.CreatedAt,
		UpdatedAt:        rec.UpdatedAt,
	}
}

type Payment struct {
	gorm.Model
	ID         int
	Payment_Id string
	Method     string
	Channel    string
	Paid_at    time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func ToDomainPayment(rec Payment) domain_transaction.Payment {
	return domain_transaction.Payment{
		ID:         rec.ID,
		Payment_Id: rec.Payment_Id,
		Method:     rec.Method,
		Channel:    rec.Channel,
		Paid_at:    rec.Paid_at,
		CreatedAt:  rec.CreatedAt,
		UpdatedAt:  rec.UpdatedAt,
	}
}

type Detail_Transaction struct {
	gorm.Model
	ID                  int
	Product_Detail_code string
	Transaction_Code    string
	ID_Customer         string
	Customer_Name       string
	Price               int
	Fee                 int
	Amount              int
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

func ToDomain(rec Detail_Transaction) domain_transaction.Detail_Transaction {
	return domain_transaction.Detail_Transaction{
		ID:                  rec.ID,
		Product_Detail_code: rec.Product_Detail_code,
		Transaction_Code:    rec.Transaction_Code,
		ID_Customer:         rec.ID_Customer,
		Customer_Name:       rec.Customer_Name,
		Price:               rec.Price,
		Fee:                 rec.Fee,
		Amount:              rec.Amount,
		CreatedAt:           rec.CreatedAt,
		UpdatedAt:           rec.UpdatedAt,
	}
}
