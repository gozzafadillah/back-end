package domain_transaction

import (
	"time"
)

type Transaction struct {
	ID               int
	Transaction_Code string
	ID_Pelanggan     string
	Phone            string
	Amount           int
	Payment_Id       string
	Status           string
}

type Payment struct {
	Payment_Id string
	Method     string
	Channel    string
	Paid_at    time.Time
}

type Detail_Transaction struct {
	Product_Detail_code string
	Transaction_Code    string
}
