package request

import (
	domain_transaction "ppob/transaction/domain"
	"time"
)

type Transaction struct {
	Transaction_Code string
	ID_Customer      string
	Phone            string
	Amount           int
	Payment_Id       string
	Status           string
}

func TodomainTransaction(req Transaction) domain_transaction.Transaction {
	return domain_transaction.Transaction{
		Transaction_Code: req.Transaction_Code,
		ID_Customer:      req.ID_Customer,
		Phone:            req.Phone,
		Amount:           req.Amount,
		Payment_Id:       req.Payment_Id,
		Status:           req.Status,
	}
}

type Payment struct {
	Payment_Id string
	Method     string
	Channel    string
	Paid_at    time.Time
}

func TodomainPayment(req Payment) domain_transaction.Payment {
	return domain_transaction.Payment{
		Payment_Id: req.Payment_Id,
		Method:     req.Method,
		Channel:    req.Channel,
		Paid_at:    req.Paid_at,
	}
}

type Detail_Transaction struct {
	Product_Detail_code string
	Transaction_Code    string
	ID_Customer         string `json:"id_customer" form:"id_customer" validate:"required"`
	Customer_Name       string `json:"customer_name" form:"customer_name" validate:"required"`
	Price               int
	Fee                 int
	Amount              int
}

func TodomainDetail(req Detail_Transaction) domain_transaction.Detail_Transaction {
	return domain_transaction.Detail_Transaction{
		Product_Detail_code: req.Product_Detail_code,
		Transaction_Code:    req.Transaction_Code,
		ID_Customer:         req.ID_Customer,
		Customer_Name:       req.Customer_Name,
		Price:               req.Price,
		Fee:                 req.Fee,
		Amount:              req.Amount,
	}
}
