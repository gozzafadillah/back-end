package mysql_users

import (
	domain_users "ppob/users/domain"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID        int
	Name      string
	DOB       time.Time
	Slug      string
	Email     string
	Password  string
	Phone     string
	Image     string
	Status    bool
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Account struct {
	gorm.Model
	ID    int
	Phone string
	Saldo int
	Pin   string
}

func toDomain(rec Users) domain_users.Users {
	return domain_users.Users{
		ID:        rec.ID,
		Name:      rec.Name,
		Slug:      rec.Slug,
		DOB:       rec.DOB,
		Email:     rec.Email,
		Password:  rec.Password,
		Phone:     rec.Phone,
		Image:     rec.Image,
		Status:    rec.Status,
		Role:      rec.Role,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}

func ToDomainAccount(rec Account) domain_users.Account {
	return domain_users.Account{
		ID:    rec.ID,
		Phone: rec.Phone,
		Saldo: rec.Saldo,
		Pin:   rec.Pin,
	}
}
