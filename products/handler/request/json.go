package request

import (
	domain_products "ppob/products/domain"
)

// Request Product
type RequestJSON struct {
	Code         string
	Name         string `json:"name" form:"name" validate:"required"`
	Image        string `json:"img" form:"img"`
	Place_Holder string `json:"placeholder" form:"placeholder"`
	Category_id  int
	File         interface{} `json:"file,omitempty"`
}

// request Category Product
type RequestJSONCategory struct {
	Name  string      `json:"name" form:"name" validate:"required"`
	Image string      `json:"img" form:"img"`
	File  interface{} `json:"file,omitempty"`
}

// Data for Detail Product
type DataDetail struct {
	Product_Code string
	Name         string `json:"name" form:"name"`
	Code         string
	Description  string `json:"description" form:"description"`
	Price        int    `json:"price" form:"price"`
	Status       bool
}

func ToDomain(req RequestJSON) domain_products.Products {
	return domain_products.Products{
		Code:         req.Code,
		Name:         req.Name,
		Place_Holder: req.Place_Holder,
		Image:        req.Image,
		Category_Id:  req.Category_id,
		Status:       true,
	}
}

func ToDomainCategory(req RequestJSONCategory) domain_products.Category_Product {
	return domain_products.Category_Product{
		Name:   req.Name,
		Image:  req.Image,
		Status: true,
	}
}

func ToDomainDetail(data DataDetail) domain_products.Detail_Product {
	return domain_products.Detail_Product{
		Product_Code: data.Product_Code,
		Name:         data.Name,
		Code:         data.Code,
		Price:        data.Price,
		Description:  data.Description,
		Status:       true,
	}
}
