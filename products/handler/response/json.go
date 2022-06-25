package response

import domain_products "ppob/products/domain"

type ResponseJSONProduct struct {
	ID           int
	Name         string
	Code         string
	Place_Holder string
	Category_Id  int
	Image        string
	Status       bool
}

func FromDomainProduct(domain domain_products.Products) ResponseJSONProduct {
	return ResponseJSONProduct{
		ID:           domain.ID,
		Name:         domain.Name,
		Code:         domain.Code,
		Place_Holder: domain.Place_Holder,
		Category_Id:  domain.Category_Id,
		Image:        domain.Image,
		Status:       domain.Status,
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
	ID           int
	Product_Code string
	Name         string
	Price        int
	Description  string
	Status       bool
}

func FromDomainDetail(res domain_products.Detail_Product) ResponseJSONDetail {
	return ResponseJSONDetail{
		ID:           res.ID,
		Product_Code: res.Product_Code,
		Price:        res.Price,
		Description:  res.Description,
		Status:       res.Status,
	}
}
