package domain_transaction

import (
	"time"
)

type Transaction struct {
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

type Payment struct {
	ID         int
	Payment_Id string
	Method     string
	Channel    string
	Paid_at    time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Detail_Transaction struct {
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
