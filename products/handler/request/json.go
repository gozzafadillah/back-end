package request

import (
	domain_products "ppob/products/domain"
)

// Request Product
type RequestJSON struct {
	Name          string `json:"name"`
	Slug          string `json:"slug"`
	Image         string `json:"image"`
	Description   string `json:"description"`
	Price         int    `json:"price"`
	Category_Name string `json:"category_name"`
}

type NewRequest struct {
	Code        string
	Name        string
	Slug        string
	Image       string
	Description string
	Category_Id int
}

// request Category Product
type RequestJSONCategory struct {
	Name string `json:"name"`
}

// Data for Detail Product
type DataDetail struct {
	Code        string
	Price       int
	Description string
	Status      bool
}

func ToDomain(req NewRequest) domain_products.Products {
	return domain_products.Products{
		Code:        req.Code,
		Name:        req.Name,
		Slug:        req.Slug,
		Image:       req.Image,
		Category_Id: req.Category_Id,
		Status:      true,
	}
}

func ToDomainCategory(req RequestJSONCategory) domain_products.Category_Product {
	return domain_products.Category_Product{
		Name: req.Name,
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
