package request

import (
	domain_products "ppob/products/domain"
)

// Request Product
type RequestJSON struct {
	Code        string
	Name        string `json:"name" validate:"required"`
	Image       string `json:"image"`
	Category_id int    `json:"category_id" validate:"required"`
}

// request Category Product
type RequestJSONCategory struct {
	Name string `json:"name" validate:"required"`
}

// Data for Detail Product
type DataDetail struct {
	Code        string
	Description string `json:"description"`
	Price       int    `json:"price" validate:"required"`
	Status      bool
}

func ToDomain(req RequestJSON) domain_products.Products {
	return domain_products.Products{
		Code:        req.Code,
		Name:        req.Name,
		Image:       req.Image,
		Category_Id: req.Category_id,
		Status:      true,
	}
}

func ToDomainCategory(req RequestJSONCategory) domain_products.Category_Product {
	return domain_products.Category_Product{
		Name:   req.Name,
		Status: true,
	}
}

func ToDomainDetail(data DataDetail) domain_products.Detail_Product {
	return domain_products.Detail_Product{
		Code:        data.Code,
		Price:       data.Price,
		Description: data.Description,
		Status:      true,
	}
}
