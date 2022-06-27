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

type Callback_Invoice struct {
	ID                     string    `json:"id"`
	ExternalID             string    `json:"external_id"`
	UserID                 string    `json:"user_id"`
	IsHigh                 bool      `json:"is_high"`
	PaymentMethod          string    `json:"payment_method"`
	Status                 string    `json:"status"`
	MerchantName           string    `json:"merchant_name"`
	Amount                 int       `json:"amount"`
	PaidAmount             int       `json:"paid_amount"`
	BankCode               string    `json:"bank_code"`
	PaidAt                 time.Time `json:"paid_at"`
	PayerEmail             string    `json:"payer_email"`
	Description            string    `json:"description"`
	AdjustedReceivedAmount int       `json:"adjusted_received_amount"`
	FeesPaidAmount         int       `json:"fees_paid_amount"`
	Updated                time.Time `json:"updated"`
	Created                time.Time `json:"created"`
	Currency               string    `json:"currency"`
	PaymentChannel         string    `json:"payment_channel"`
	PaymentDestination     string    `json:"payment_destination"`
	ID_Customer            string
	Phone                  string
}

func TodomainTransactionFromCallback(req Callback_Invoice) domain_transaction.Transaction {
	return domain_transaction.Transaction{
		Transaction_Code: req.ExternalID,
		ID_Customer:      req.ID_Customer,
		Phone:            req.Phone,
		Amount:           req.Amount,
		Payment_Id:       req.ID,
		Status:           req.Status,
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
