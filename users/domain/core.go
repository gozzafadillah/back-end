package domain_users

import "time"

type Users struct {
	ID         int
	Name       string
	Slug       string
	Email      string
	Password   string
	Phone      string
	Status     bool
	Address_Id int
	Role       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
