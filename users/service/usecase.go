package service_users

import (
	"errors"
	"ppob/app/middlewares"
	domain_users "ppob/users/domain"
)

type UsersService struct {
	Repository domain_users.Repository
	jwtauth    *middlewares.ConfigJwt
}

func NewUsersService(repo domain_users.Repository, jwt *middlewares.ConfigJwt) domain_users.Service {
	return UsersService{
		Repository: repo,
		jwtauth:    jwt,
	}
}

// Register implements domain_users.Service
func (us UsersService) Register(domain domain_users.Users) (domain_users.Users, error) {
	id, err := us.Repository.Store(domain)
	if err != nil {
		return domain_users.Users{}, errors.New("gagal ditambahkan")
	}
	data, err := us.Repository.GetById(id)
	if err != nil {
		return domain_users.Users{}, errors.New("gagal ditambahkan")
	}
	return data, nil
}

// Login implements domain_users.Service
func (us UsersService) Login(email string, password string) (string, error) {
	data, err := us.Repository.CheckEmailPassword(email, password)
	if err != nil {
		return "", errors.New("data not found")
	}
	token, err := us.jwtauth.GenerateToken(data.ID, data.Status)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
