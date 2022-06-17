package service_products

import (
	"errors"
	"fmt"
	domain_products "ppob/products/domain"
	"sort"

	"ppob/helper/slug"
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

// GetProductByCategory implements domain_products.Service
func (ps ProductService) GetProductByCategory(id int) []domain_products.Products {
	data := ps.Repository.GetByCategory(id)
	return data
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
func (ps ProductService) InsertData(domain domain_products.Products) error {
	domain.Code = slug.GenerateSlug(domain.Name)
	err := ps.Repository.Store(domain)
	if err != nil {
		return errors.New("internal server error")
	}

	return nil
}

// Delete implements domain_products.Service
func (ps ProductService) Destroy(id int) error {
	data, err := ps.GetProduct(id)
	if err != nil {
		return errors.New("data not found")
	}

	err = ps.Repository.Delete(data.ID)
	if err != nil {
		return errors.New("delete failed")
	}

	err = ps.Repository.DeleteDetails(data.Code)
	if err != nil {
		return errors.New("delete failed")
	}

	return nil
}

// Edit implements domain_products.Service
func (ps ProductService) Edit(id int, domain domain_products.Products) error {
	data, err := ps.GetProduct(id)
	if err != nil {
		return errors.New("data not found")
	}
	domain.Code = slug.GenerateSlug(domain.Name)
	err = ps.Repository.Update(id, domain)
	if err != nil {
		return errors.New("update failed")
	}
	err = ps.Repository.UpdateDetails(data.Code, domain.Code)
	if err != nil {
		return errors.New("update failed")
	}
	return nil
}

// InsertDetail implements domain_products.Service
func (ps ProductService) InsertDetail(code string, domain domain_products.Detail_Product) error {

	err := ps.Repository.StoreDetail(code, domain)
	if err != nil {
		return errors.New("insert data failed")
	}
	return nil
}

// GetDetails implements domain_products.Service
func (ps ProductService) GetDetails(code string) []domain_products.Detail_Product {
	data, err := ps.Repository.GetDetailsByCode(code)
	sort.Slice(data, func(i, j int) bool {
		return data[i].Price < data[j].Price
	})
	for _, v := range data {
		fmt.Println(v)
	}
	if err != nil {
		return []domain_products.Detail_Product{}
	}
	return data
}

// EditDetail implements domain_products.Service
func (ps ProductService) EditDetail(id int, domain domain_products.Detail_Product) error {
	err := ps.Repository.UpdateDetail(id, domain)
	if err != nil {
		return err
	}
	return nil
}

// DestroyDetail implements domain_products.Service
func (ps ProductService) DestroyDetail(id int) error {
	err := ps.Repository.DeleteDetail(id)
	if err != nil {
		return errors.New("data failed destroy")
	}
	return nil
}

// InsertCategory implements domain_products.Service
func (ps ProductService) InsertCategory(domain domain_products.Category_Product) error {
	err := ps.Repository.StoreCategory(domain)
	if err != nil {
		return errors.New("insert data failed")
	}
	return nil
}

// GetCategory implements domain_products.Service
func (ps ProductService) GetCategory(id int) (domain_products.Category_Product, error) {
	data, err := ps.Repository.GetCategoryById(id)
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

// EditCategory implements domain_products.Service
func (ps ProductService) EditCategory(id int, domain domain_products.Category_Product) error {
	err := ps.Repository.UpdateCategory(id, domain)
	if err != nil {
		return err
	}
	return nil
}

// DestroyCategory implements domain_products.Service
func (ps ProductService) DestroyCategory(id int) error {
	err := ps.Repository.DeleteCategory(id)
	if err != nil {
		return errors.New("data failed delete")
	}
	return nil
}
