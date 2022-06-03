package domain_users

import "time"

type Users struct {
	ID         int
	Name       string
	Slug       string
	Email      string
	Password   string
	Phone      string
	Image      string
	Status     bool
	Account_id int
	Role       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
