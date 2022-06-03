package mysql_products

import (
	"errors"
	domain_products "ppob/products/domain"

	"gorm.io/gorm"
)

type ProductsRepo struct {
	DB *gorm.DB
}

func NewProductsRepository(db *gorm.DB) domain_products.Repository {
	return ProductsRepo{
		DB: db,
	}
}

// GetCategoryByName implements domain_products.Repository
func (pr ProductsRepo) GetCategoryByName(name string) (domain_products.Category_Product, error) {
	var rec Category_Product
	err := pr.DB.Where("name = ?", name).First(&rec).Error
	return ToDomainCategory(rec), err
}

// Delete implements domain_products.Repository
func (pr ProductsRepo) Delete(code string) error {
	var rec Products
	err := pr.DB.Unscoped().Delete(&rec, code).Error
	return err
}

// GetAll implements domain_products.Repository
func (pr ProductsRepo) GetAll() ([]domain_products.Products, error) {
	var rec []Products
	err := pr.DB.Find(&rec).Error
	temp := []domain_products.Products{}
	for _, value := range rec {
		temp = append(temp, ToDomain(value))
	}
	return temp, err
}

// GetByCode implements domain_products.Repository
func (pr ProductsRepo) GetByCode(code string) (domain_products.Products, error) {
	var rec Products
	err := pr.DB.Where("code = ?", code).First(&rec).Error
	return ToDomain(rec), err
}

// Store implements domain_products.Repository
func (pr ProductsRepo) Store(domain domain_products.Products) (string, error) {
	err := pr.DB.Save(domain).Error
	return domain.Code, err
}

// Update implements domain_products.Repository
func (pr ProductsRepo) Update(code string, domain domain_products.Products) (domain_products.Products, error) {
	newRecord := map[string]interface{}{
		"Name":        domain.Name,
		"Slug":        domain.Slug,
		"Description": domain.Description,
		"Price":       domain.Price,
		"Category_Id": domain.Category_Id,
		"Status":      domain.Status,
		"UpdatedAt":   domain.UpdatedAt,
	}
	update := pr.DB.Model(&domain).Where("code = ?", code).Updates(newRecord).RowsAffected
	var err error
	if update == 0 {
		err = errors.New("data not found")
	}
	return domain, err
}

// GetCategories implements domain_products.Repository
func (ProductsRepo) GetCategories() ([]domain_products.Category_Product, error) {
	panic("unimplemented")
}

// StoreCategory implements domain_products.Repository
func (ProductsRepo) StoreCategory(domain domain_products.Category_Product) (domain_products.Category_Product, error) {
	panic("unimplemented")
}
