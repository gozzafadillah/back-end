package mysql_products

import (
	domain_products "ppob/products/domain"
	"time"

	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
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

func ToDomain(rec Products) domain_products.Products {
	return domain_products.Products{
		ID:          rec.ID,
		Code:        rec.Code,
		Name:        rec.Name,
		Slug:        rec.Slug,
		Description: rec.Description,
		Price:       rec.Price,
		Category_Id: rec.Category_Id,
		Status:      rec.Status,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
}

type Category_Product struct {
	ID     int
	Name   string
	Status bool
}

func ToDomainCategory(rec Category_Product) domain_products.Category_Product {
	return domain_products.Category_Product{
		ID:     rec.ID,
		Name:   rec.Name,
		Status: rec.Status,
	}
}
