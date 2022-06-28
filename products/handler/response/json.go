package response

import domain_products "ppob/products/domain"

type ResponseJSONProduct struct {
	ID           int
	Name         string
	Product_Slug string
	Category_Id  int
	Image        string
	Status       bool
}

func FromDomainProduct(domain domain_products.Products) ResponseJSONProduct {
	return ResponseJSONProduct{
		ID:           domain.ID,
		Name:         domain.Name,
		Product_Slug: domain.Product_Slug,
		Category_Id:  domain.Category_Id,
		Image:        domain.Image,
		Status:       domain.Status,
	}
}

type ResponseJSONPCategory struct {
	ID     int
	Name   string
	Icon   string
	Status bool
}

func FromDomainCategory(res domain_products.Category_Product) ResponseJSONPCategory {
	return ResponseJSONPCategory{
		ID:     res.ID,
		Name:   res.Name,
		Icon:   res.Icon,
		Status: res.Status,
	}
}

type ResponseJSONDetail struct {
	ID           int
	Product_Slug string
	Name         string
	Detail_Slug  string
	Price        int
	Status       bool
}

func FromDomainDetail(res domain_products.Detail_Product) ResponseJSONDetail {
	return ResponseJSONDetail{
		ID:           res.ID,
		Product_Slug: res.Product_Slug,
		Name:         res.Name,
		Price:        res.Price,
		Detail_Slug:  res.Detail_Slug,
		Status:       res.Status,
	}
}
