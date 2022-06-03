package mysql_users

import (
	domain_users "ppob/users/domain"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID         int
	Name       string
	Slug       string
	Email      string
	Password   string
	Phone      string
	Status     bool
	Account_id int
	Role       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func toDomain(rec Users) domain_users.Users {
	return domain_users.Users{
		ID:         rec.ID,
		Name:       rec.Name,
		Slug:       rec.Slug,
		Email:      rec.Email,
		Password:   rec.Password,
		Phone:      rec.Phone,
		Status:     rec.Status,
		Account_id: rec.Account_id,
		Role:       rec.Role,
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
	}
}
