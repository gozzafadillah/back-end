package domain_users_test

import (
	"errors"
	"os"
	"ppob/app/middlewares"
	domain_users "ppob/users/domain"
	usersMocks "ppob/users/domain/mocks"
	service_users "ppob/users/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userService domain_users.Service
	userDomain  domain_users.Users
	userRepo    usersMocks.Repository
)

func TestMain(m *testing.M) {
	userService = service_users.NewUsersService(&userRepo, &middlewares.ConfigJwt{})
	userDomain = domain_users.Users{
		ID:     1,
		Name:   "Muhammad Fadillah Abdul Aziz",
		Status: true,
		Role:   "customer",
	}
	os.Exit(m.Run())
}

func TestLogin(t *testing.T) {
	t.Run("login success", func(t *testing.T) {
		userRepo.On("CheckEmailPassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(userDomain, nil).Once()
		userRepo.On("GenerateToken", mock.AnythingOfType("int"), mock.AnythingOfType("bool")).Return("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwic3RhdHVzIjp0cnVlfQ.vcnqONfIFjRBD5O4a8LDZXNt2afy4rV2NjmpNBDiAqE", nil).Once()
		res, err := userService.Login("Aziz", "12345")

		assert.NoError(t, err)
		assert.Equal(t, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwic3RhdHVzIjp0cnVlfQ.vcnqONfIFjRBD5O4a8LDZXNt2afy4rV2NjmpNBDiAqE", res)
	})
	t.Run("login Failed Bad Request", func(t *testing.T) {
		userRepo.On("CheckEmailPassword", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(domain_users.Users{}, errors.New("data not found")).Once()
		userRepo.On("GenerateToken", mock.AnythingOfType("int"), mock.AnythingOfType("bool")).Return("", errors.New("failed to generate token")).Once()
		res, err := userService.Login("", "12345")

		assert.Error(t, err)
		assert.Equal(t, "", res)
	})
}

func TestRegister(t *testing.T) {
	t.Run("Register Success", func(t *testing.T) {
		userRepo.On("Store", mock.Anything).Return(1, nil).Once()
		userRepo.On("GetById", mock.Anything).Return(userDomain, nil)
		res, err := userService.Register(userDomain)

		assert.NoError(t, err)
		assert.Equal(t, 1, res.ID)
	})
	t.Run("Register Data Not Found", func(t *testing.T) {
		userRepo.On("Store", mock.Anything).Return(0, errors.New("bad request")).Once()
		userRepo.On("GetById", mock.Anything).Return(userDomain, nil)
		res, err := userService.Register(domain_users.Users{})

		assert.Error(t, err)
		assert.Equal(t, domain_users.Users{}, res)
	})
}
