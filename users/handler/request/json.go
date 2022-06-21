package request

import (
	domain_users "ppob/users/domain"
)

type RequestJSONUser struct {
	Name     string      `json:"name" form:"name" validate:"required"`
	Email    string      `json:"email" form:"email" validate:"required,email"`
	Password string      `json:"password" form:"password" validate:"required"`
	Phone    string      `json:"phone" form:"phone" validate:"required"`
	Image    string      `json:"img" form:"img"`
	File     interface{} `json:"file,omitempty"`
}

type RequestJSONAccount struct {
	Phone string
	Saldo int
	Pin   string `json:"pin" form:"pin"`
}

type RequestJSONLogin struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

type RequestJSONVerif struct {
	Code string `json:"code" validate:"required"`
}

type RequestJSONRefresh struct {
	Email  string `json:"email" validate:"email"`
	Status bool
}

func ToDomainUser(req RequestJSONUser) domain_users.Users {
	return domain_users.Users{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Phone:    req.Phone,
		Image:    req.Image,
		Status:   false,
		Role:     "customer",
	}
}
func ToDomainLogin(req RequestJSONLogin) domain_users.Users {
	return domain_users.Users{
		Email:    req.Email,
		Password: req.Password,
	}
}

func ToDomainAccount(req RequestJSONAccount) domain_users.Account {
	return domain_users.Account{
		Phone: req.Phone,
		Saldo: 0,
		Pin:   req.Pin,
	}
}

func ToDomainVerif(req RequestJSONVerif) domain_users.UserVerif {
	return domain_users.UserVerif{
		Code:   req.Code,
		Status: false,
	}
}

func ToDomainReVerif(req RequestJSONRefresh) domain_users.UserVerif {
	return domain_users.UserVerif{
		Email:  req.Email,
		Status: req.Status,
	}
}
