package adminrepo_test

import (
	"fmt"
	"server/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAdmin(t *testing.T) {
	assert := assert.New(t)
	mockAdminRepo := new(MockAdminRepository)
	want := entities.Admin{
		Username: "admin2023",
		Password: "password_encrypted",
	}
	mockAdminRepo.On("GetAdmin").Return(want)

	got := mockAdminRepo.GetAdmin()

	assert.Equal(want, got, fmt.Sprintf("Method GetAdmin is supose to %v, but got %s", want, got))
}

func TestReadAdminJSON(t *testing.T) {
	assert := assert.New(t)
	mockAdminRepo := new(MockAdminRepository)
	want := entities.Admin{
		Username: "admin2023",
		Password: "password_encrypted",
	}

	mockAdminRepo.On("ReadAdminJSON").Return(want, nil)
	got, _ := mockAdminRepo.ReadAdminJSON()

	assert.Equal(want, got, fmt.Sprintf("Method ReadAdminJSON is supose to %#v, but got %#v", want, got))
}

func TestWriteAdminJson(t *testing.T) {
	assert := assert.New(t)
	mockAdminRepo := new(MockAdminRepository)
	input := &entities.Admin{
		Username: "admin2023",
		Password: "password_encrypted",
	}

	mockAdminRepo.On("WriteAdminJSON", input).Return(nil)
	got := mockAdminRepo.WriteAdminJSON(input)

	assert.Nil( got, fmt.Sprintf("Method ReadAdminJSON is supose to nil, but got %#v", got))
}