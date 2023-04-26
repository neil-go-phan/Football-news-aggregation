package adminrepo

import (
	"fmt"
	"server/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

var PATH = "./testJson/"

func TestNewAdminRepo(t *testing.T) {
	assert := assert.New(t)
	contructorAdmin := entities.Admin{
		Username: "admin2023",
		Password: "password_encrypted",
	}
	want := &adminRepo{
		admin: entities.Admin{
			Username: "admin2023",
			Password: "password_encrypted",
		},
		path: PATH,
	}

	got := NewAdminRepo(contructorAdmin, PATH)

	assert.Equal(want, got, fmt.Sprintf("Method NewAdminRepo is supose to %v, but got %s", want, got))
}

func TestGetAdmin(t *testing.T) {
	assert := assert.New(t)
	contructorAdmin := entities.Admin{
		Username: "admin2023",
		Password: "password_encrypted",
	}
	want := entities.Admin{
		Username: "admin2023",
		Password: "password_encrypted",
	}

	adminRepo := NewAdminRepo(contructorAdmin, PATH)

	got := adminRepo.GetAdmin()

	assert.Equal(want, got, fmt.Sprintf("Method GetAdmin is supose to %v, but got %s", want, got))
}

func TestReadAdminJSONSuccess(t *testing.T) {
	assert := assert.New(t)
	contructorAdmin := entities.Admin{
		Username: "admin2023test",
		Password: "password_encryptedtest",
	}
	want := entities.Admin{
		Username: "admin2023",
		Password: "password_encrypted",
	}

	adminRepo := NewAdminRepo(contructorAdmin, PATH)

	got, _ := adminRepo.ReadAdminJSON()

	assert.Equal(want, got, fmt.Sprintf("Method ReadAdminJSON is supose to %#v, but got %#v", want, got))
}

func TestReadAdminJSONCantOpenFile(t *testing.T) {
	assert := assert.New(t)
	contructorAdmin := entities.Admin{
		Username: "admin2023test",
		Password: "password_encryptedtest",
	}
	want := "json file not found"

	adminRepo := NewAdminRepo(contructorAdmin, "./testJson/fail/")

	_, got := adminRepo.ReadAdminJSON()

	assert.Errorf(got, want, fmt.Sprintf("Method ReadAdminJSON is supose to %#v, but got %#v", want, got))
}

func TestWriteAdminJsonSuccess(t *testing.T) {
	assert := assert.New(t)
	contructorAdmin := entities.Admin{
		Username: "admin2023test",
		Password: "password_encryptedtest",
	}

	writedAdmin := &entities.Admin{
		Username: "admin2023",
		Password: "password_encrypted",
	}

	adminRepo := NewAdminRepo(contructorAdmin, PATH)
	got := adminRepo.WriteAdminJSON(writedAdmin)

	assert.Nil(got, fmt.Sprintf("Method WriteAdminJson is supose to nil, but got %#v", got))
}

func TestWriteAdminJsonCantOpenFile(t *testing.T) {
	assert := assert.New(t)
	contructorAdmin := entities.Admin{
		Username: "admin2023test",
		Password: "password_encryptedtest",
	}

	writedAdmin := &entities.Admin{
		Username: "admin2023",
		Password: "password_encrypted",
	}

	want := "json file not found"

	adminRepo := NewAdminRepo(contructorAdmin, "./testJson/fail/")

	got := adminRepo.WriteAdminJSON(writedAdmin)

	assert.Errorf(got, want, fmt.Sprintf("Method WriteAdminJson is supose to %#v, but got %#v", want, got))
}

func TestSetAdmin(t *testing.T) {
	assert := assert.New(t)
	contructorAdmin := entities.Admin{
		Username: "admin2023test",
		Password: "password_encryptedtest",
	}

	want := entities.Admin{
		Username: "admin2023",
		Password: "password_encrypted",
	}

	adminRepo := NewAdminRepo(contructorAdmin, PATH)
	adminRepo.SetAdmin(want)

	got,_ := adminRepo.ReadAdminJSON()

	assert.Equal(want, got, fmt.Sprintf("Method ReadAdminJSON is supose to %#v, but got %#v", want, got))
}