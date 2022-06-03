package request

import (
	domain_products "ppob/products/domain"
)

type RequestJSON struct {
	Name          string `json:"name"`
	Slug          string `json:"slug"`
	Description   string `json:"description"`
	Price         int    `json:"price"`
	Category_Name string `json:"category_name"`
}

type NewRequest struct {
	Code        string
	Name        string
	Slug        string
	Description string
	Price       int
	Category_Id int
}

func ToDomain(req NewRequest) domain_products.Products {
	return domain_products.Products{
		Code:        req.Code,
		Name:        req.Name,
		Slug:        req.Slug,
		Description: req.Description,
		Price:       req.Price,
		Category_Id: req.Category_Id,
	}
}
