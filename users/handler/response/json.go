package response

import (
	domain_users "ppob/users/domain"
	"time"
)

type ResponseJSONUsers struct {
	Name   string
	DOB    time.Time
	Email  string
	Phone  string
	Image  string
	Status bool
}

type ResponseJSONAccount struct {
	ID    int
	Phone string
	Saldo int
}

type ResponseJSONVerif struct {
	Phone  string
	Code   string
	Status bool
}

func FromDomainUsers(domain domain_users.Users) ResponseJSONUsers {
	return ResponseJSONUsers{
		Name:   domain.Name,
		DOB:    domain.DOB,
		Email:  domain.Email,
		Phone:  domain.Phone,
		Image:  domain.Image,
		Status: domain.Status,
	}
}
func FromDomainAccount(domain domain_users.Account) ResponseJSONAccount {
	return ResponseJSONAccount{
		ID:    domain.ID,
		Phone: domain.Phone,
		Saldo: domain.Saldo,
	}
}
func FromDomainVerif(domain domain_users.UserVerif) ResponseJSONVerif {
	return ResponseJSONVerif{
		Phone:  domain.Phone,
		Code:   domain.Code,
		Status: domain.Status,
	}
}
