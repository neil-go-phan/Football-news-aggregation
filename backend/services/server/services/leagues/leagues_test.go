package leaguesservices

import (
	"testing"

	"server/entities"
	mock "server/services/leagues/mock"

	"github.com/stretchr/testify/assert"
)

func TestListLeagues(t *testing.T) {
	mockRepoLeague := new(mock.MockLeaguesRepository)
	mockRepoTags := new(mock.MockTagRepository)
	service := NewleaguesService(mockRepoLeague, mockRepoTags)
	assert := assert.New(t)

	want := entities.Leagues {
		Leagues: []entities.League{
			{LeagueName: "League 1", Active: true},
			{LeagueName: "League 2", Active: true},
			{LeagueName: "League 3", Active: false},
		},
	}

	mockRepoLeague.On("GetLeagues").Return(want)

	got := service.ListLeagues()

	assert.Equal(want, got)
}

func TestGetLeaguesNameActive(t *testing.T) {
	mockRepoLeague := new(mock.MockLeaguesRepository)
	mockRepoTags := new(mock.MockTagRepository)
	service := NewleaguesService(mockRepoLeague, mockRepoTags)
	assert := assert.New(t)

	want:= []string{"League 1", "League 2"}

	mockRepoLeague.On("GetLeaguesNameActive").Return(want)

	got := service.GetLeaguesNameActive()

	assert.Equal(want, got)
}

// func TestChangeStatus(t *testing.T) {
// 	mockRepoLeague := new(mock.MockLeaguesRepository)
// 	mockRepoTags := new(mock.MockTagRepository)
// 	service := NewleaguesService(mockRepoLeague, mockRepoTags)
// 	assert := assert.New(t)

// 	leagues := entities.Leagues {
// 		Leagues: []entities.League{
// 			{LeagueName: "League 1", Active: true},
// 			{LeagueName: "League 2", Active: true},
// 			{LeagueName: "League 3", Active: false},
// 		},
// 	}

// 	mockRepoLeague.On("GetLeagues").Return(leagues)

// }