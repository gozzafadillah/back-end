package request

import (
	domain_transaction "ppob/transaction/domain"
	"time"
)

type Transaction struct {
	Transaction_Code    string
	ID_Customer         string
	Phone               string
	Category_Slug       string
	Detail_Product_Slug string
	Amount              int
	Payment_Id          string
	Status              string
}

func TodomainTransaction(req Transaction) domain_transaction.Transaction {
	return domain_transaction.Transaction{
		Transaction_Code:    req.Transaction_Code,
		ID_Customer:         req.ID_Customer,
		Phone:               req.Phone,
		Category_Slug:       req.Category_Slug,
		Detail_Product_Slug: req.Detail_Product_Slug,
		Amount:              req.Amount,
		Payment_Id:          req.Payment_Id,
		Status:              req.Status,
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
	ID             string    `json:"id"`
	PaymentMethod  string    `json:"payment_method"`
	Status         string    `json:"status"`
	PaidAmount     int       `json:"paid_amount"`
	PaymentChannel string    `json:"payment_channel"`
	PaidAt         time.Time `json:"paid_at"`
}

func ToDomainCallBack(req Callback_Invoice) domain_transaction.Callback_Invoice {
	return domain_transaction.Callback_Invoice{
		ID:             req.ID,
		PaymentMethod:  req.PaymentMethod,
		Status:         req.Status,
		PaidAmount:     req.PaidAmount,
		PaymentChannel: req.PaymentChannel,
		PaidAt:         req.PaidAt,
	}
}

type Detail_Transaction struct {
	Product_Detail_code string
	Transaction_Code    string
	ID_Customer         string `json:"id_customer" form:"id_customer" validate:"required"`
	Customer_Name       string `json:"customer_name" form:"customer_name" validate:"required"`
	Price               int
	Fee                 int
}

func TodomainDetail(req Detail_Transaction) domain_transaction.Detail_Transaction {
	return domain_transaction.Detail_Transaction{
		Product_Detail_code: req.Product_Detail_code,
		Transaction_Code:    req.Transaction_Code,
		ID_Customer:         req.ID_Customer,
		Customer_Name:       req.Customer_Name,
		Price:               req.Price,
		Fee:                 2000,
	}
}
