package leaguesservices

import (
	"testing"

	"server/entities"
	mock "server/services/leagues/mock"

	"github.com/stretchr/testify/assert"
)

func TestListLeagues(t *testing.T) {
	mockRepoLeague := new(mock.MockLeaguesRepository)
	mockServiceTags := new(mock.MockTagsServices)
	service := NewleaguesService(mockRepoLeague, mockServiceTags)
	assert := assert.New(t)

	want := &[]entities.League{
		{LeagueName: "League 1", Active: true},
		{LeagueName: "League 2", Active: true},
		{LeagueName: "League 3", Active: false},
	}

	mockRepoLeague.On("List").Return(want, nil)

	got, err := service.ListLeagues()
	assert.Nil(err)
	assert.Equal(want, got)
}

func TestGetLeaguesNameActive(t *testing.T) {
	mockRepoLeague := new(mock.MockLeaguesRepository)
	mockServiceTags := new(mock.MockTagsServices)
	service := NewleaguesService(mockRepoLeague, mockServiceTags)
	assert := assert.New(t)

	want := []string{"League 1", "League 2"}
	leagues := &[]entities.League{
		{LeagueName: "League 1", Active: true},
		{LeagueName: "League 2", Active: true},
	}
	mockRepoLeague.On("GetLeaguesNameActive").Return(leagues, nil)

	got, err := service.GetLeaguesNameActive()

	assert.Nil(err)
	assert.Equal(want, got)
}

func TestCreateLeague(t *testing.T) {
	mockRepoLeague := new(mock.MockLeaguesRepository)
	mockServiceTags := new(mock.MockTagsServices)
	service := NewleaguesService(mockRepoLeague, mockServiceTags)
	assert := assert.New(t)

	leagues := &entities.League{
		LeagueName: "League 1", Active: false,
	}

	mockRepoLeague.On("Create", leagues).Return(nil)

	err := service.CreateLeague("League 1")

	assert.Nil(err, "no error")
}

func TestGetLeaguesName(t *testing.T) {
	mockRepoLeague := new(mock.MockLeaguesRepository)
	mockServiceTags := new(mock.MockTagsServices)
	service := NewleaguesService(mockRepoLeague, mockServiceTags)
	assert := assert.New(t)

	want := []string{"League 1", "League 2","League 3"}
	leagues := &[]entities.League{
		{LeagueName: "League 1", Active: true},
		{LeagueName: "League 2", Active: true},
		{LeagueName: "League 3", Active: false},
	}
	mockRepoLeague.On("GetLeaguesName").Return(leagues, nil)

	got, err := service.GetLeaguesName()

	assert.Nil(err)
	assert.Equal(want, got)
}

// func TestChangeStatusSuccess(t *testing.T) {
// 	mockRepoLeague := new(mock.MockLeaguesRepository)
// 	mockRepoTags := new(mock.MockTagRepository)
// 	service := NewleaguesService(mockRepoLeague, mockRepoTags)
// 	assert := assert.New(t)

// 	leagues := entities.Leagues{
// 		Leagues: []entities.League{
// 			{LeagueName: "League 1", Active: true},
// 			{LeagueName: "League 2", Active: true},
// 			{LeagueName: "League 3", Active: false},
// 		},
// 	}

// 	mockRepoLeague.On("GetLeagues").Return(leagues)
// 	mockRepoLeague.On("WriteLeaguesJSON", leagues).Return(nil)
// 	mockRepoTags.On("DeleteTag", "League 1").Return(nil)
// 	mockRepoTags.On("AddTag", "League 1").Return(nil)
// 	_, err := service.ChangeStatus("League 1")

// 	assert.Nil(err, "no error")

// }
