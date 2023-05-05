package adminservices

import (
	"testing"
	"fmt"
	"server/entities"
	mock "server/services/admin/mock"

	"github.com/stretchr/testify/assert"
)

func TestCheckAdminUsernameTokenSuccess(t *testing.T) {
	mockRepoAdmin := new(mock.MockAdminRepository)
	service := NewAdminService(mockRepoAdmin)
	assert := assert.New(t)

	admin := entities.Admin{
		Username: "username",
		Password: "password",
	}
	mockRepoAdmin.On("GetAdmin").Return(admin)

	got := service.CheckAdminUsernameToken("username")

	assert.Nil(got)
}

func TestCheckAdminUsernameTokenFail(t *testing.T) {
	mockRepoAdmin := new(mock.MockAdminRepository)
	service := NewAdminService(mockRepoAdmin)
	assert := assert.New(t)

	admin := entities.Admin{
		Username: "username",
		Password: "password",
	}
	mockRepoAdmin.On("GetAdmin").Return(admin)

	got := service.CheckAdminUsernameToken("usernamefail")

	assert.Error(got, "Expected an error for case check admin fail")
}

func TestGetAdminUsernameSuccess(t *testing.T) {
	mockRepoAdmin := new(mock.MockAdminRepository)
	service := NewAdminService(mockRepoAdmin)
	assert := assert.New(t)

	admin := entities.Admin{
		Username: "username",
		Password: "password",
	}
	mockRepoAdmin.On("GetAdmin").Return(admin)
	want := admin.Username
	got, err := service.GetAdminUsername("username")
	assert.Nil(err)
	assert.Equal(want, got)
}

func TestGetAdminUsernameFail(t *testing.T) {
	mockRepoAdmin := new(mock.MockAdminRepository)
	service := NewAdminService(mockRepoAdmin)
	assert := assert.New(t)

	admin := entities.Admin{
		Username: "username",
		Password: "password",
	}
	mockRepoAdmin.On("GetAdmin").Return(admin)

	_, err := service.GetAdminUsername("usernamefail")
	assert.Error(err, "Expected an error for case GetAdminUsername fail")
}

func TestChangePasswordGetAdminFail(t *testing.T) {
	mockRepoAdmin := new(mock.MockAdminRepository)
	service := NewAdminService(mockRepoAdmin)
	assert := assert.New(t)

	admin := entities.Admin{
		Username: "username",
		Password: "password",
	}
	adminWithConfirmPassword := AdminWithConfirmPassword{
		Username: "usernamefail",
		Password: "password",
		PasswordConfirmation: "password",
	}
	usernameFromToken := "username"
	mockRepoAdmin.On("GetAdmin").Return(admin)

  err := service.ChangePassword(&adminWithConfirmPassword, usernameFromToken)
	assert.Error(err, "Expected an error for case ChangePassword fail")
}

func TestChangePasswordValidateAdminFail(t *testing.T) {
	mockRepoAdmin := new(mock.MockAdminRepository)
	service := NewAdminService(mockRepoAdmin)
	assert := assert.New(t)

	admin := entities.Admin{
		Username: "username@@@toolong",
		Password: "password",
	}
	adminWithConfirmPassword := AdminWithConfirmPassword{
		Username: "username@@@toolong",
		Password: "password",
		PasswordConfirmation: "password",
	}
	usernameFromToken := "username"
	mockRepoAdmin.On("GetAdmin").Return(admin)

  err := service.ChangePassword(&adminWithConfirmPassword, usernameFromToken)
	assert.Error(err, "Expected an error for case ChangePassword fail")
}

func TestChangePasswordNotMatchPasswordConfirmation(t *testing.T) {
	mockRepoAdmin := new(mock.MockAdminRepository)
	service := NewAdminService(mockRepoAdmin)
	assert := assert.New(t)

	admin := entities.Admin{
		Username: "username",
		Password: "password",
	}
	adminWithConfirmPassword := AdminWithConfirmPassword{
		Username: "username",
		Password: "password",
		PasswordConfirmation: "passwordheelo",
	}
	usernameFromToken := "username"
	mockRepoAdmin.On("GetAdmin").Return(admin)

  err := service.ChangePassword(&adminWithConfirmPassword, usernameFromToken)
	assert.Error(err, "Expected an error for case ChangePassword fail")
}

func TestChangePasswordWriteAdminJsonFail(t *testing.T) {
	mockRepoAdmin := new(mock.MockAdminRepository)
	service := NewAdminService(mockRepoAdmin)
	assert := assert.New(t)

	admin := entities.Admin{
		Username: "username",
		Password: "password",
	}
	adminWithConfirmPassword := AdminWithConfirmPassword{
		Username: "username",
		Password: "password",
		PasswordConfirmation: "password",
	}
	usernameFromToken := "username"
	mockRepoAdmin.On("GetAdmin").Return(admin)
	mockRepoAdmin.On("WriteAdminJSON", &admin).Return(fmt.Errorf("can not write admin"))

  err := service.ChangePassword(&adminWithConfirmPassword, usernameFromToken)
	assert.Error(err, "Expected an error for case ChangePassword fail")
}

func TestChangePasswordReadAdminJsonFail(t *testing.T) {
	mockRepoAdmin := new(mock.MockAdminRepository)
	service := NewAdminService(mockRepoAdmin)
	assert := assert.New(t)

	admin := entities.Admin{
		Username: "username",
		Password: "password",
	}
	adminWithConfirmPassword := AdminWithConfirmPassword{
		Username: "username",
		Password: "password",
		PasswordConfirmation: "password",
	}
	usernameFromToken := "username"
	mockRepoAdmin.On("GetAdmin").Return(admin)
	mockRepoAdmin.On("WriteAdminJSON", &admin).Return(nil)
	mockRepoAdmin.On("ReadAdminJSON").Return(admin,fmt.Errorf("cant read admin json"))

  err := service.ChangePassword(&adminWithConfirmPassword, usernameFromToken)
	assert.Error(err, "Expected an error for case ChangePassword fail")
}