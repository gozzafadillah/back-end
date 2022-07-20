package handler_products

import (
	domain_products "ppob/products/domain"

	"github.com/go-playground/validator/v10"
)

type ProductsHandler struct {
	Service    domain_products.Service
	Validation *validator.Validate
}

func NewProductsHandler(service domain_products.Service) ProductsHandler {
	return ProductsHandler{
		Service:    service,
		Validation: validator.New(),
	}
}
