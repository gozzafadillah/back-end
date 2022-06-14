package mysql_users

import (
	"errors"
	"fmt"
	"ppob/helper/encryption"
	domain_users "ppob/users/domain"

	"gorm.io/gorm"
)

type UsersRepo struct {
	DB *gorm.DB
}

func NewUsersRepo(db *gorm.DB) domain_users.Repository {
	return UsersRepo{
		DB: db,
	}
}

// GetAllUser implements domain_users.Repository
func (ur UsersRepo) GetAllUser() ([]domain_users.Users, error) {
	var rec []Users
	err := ur.DB.Find(&rec).Error
	var SliceRec []domain_users.Users
	for _, value := range rec {
		SliceRec = append(SliceRec, ToDomain(value))
	}
	return SliceRec, err
}

// GetByPhone implements domain_users.Repository
func (ur UsersRepo) GetByPhone(phone string) (domain_users.Users, error) {
	var rec Users
	err := ur.DB.Where("phone = ?", phone).First(&rec).Error
	return ToDomain(rec), err
}

// // GetById implements domain_users.Repository
// func (ur UsersRepo) GetById(id int) (domain_users.Users, error) {
// 	rec := Users{}
// 	err := ur.DB.Where("id=?", id).First(&rec).Error
// 	return ToDomain(rec), err
// }

// Store implements domain_users.Repository
func (ur UsersRepo) Store(domain domain_users.Users) (string, error) {
	err := ur.DB.Save(&domain).Error
	return domain.Phone, err
}

// Update implements domain_users.Repository
func (ur UsersRepo) Update(phone string, domain domain_users.Users) error {
	data := map[string]interface{}{
		"Name":     domain.Name,
		"Email":    domain.Email,
		"Password": domain.Password,
		"Phone":    domain.Phone,
		"DOB":      domain.DOB,
		"Image":    domain.Image,
	}
	fmt.Println("data update :", data)
	err := ur.DB.Model(&domain).Where("phone = ?", phone).Updates(data).RowsAffected
	if err == 0 {
		return errors.New("users not found")
	}
	return nil
}

// CheckEmailPassword implements domain_users.Repository
func (ur UsersRepo) CheckEmailPassword(email string, password string) (domain_users.Users, error) {
	var rec Users
	err := ur.DB.Where("email = ?", email).First(&rec).Error
	if err != nil {
		return domain_users.Users{}, errors.New("data not found")
	}

	data := encryption.CheckPasswordHash(password, rec.Password)
	fmt.Println(data)
	if !data || !rec.Status {
		return domain_users.Users{}, errors.New("password miss match")
	}

	return ToDomain(rec), nil
}

// StoreAccount implements domain_users.Repository
func (ur UsersRepo) StoreAccount(domain domain_users.Account) (domain_users.Account, error) {
	err := ur.DB.Save(&domain).Error
	return domain, err
}

// GetUserAccount implements domain_users.Repository
func (ur UsersRepo) GetUserAccount(phone string) (domain_users.Account, error) {
	var rec Account
	err := ur.DB.Where("phone = ?", phone).First(&rec).Error
	return ToDomainAccount(rec), err
}

// CheckOTP implements domain_users.Repository
func (us UsersRepo) CheckOTP(phone string) (bool, error) {
	var rec UserVerif
	err := us.DB.Where("phone = ?", phone).First(&rec).Error
	if err != nil {
		return false, err
	}
	return true, err
}

// StoreVerif implements domain_users.Repository
func (us UsersRepo) StoreVerif(domain domain_users.UserVerif) (string, error) {
	err := us.DB.Save(&domain).Error
	return domain.Phone, err
}
