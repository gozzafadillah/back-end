package service_users

import (
	"errors"
	"fmt"
	"ppob/app/middlewares"
	domain_users "ppob/users/domain"
)

type UsersService struct {
	Repository domain_users.Repository
	jwtauth    *middlewares.ConfigJwt
}

// Register implements domain_users.Service
func (us UsersService) Register(domain domain_users.Users) (domain_users.Users, error) {
	id, err := us.Repository.Store(domain)
	if err != nil {
		return domain_users.Users{}, errors.New("gagal di tambahkan")
	}
	data, err := us.Repository.GetById(id)
	if err != nil {
		return domain_users.Users{}, errors.New("gagal di tambahkan")
	}
	return data, nil
}

func NewUsersService(repo domain_users.Repository, jwt *middlewares.ConfigJwt) domain_users.Service {
	return UsersService{
		Repository: repo,
		jwtauth:    jwt,
	}
}

// Login implements domain_users.Service
func (us UsersService) Login(email string, password string) (string, error) {
	data, err := us.Repository.CheckEmailPassword(email, password)
	fmt.Println("data login : ", data)
	if err != nil {
		return "", errors.New("failed to generate token, user not found")
	}
	token, err := us.jwtauth.GenerateToken(data.ID, data.Status)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
