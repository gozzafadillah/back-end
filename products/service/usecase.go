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

// GetProduct implements domain_products.Service
func (ps ProductService) GetProduct(id int) (domain_products.Products, error) {
	data, err := ps.Repository.GetByID(id)
	if err != nil {
		return domain_products.Products{}, errors.New("get product failed")
	}
	return data, nil
}

// GetProducts implements domain_products.Service
func (ps ProductService) GetProducts() ([]domain_products.Products, error) {
	datas, err := ps.Repository.GetAll()
	if err != nil {
		return []domain_products.Products{}, errors.New("get products failed")
	}
	return datas, nil
}

// InsertData implements domain_products.Service
func (ps ProductService) InsertData(domain domain_products.Products) (domain_products.Products, error) {
	data, err := ps.Repository.Store(domain)
	if err != nil {
		return domain_products.Products{}, errors.New("store data product failed")
	}

	return data, nil
}

// Delete implements domain_products.Service
func (ps ProductService) Destroy(id int) error {
	err := ps.Repository.Delete(id)
	if err != nil {
		return errors.New("delete failed")
	}
	return nil
}

// Edit implements domain_products.Service
func (ps ProductService) Edit(id int, domain domain_products.Products) (domain_products.Products, error) {
	data, err := ps.Repository.Update(id, domain)
	if err != nil {
		return domain_products.Products{}, errors.New("update failed")
	}
	return data, nil
}

// InsertCategory implements domain_products.Service
func (ps ProductService) InsertCategory(domain domain_products.Category_Product) (domain_products.Category_Product, error) {
	data, err := ps.Repository.StoreCategory(domain)
	if err != nil {
		return domain_products.Category_Product{}, errors.New("insert data failed")
	}
	return data, nil
}

// GetCategory implements domain_products.Service
func (ps ProductService) GetCategory(name string) (domain_products.Category_Product, error) {
	data, err := ps.Repository.GetCategoryByName(name)
	if err != nil {
		return domain_products.Category_Product{}, errors.New("data category not found")
	}
	return data, nil
}

// GetCategories implements domain_products.Service
func (ps ProductService) GetCategories() ([]domain_products.Category_Product, error) {
	data, err := ps.Repository.GetCategories()
	if err != nil {
		return []domain_products.Category_Product{}, errors.New("data empty")
	}
	return data, nil
}

// DestroyCategory implements domain_products.Service
func (ps ProductService) DestroyCategory(id int) error {
	err := ps.Repository.DeleteCategory(id)
	if err != nil {
		return errors.New("data failed delete")
	}
	return nil
}

// InsertDetail implements domain_products.Service
func (ps ProductService) InsertDetail(domain domain_products.Detail_Product) (domain_products.Detail_Product, error) {
	data, err := ps.Repository.StoreDetail(domain)
	if err != nil {
		return domain_products.Detail_Product{}, errors.New("insert data failed")
	}
	return data, nil
}

// GetDetails implements domain_products.Service
func (ps ProductService) GetDetails(id int) ([]domain_products.Detail_Product, error) {
	data, err := ps.Repository.GetDetaislByID(id)
	if err != nil {
		return []domain_products.Detail_Product{}, errors.New("details product not found")
	}
	return data, nil
}

// DestroyDetail implements domain_products.Service
func (ps ProductService) DestroyDetail(id int) error {
	err := ps.Repository.DeleteDetail(id)
	if err != nil {
		return errors.New("data failed destroy")
	}
	return nil
}
