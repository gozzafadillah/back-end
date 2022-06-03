package domain_products

import "time"

type Products struct {
	ID          int
	Code        string
	Name        string
	Slug        string
	Description string
	Price       int
	Category_Id int
	Status      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Category_Product struct {
	ID     int
	Name   string
	Status bool
}
