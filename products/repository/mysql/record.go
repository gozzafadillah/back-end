package mysql_products

import (
	domain_products "ppob/products/domain"
	"time"

	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
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

func ToDomain(rec Products) domain_products.Products {
	return domain_products.Products{
		ID:           rec.ID,
		Code:         rec.Code,
		Name:         rec.Name,
		Image:        rec.Image,
		Place_Holder: rec.Place_Holder,
		Category_Id:  rec.Category_Id,
		Status:       rec.Status,
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	}
}

type Category_Product struct {
	gorm.Model
	ID        int
	Name      string
	Image     string
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ToDomainCategory(rec Category_Product) domain_products.Category_Product {
	return domain_products.Category_Product{
		ID:     rec.ID,
		Name:   rec.Name,
		Image:  rec.Image,
		Status: rec.Status,
	}
}

type Detail_Product struct {
	gorm.Model
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

func ToDomainDetail(rec Detail_Product) domain_products.Detail_Product {
	return domain_products.Detail_Product{
		ID:           rec.ID,
		Product_Code: rec.Product_Code,
		Name:         rec.Name,
		Code:         rec.Code,
		Price:        rec.Price,
		Description:  rec.Description,
		Status:       rec.Status,
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	}
}
