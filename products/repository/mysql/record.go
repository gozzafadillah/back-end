package mysql_products

import (
	domain_products "ppob/products/domain"
	"time"

	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	ID           int
	Product_Slug string `gorm:"unique"`
	Name         string
	Image        string
	Category_Id  int
	Status       bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func ToDomainProduct(rec Products) domain_products.Products {
	return domain_products.Products{
		ID:           rec.ID,
		Product_Slug: rec.Product_Slug,
		Name:         rec.Name,
		Image:        rec.Image,
		Category_Id:  rec.Category_Id,
		Status:       rec.Status,
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	}
}

type Category_Product struct {
	gorm.Model
	ID            int
	Name          string
	Category_Slug string `gorm:"unique"`
	Image         string
	Status        bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func ToDomainCategory(rec Category_Product) domain_products.Category_Product {
	return domain_products.Category_Product{
		ID:            rec.ID,
		Name:          rec.Name,
		Category_Slug: rec.Category_Slug,
		Image:         rec.Image,
		Status:        rec.Status,
	}
}

type Detail_Product struct {
	gorm.Model
	ID           int
	Product_Slug string `gorm:"unique"`
	Name         string
	Detail_Slug  string
	Price        int
	Status       bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func ToDomainDetail(rec Detail_Product) domain_products.Detail_Product {
	return domain_products.Detail_Product{
		ID:           rec.ID,
		Product_Slug: rec.Product_Slug,
		Name:         rec.Name,
		Detail_Slug:  rec.Detail_Slug,
		Price:        rec.Price,
		Status:       rec.Status,
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	}
}
