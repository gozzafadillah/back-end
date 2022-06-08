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
func (pr ProductsRepo) GetByID(id int) (domain_products.Products, error) {
	var rec Products
	err := pr.DB.Where("id = ?", id).First(&rec).Error
	return ToDomain(rec), err
}

// Store implements domain_products.Repository
func (pr ProductsRepo) Store(domain domain_products.Products) (domain_products.Products, error) {
	err := pr.DB.Save(domain).Error
	return domain, err
}

// Update implements domain_products.Repository
func (pr ProductsRepo) Update(id int, domain domain_products.Products) (domain_products.Products, error) {
	newRecord := map[string]interface{}{
		"Name":        domain.Name,
		"Slug":        domain.Slug,
		"Category_Id": domain.Category_Id,
		"Status":      domain.Status,
		"UpdatedAt":   domain.UpdatedAt,
	}
	update := pr.DB.Model(&domain).Where("id = ?", id).Updates(newRecord).RowsAffected
	var err error
	if update == 0 {
		err = errors.New("data not found")
	}
	return domain, err
}

// Delete implements domain_products.Repository
func (pr ProductsRepo) Delete(id int) error {
	var rec Products
	err := pr.DB.Unscoped().Delete(&rec, id).Error
	return err
}

// GetCategories implements domain_products.Repository
func (pr ProductsRepo) GetCategories() ([]domain_products.Category_Product, error) {
	var rec []Category_Product
	err := pr.DB.Find(&rec).Error
	sliceRec := []domain_products.Category_Product{}
	for _, value := range rec {
		sliceRec = append(sliceRec, ToDomainCategory(value))
	}
	return sliceRec, err
}

// StoreCategory implements domain_products.Repository
func (pr ProductsRepo) StoreCategory(domain domain_products.Category_Product) (domain_products.Category_Product, error) {
	err := pr.DB.Save(&domain).Error
	return domain, err
}

// GetCategoryByName implements domain_products.Repository
func (pr ProductsRepo) GetCategoryByName(name string) (domain_products.Category_Product, error) {
	var rec Category_Product
	err := pr.DB.Where("name = ?", name).First(&rec).Error
	return ToDomainCategory(rec), err
}

// GetDetailsByID implements domain_products.Repository
func (pr ProductsRepo) GetDetailsByID(id int) ([]domain_products.Detail_Product, error) {
	var rec []Detail_Product
	err := pr.DB.Where("id = ?", id).First(&rec).Error
	var sliceRec []domain_products.Detail_Product
	for _, value := range rec {
		sliceRec = append(sliceRec, ToDomainDetail(value))
	}
	return sliceRec, err
}

// StoreDetail implements domain_products.Repository
func (pr ProductsRepo) StoreDetail(domain domain_products.Detail_Product) (domain_products.Detail_Product, error) {
	err := pr.DB.Save(&domain).Error
	return domain, err
}

// DeleteCategory implements domain_products.Repository
func (pr ProductsRepo) DeleteCategory(id int) error {
	var rec Category_Product
	err := pr.DB.Unscoped().Delete(&rec, id).Error
	return err
}

// GetDetaislByID implements domain_products.Repository
func (pr ProductsRepo) GetDetaislByID(id int) ([]domain_products.Detail_Product, error) {
	var rec []Detail_Product
	err := pr.DB.Where("id = ?", id).First(&rec).Error
	var sliceRec []domain_products.Detail_Product
	for _, value := range rec {
		sliceRec = append(sliceRec, ToDomainDetail(value))
	}
	return sliceRec, err
}

// DeleteDetail implements domain_products.Repository
func (pr ProductsRepo) DeleteDetail(id int) error {
	var rec Detail_Product
	err := pr.DB.Unscoped().Delete(&rec, id).Error
	return err
}
