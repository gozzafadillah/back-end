package domain_products

import "time"

type Products struct {
	ID           int
	Product_Slug string
	Name         string
	Image        string
	Category_Id  int
	Status       bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Category_Product struct {
	ID            int
	Name          string
	Category_Slug string
	Image         string
	Status        bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Detail_Product struct {
	ID           int
	Product_Slug string
	Name         string
	Detail_Slug  string
	Price        int
	Status       bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
