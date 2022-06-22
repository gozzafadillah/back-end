package service_users

import (
	"errors"
	"fmt"
	"ppob/app/middlewares"
	"ppob/helper/mailjet"
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

// GetUsers implements domain_users.Service
func (us UsersService) GetUsers() ([]domain_users.Users, error) {
	data, err := us.Repository.GetAllUser()
	if err != nil {
		return []domain_users.Users{}, errors.New("internal server error")
	}
	return data, nil
}

// GetUserPhone implements domain_users.Service
func (us UsersService) GetUserPhone(phone string) (domain_users.Users, error) {
	data, err := us.Repository.GetByPhone(phone)
	if err != nil {
		return domain_users.Users{}, errors.New("bad request")
	}
	return data, nil
}

// EditUser implements domain_users.Service
func (us UsersService) EditUser(phone string, domain domain_users.Users) error {
	err := us.Repository.Update(phone, domain)
	if err != nil {
		return errors.New("bad request")
	}
	return nil
}

// Register implements domain_users.Service
func (us UsersService) Register(domain domain_users.Users) (domain_users.Users, error) {
	phone, err := us.Repository.Store(domain)
	if err != nil {
		return domain_users.Users{}, errors.New("internal server error")
	}
	data, err := us.Repository.GetByPhone(phone)
	if err != nil {
		return domain_users.Users{}, errors.New("bad request")
	}

	return data, nil
}

// Login implements domain_users.Service
func (us UsersService) Login(email string, password string) (string, error) {
	data, err := us.Repository.CheckEmailPassword(email, password)
	if err != nil {
		return "", err
	}
	token, err := us.jwtauth.GenerateToken(data.ID, data.Phone, data.Status)
	if err != nil {
		return "", errors.New("internal server error")
	}

	return token, nil
}

// InsertAccount implements domain_users.Service
func (us UsersService) InsertAccount(domain domain_users.Account) (domain_users.Account, error) {
	data, err := us.Repository.StoreAccount(domain)
	if err != nil {
		return domain_users.Account{}, errors.New("internal server error")
	}
	return data, nil
}

// GetUserAccount implements domain_users.Service
func (us UsersService) GetUserAccount(phone string) domain_users.Account {
	data, err := us.Repository.GetUserAccount(phone)
	if err != nil {
		return domain_users.Account{}
	}
	return data
}

// AddUserVerif implements domain_users.Service
func (us UsersService) AddUserVerif(code, email, name string) error {
	fmt.Println(code, email, name)
	var data = []byte(`{
		"Messages":[
				{
						"From": {
								"Email": "gozza15bdg@gmail.com",
								"Name": "Muhammad Fadillah Abdul Aziz"
						},
						"To": [
								{
										"Email": "` + email + `",
										"Name": "` + name + `"
								}
						],
						"Subject": "Verification OTP",
						"TextPart": "Code Generator",
						"HTMLPart": "<center><h2>OTP Code</h2><br /> <b><u>` + code + `</u></b> </center>"
				}
		]
	}`)
	err := us.Repository.StoreOtpUserVerif(code, email)
	if err != nil {
		return errors.New("internal server error")
	}
	mailjet.Mailjet(data)

	return nil

}

// Verif implements domain_users.Service
func (us UsersService) Verif(code string) error {
	data, err := us.Repository.Verif(code)
	if err != nil {
		return errors.New("bad request")
	}
	err = us.Repository.ChangeStatusVerif(data.Email)
	if err != nil {
		return errors.New("internal server error")
	}

	err = us.Repository.ChangeStatusUsers(data.Email)
	if err != nil {
		return errors.New("internal server error")
	}

	return nil
}
