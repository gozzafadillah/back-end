package mysql_products

import (
	"errors"
	"fmt"
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

// GetByCategory implements domain_products.Repository
func (pr ProductsRepo) GetByCategory(id int) []domain_products.Products {
	rec := []Products{}
	sliceProduct := []domain_products.Products{}
	pr.DB.Where("category_id = ?", id).Find(&rec)
	for _, value := range rec {
		sliceProduct = append(sliceProduct, ToDomain(value))
	}
	return sliceProduct
}

// GetByCode implements domain_products.Repository
func (pr ProductsRepo) GetByID(id int) (domain_products.Products, error) {
	var rec Products
	err := pr.DB.Where("id = ?", id).First(&rec).Error
	return ToDomain(rec), err
}

// Store implements domain_products.Repository
func (pr ProductsRepo) Store(domain domain_products.Products) error {
	err := pr.DB.Save(&domain).Error
	return err
}

// Update implements domain_products.Repository
func (pr ProductsRepo) Update(id int, domain domain_products.Products) error {
	newRecord := map[string]interface{}{
		"Name":        domain.Name,
		"Code":        domain.Code,
		"Category_Id": domain.Category_Id,
		"Image":       domain.Image,
	}
	update := pr.DB.Model(&domain).Where("id = ?", id).Updates(newRecord).RowsAffected
	var err error
	if update == 0 {
		err = errors.New("update failed")
	}
	return err
}

// Delete implements domain_products.Repository
func (pr ProductsRepo) Delete(id int) error {
	var rec Products
	err := pr.DB.Unscoped().Delete(&rec, id).Error
	return err
}

// GetDetailsByCode implements domain_products.Repository
func (pr ProductsRepo) GetDetailsByCode(code string) ([]domain_products.Detail_Product, error) {
	var rec []Detail_Product
	err := pr.DB.Where("code = ?", code).Find(&rec).Error
	var sliceRec []domain_products.Detail_Product
	for _, value := range rec {
		sliceRec = append(sliceRec, ToDomainDetail(value))
	}
	return sliceRec, err
}

// StoreDetail implements domain_products.Repository
func (pr ProductsRepo) StoreDetail(code string, domain domain_products.Detail_Product) error {
	cek := pr.DB.Where("code = ?", code).First(&Products{}).Error
	if cek == nil {
		domain.Code = code
		err := pr.DB.Save(&domain).Error
		return err
	}
	return cek
}

// UpdateDetails implements domain_products.Repository
func (pr ProductsRepo) UpdateDetails(codeLama string, codeBaru string) error {
	rec := Detail_Product{}
	err := pr.DB.Model(&rec).Where("code = ?", codeLama).Update("code", codeBaru).Error
	return err
}

// UpdateDetail implements domain_products.Repository
func (pr ProductsRepo) UpdateDetail(id int, domain domain_products.Detail_Product) error {
	newRecord := map[string]interface{}{
		"Description": domain.Description,
		"Price":       domain.Price,
	}
	fmt.Println("data ", newRecord)
	update := pr.DB.Model(&domain).Where("id = ?", id).Updates(newRecord).RowsAffected
	if update == 0 {
		return errors.New("update failed")
	}
	return nil
}

// DeleteDetail implements domain_products.Repository
func (pr ProductsRepo) DeleteDetail(id int) error {
	var rec Detail_Product
	err := pr.DB.Unscoped().Delete(&rec, id).RowsAffected
	if err == 0 {
		return errors.New("delete failed")
	}
	return nil
}

// DeleteCategory implements domain_products.Repository
func (pr ProductsRepo) DeleteCategory(id int) error {
	var rec Category_Product
	err := pr.DB.Unscoped().Delete(&rec, id).RowsAffected
	if err == 0 {
		return errors.New("delete failed")
	}
	return nil
}

// DeleteDetails implements domain_products.Repository
func (pr ProductsRepo) DeleteDetails(code string) error {
	rec := Detail_Product{}
	err := pr.DB.Unscoped().Where("code = ?", code).Delete(&rec).RowsAffected
	if err == 0 {
		return errors.New("delete failed")
	}
	return nil
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
func (pr ProductsRepo) StoreCategory(domain domain_products.Category_Product) error {
	err := pr.DB.Save(&domain).Error
	return err
}

// UpdateCategory implements domain_products.Repository
func (pr ProductsRepo) UpdateCategory(id int, domain domain_products.Category_Product) error {
	var rec = Category_Product{}
	newRecord := map[string]interface{}{
		"Name": domain.Name,
	}
	var err error
	update := pr.DB.Model(&rec).Where("id = ?", id).Updates(newRecord).RowsAffected
	if update == 0 {
		err = errors.New("update failed")
	}
	return err
}

// GetCategoryById implements domain_products.Repository
func (pr ProductsRepo) GetCategoryById(id int) (domain_products.Category_Product, error) {
	var rec Category_Product
	err := pr.DB.Where("id = ?", id).First(&rec).Error
	return ToDomainCategory(rec), err
}
