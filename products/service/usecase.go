package service_products

import (
	"errors"
	domain_products "ppob/products/domain"
)

type ProductService struct {
	Repository domain_products.Repository
}

func NewProductsService(repo domain_products.Repository) domain_products.Service {
	return ProductService{
		Repository: repo,
	}
}

// GetCategory implements domain_products.Service
func (ProductService) GetCategory(name string) (domain_products.Category_Product, error) {
	panic("unimplemented")
}

// Delete implements domain_products.Service
func (ps ProductService) Delete(code string) error {
	err := ps.Repository.Delete(code)
	if err != nil {
		return errors.New("delete failed")
	}
	return nil
}

// Edit implements domain_products.Service
func (ps ProductService) Edit(code string, domain domain_products.Products) (domain_products.Products, error) {
	data, err := ps.Repository.Update(code, domain)
	if err != nil {
		return domain_products.Products{}, errors.New("update fail")
	}
	return data, nil
}

// GetProduct implements domain_products.Service
func (ps ProductService) GetProduct(code string) (domain_products.Products, error) {
	data, err := ps.Repository.GetByCode(code)
	if err != nil {
		return domain_products.Products{}, errors.New("get product fail")
	}
	return data, nil
}

// GetProducts implements domain_products.Service
func (ps ProductService) GetProducts() ([]domain_products.Products, error) {
	datas, err := ps.Repository.GetAll()
	if err != nil {
		return []domain_products.Products{}, errors.New("get products fail")
	}
	return datas, nil
}

// InsertData implements domain_products.Service
func (ps ProductService) InsertData(domain domain_products.Products) (domain_products.Products, error) {
	code, err := ps.Repository.Store(domain)
	if err != nil {
		return domain_products.Products{}, errors.New("store data fail")
	}
	data, err := ps.Repository.GetByCode(code)
	if err != nil {
		return domain_products.Products{}, errors.New("fail search data store")
	}
	return data, nil
}

// InsertCategory implements domain_products.Service
func (ProductService) InsertCategory(domain domain_products.Category_Product) (domain_products.Category_Product, error) {
	panic("unimplemented")
}
