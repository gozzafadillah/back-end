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

// GetByEmail implements domain_users.Repository
func (ur UsersRepo) GetByEmail(email string) (domain_users.Users, error) {
	var rec Users
	err := ur.DB.Where("email = ?", email).First(&rec).Error
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
		return domain_users.Users{}, errors.New("email or password miss macth")
	}

	data := encryption.CheckPasswordHash(password, rec.Password)
	fmt.Println(data)
	if !data {
		return domain_users.Users{}, errors.New("email or password miss macth")
	}
	if !rec.Status {
		return domain_users.Users{}, errors.New("unauthorized account, please contact customer service")
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

// StoreOtpUserVerif implements domain_users.Repository
func (ur UsersRepo) StoreOtpUserVerif(code string, email string) error {
	rec := UserVerif{
		Email: email,
		Code:  code,
	}
	err := ur.DB.Where("email = ?", email).Save(&rec).Error
	return err
}

// Verif implements domain_users.Repository
func (ur UsersRepo) Verif(code string) (domain_users.UserVerif, error) {
	rec := UserVerif{}
	err := ur.DB.Where("code = ?", code).First(&rec).Error
	return ToDomainVerif(rec), err
}

// ChangeStatusUsers implements domain_users.Repository
func (ur UsersRepo) ChangeStatusUsers(email string) error {
	rec := Users{}
	err := ur.DB.Model(&rec).Where("email = ?", email).Update("status", true).Error
	return err
}

// ChangeStatus implements domain_users.Repository
func (ur UsersRepo) ChangeStatusVerif(email string) error {
	rec := UserVerif{}
	err := ur.DB.Model(&rec).Where("email = ?", email).Update("status", true).Error
	return err
}
