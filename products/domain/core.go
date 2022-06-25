package domain_products

import "time"

type Products struct {
	ID           int
	Code         string
	Name         string
	Image        string
	Place_Holder string
	Category_Id  int
	Status       bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Category_Product struct {
	ID        int
	Name      string
	Image     string
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Detail_Product struct {
	ID           int
	Product_Code string
	Name         string
	Code         string
	Price        int
	Description  string
	Status       bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
