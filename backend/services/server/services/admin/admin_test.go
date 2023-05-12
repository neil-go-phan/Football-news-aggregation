package adminservices

import (
	"fmt"
	"server/entities"
	mock "server/services/admin/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckAdminUsername_Token_Success(t *testing.T) {
	mockRepoAdmin := new(mock.MockAdminRepository)
	service := NewAdminService(mockRepoAdmin)
	assert := assert.New(t)

	admin := &entities.Admin{
		Username: "username",
		Password: "password",
	}
	mockRepoAdmin.On("Get", admin.Username).Return(admin, nil)

	got := service.CheckAdminUsernameToken("username")

	assert.Nil(got)
}

func TestCheckAdminUsername_Token_Fail(t *testing.T) {
	mockRepoAdmin := new(mock.MockAdminRepository)
	service := NewAdminService(mockRepoAdmin)
	assert := assert.New(t)

	admin := &entities.Admin{
		Username: "username",
		Password: "password",
	}
	mockRepoAdmin.On("Get", "usernamefail").Return(admin, fmt.Errorf("fail"))

	got := service.CheckAdminUsernameToken("usernamefail")

	assert.Error(got, "Expected an error for case check admin fail")
}

func TestGetAdminUsername_Success(t *testing.T) {
	mockRepoAdmin := new(mock.MockAdminRepository)
	service := NewAdminService(mockRepoAdmin)
	assert := assert.New(t)

	admin := &entities.Admin{
		Username: "username",
		Password: "password",
	}
	mockRepoAdmin.On("Get", admin.Username).Return(admin, nil)
	want := admin.Username
	got, err := service.GetAdminUsername("username")
	assert.Nil(err)
	assert.Equal(want, got)
}

func TestGetAdminUsername_Fail(t *testing.T) {
	mockRepoAdmin := new(mock.MockAdminRepository)
	service := NewAdminService(mockRepoAdmin)
	assert := assert.New(t)

	admin := &entities.Admin{
		Username: "username",
		Password: "password",
	}
	mockRepoAdmin.On("Get", "usernamefail").Return(admin, fmt.Errorf("fail"))

	_, err := service.GetAdminUsername("usernamefail")
	assert.Error(err, "Expected an error for case GetAdminUsername fail")
}

func TestChangePasswordGetAdmin_Fail(t *testing.T) {
	mockRepoAdmin := new(mock.MockAdminRepository)
	service := NewAdminService(mockRepoAdmin)
	assert := assert.New(t)

	admin := &entities.Admin{
		Username: "username",
		Password: "password",
	}
	adminWithConfirmPassword := AdminWithConfirmPassword{
		Username:             "usernamefail",
		Password:             "password",
		PasswordConfirmation: "password",
	}
	usernameFromToken := "username"
	mockRepoAdmin.On("Get", "usernamefail").Return(admin, nil)

	err := service.ChangePassword(&adminWithConfirmPassword, usernameFromToken)
	assert.Error(err, "Expected an error for case ChangePassword fail")
}

func TestChangePasswordValidateAdmin_Fail(t *testing.T) {
	mockRepoAdmin := new(mock.MockAdminRepository)
	service := NewAdminService(mockRepoAdmin)
	assert := assert.New(t)

	admin := &entities.Admin{
		Username: "username@@@toolong",
		Password: "password",
	}
	adminWithConfirmPassword := AdminWithConfirmPassword{
		Username:             "username@@@toolong",
		Password:             "password",
		PasswordConfirmation: "password",
	}
	usernameFromToken := "username"
	mockRepoAdmin.On("Get", "username@@@toolong").Return(admin, nil)

	err := service.ChangePassword(&adminWithConfirmPassword, usernameFromToken)
	assert.Error(err, "Expected an error for case ChangePassword fail")
}

func TestChangePasswordNotMatchPasswordConfirmation(t *testing.T) {
	mockRepoAdmin := new(mock.MockAdminRepository)
	service := NewAdminService(mockRepoAdmin)
	assert := assert.New(t)

	admin := &entities.Admin{
		Username: "username",
		Password: "password",
	}
	adminWithConfirmPassword := AdminWithConfirmPassword{
		Username:             "username",
		Password:             "password",
		PasswordConfirmation: "passwordheelo",
	}
	usernameFromToken := "username"
	mockRepoAdmin.On("Get", "username").Return(admin, nil)

	err := service.ChangePassword(&adminWithConfirmPassword, usernameFromToken)
	assert.Error(err, "Expected an error for case ChangePassword fail")
}
