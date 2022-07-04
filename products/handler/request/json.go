package request

import (
	domain_products "ppob/products/domain"
)

// Request Product
type RequestJSONProduct struct {
	Product_Slug string
	Name         string      `json:"name" form:"name" validate:"required"`
	Image        string      `json:"img" form:"img"`
	Category_id  int         `json:"category_id" form:"category_id"`
	File         interface{} `json:"file,omitempty"`
}

func ToDomain(req RequestJSONProduct) domain_products.Products {
	return domain_products.Products{
		Product_Slug: req.Product_Slug,
		Name:         req.Name,
		Image:        req.Image,
		Category_Id:  req.Category_id,
		Status:       true,
	}
}

// request Category Product
type RequestJSONCategory struct {
	Name          string `json:"name" form:"name" validate:"required"`
	Category_Slug string
	Icon          string `json:"icon" form:"icon" validate:"required"`
}

func ToDomainCategory(req RequestJSONCategory) domain_products.Category_Product {
	return domain_products.Category_Product{
		Name:          req.Name,
		Category_Slug: req.Category_Slug,
		Icon:          req.Icon,
		Status:        true,
	}
}

// Data for Detail Product
type DataDetail struct {
	Product_Slug string
	Name         string `json:"name" form:"name" validate:"required"`
	Detail_Slug  string
	Price        int `json:"price" form:"price" validate:"required"`
	Status       bool
}

func ToDomainDetail(data DataDetail) domain_products.Detail_Product {
	return domain_products.Detail_Product{
		Product_Slug: data.Product_Slug,
		Name:         data.Name,
		Detail_Slug:  data.Detail_Slug,
		Price:        data.Price,
		Status:       true,
	}
}
