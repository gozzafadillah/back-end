package request

import domain_users "ppob/users/domain"

type RequestJSON struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
}

type RequestJSONLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func ToDomain(req RequestJSON) domain_users.Users {
	return domain_users.Users{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Phone:    req.Phone,
		Status:   true,
	}
}
func ToDomainLogin(req RequestJSONLogin) domain_users.Users {
	return domain_users.Users{
		Email:    req.Email,
		Password: req.Password,
	}
}
