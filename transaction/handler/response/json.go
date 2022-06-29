package response

import (
	domain_products "ppob/products/domain"
	domain_transaction "ppob/transaction/domain"
)

type ResponseJSONDetailTransaction struct {
	Product_Detail_code string
	Transaction_Code    string
	ID_Customer         string
	Customer_Name       string
	Price               int
	Fee                 int
}

func FromDomainCheckout(domain domain_transaction.Detail_Transaction) ResponseJSONDetailTransaction {
	return ResponseJSONDetailTransaction{
		Product_Detail_code: domain.Product_Detail_code,
		Transaction_Code:    domain.Transaction_Code,
		ID_Customer:         domain.ID_Customer,
		Customer_Name:       domain.Customer_Name,
		Price:               domain.Price,
		Fee:                 domain.Fee,
	}
}

type ResponseJSONProduct struct {
	Name  string
	Image string
}

func FromDomainProduct(domain domain_products.Products) ResponseJSONProduct {
	return ResponseJSONProduct{
		Name:  domain.Name,
		Image: domain.Image,
	}
}

type ResponseJSONPCategory struct {
	Name string
}

func FromDomainCatProduct(domain domain_products.Category_Product) ResponseJSONPCategory {
	return ResponseJSONPCategory{
		Name: domain.Name,
	}
}
