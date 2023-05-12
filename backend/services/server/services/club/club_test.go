package clubservices

import (
	"server/entities"
	mock "server/services/club/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOrCreate(t *testing.T) {
	mockRepoClub := new(mock.MockClubRepository)
	service := NewClubService(mockRepoClub)
	assert := assert.New(t)

	club := &entities.Club{
		Name: "test",
		Logo: "test",
	}
	mockRepoClub.On("FirstOrCreate", club.Name, club.Logo).Return(club, nil)

	got, err := service.GetOrCreate(club.Name, club.Logo)
	assert.Nil(err)

	assert.Equal(club, got)
}

func TestGetClubByName(t *testing.T) {
	mockRepoClub := new(mock.MockClubRepository)
	service := NewClubService(mockRepoClub)
	assert := assert.New(t)

	club := &entities.Club{
		Name: "test",
		Logo: "test",
	}
	mockRepoClub.On("GetByName", club.Name).Return(club, nil)

	got, err := service.GetClubByName(club.Name)
	assert.Nil(err)

	assert.Equal(club, got)
}
