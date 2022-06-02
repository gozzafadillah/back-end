package request

import domain_users "ppob/users/domain"

type RequestJSON struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
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
