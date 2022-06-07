package domain_users

import (
	"time"
)

type Users struct {
	ID        int
	Name      string
	Slug      string
	DOB       time.Time
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
	ID    int
	Phone string
	Saldo int
	Pin   string
}

type UserVerif struct {
	ID     int
	Code   string
	Status bool
}
