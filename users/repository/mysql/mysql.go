package mysql_users

import (
	"errors"
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

// GetById implements domain_users.Repository
func (ur UsersRepo) GetById(id int) (domain_users.Users, error) {
	rec := Users{}
	err := ur.DB.Where("id=?", id).First(&rec).Error
	return toDomain(rec), err
}

// Store implements domain_users.Repository
func (ur UsersRepo) Store(domain domain_users.Users) (int, error) {
	err := ur.DB.Save(&domain).Error
	return domain.ID, err
}

// CheckEmailPassword implements domain_users.Repository
func (ur UsersRepo) CheckEmailPassword(email string, password string) (domain_users.Users, error) {
	var rec Users
	err := ur.DB.Where("email = ?", email).First(&rec).Error
	if err != nil {
		return domain_users.Users{}, errors.New("data not found")
	}

	data := encryption.CheckPasswordHash(password, rec.Password)
	if !data && !rec.Status {
		return domain_users.Users{}, err
	}

	return toDomain(rec), nil
}
