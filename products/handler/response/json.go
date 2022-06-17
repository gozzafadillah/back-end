package response

import domain_products "ppob/products/domain"

type ResponseJSONProduct struct {
	ID          int
	Name        string
	Code        string
	Category_Id int
	Status      bool
}

func FromDomainProduct(domain domain_products.Products) ResponseJSONProduct {
	return ResponseJSONProduct{
		ID:          domain.ID,
		Name:        domain.Name,
		Code:        domain.Code,
		Category_Id: domain.Category_Id,
		Status:      domain.Status,
	}
}

type ResponseJSONPCategory struct {
	ID     int
	Name   string
	Status bool
}

func FromDomainCategory(res domain_products.Category_Product) ResponseJSONPCategory {
	return ResponseJSONPCategory{
		ID:     res.ID,
		Name:   res.Name,
		Status: res.Status,
	}
}

type ResponseJSONDetail struct {
	ID          int
	Price       int
	Description string
	Status      bool
}

func FromDomainDetail(res domain_products.Detail_Product) ResponseJSONDetail {
	return ResponseJSONDetail{
		ID:          res.ID,
		Price:       res.Price,
		Description: res.Description,
		Status:      res.Status,
	}
}
