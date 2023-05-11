package leaguesservices

// import (
// 	"fmt"
// 	"testing"

// 	"server/entities"
// 	mock "server/services/leagues/mock"

// 	"github.com/stretchr/testify/assert"
// )

// func TestListLeagues(t *testing.T) {
// 	mockRepoLeague := new(mock.MockLeaguesRepository)
// 	mockRepoTags := new(mock.MockTagRepository)
// 	service := NewleaguesService(mockRepoLeague, mockRepoTags)
// 	assert := assert.New(t)

// 	want := entities.Leagues {
// 		Leagues: []entities.League{
// 			{LeagueName: "League 1", Active: true},
// 			{LeagueName: "League 2", Active: true},
// 			{LeagueName: "League 3", Active: false},
// 		},
// 	}

// 	mockRepoLeague.On("GetLeagues").Return(want)

// 	got := service.ListLeagues()

// 	assert.Equal(want, got)
// }

// func TestGetLeaguesNameActive(t *testing.T) {
// 	mockRepoLeague := new(mock.MockLeaguesRepository)
// 	mockRepoTags := new(mock.MockTagRepository)
// 	service := NewleaguesService(mockRepoLeague, mockRepoTags)
// 	assert := assert.New(t)

// 	want:= []string{"League 1", "League 2"}

// 	mockRepoLeague.On("GetLeaguesNameActive").Return(want)

// 	got := service.GetLeaguesNameActive()

// 	assert.Equal(want, got)
// }

// func TestChangeStatusSuccess(t *testing.T) {
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
// 	mockRepoLeague.On("WriteLeaguesJSON", leagues).Return(nil)
// 	mockRepoTags.On("DeleteTag", "League 1").Return(nil)
// 	mockRepoTags.On("AddTag", "League 1").Return(nil)
// 	_, err := service.ChangeStatus("League 1")

// 	assert.Nil(err, "no error")

// }

// func TestChangeStatusCantWriteJSON(t *testing.T) {
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
// 	mockRepoLeague.On("WriteLeaguesJSON", leagues).Return(fmt.Errorf("cant write"))
// 	mockRepoTags.On("DeleteTag", "League 1").Return(nil)
// 	mockRepoTags.On("AddTag", "League 1").Return(nil)
// 	_, err := service.ChangeStatus("League 1")

// 	assert.Error(err, "Expected an error for case can not write json")

// }

// func TestChangeStatusCantDeleteTag(t *testing.T) {
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
// 	mockRepoLeague.On("WriteLeaguesJSON", leagues).Return(nil)
// 	mockRepoTags.On("DeleteTag", "League 1").Return(fmt.Errorf("cant delete tag"))
// 	mockRepoTags.On("AddTag", "League 1").Return(nil)
// 	_, err := service.ChangeStatus("League 1")

// 	assert.Error(err, "Expected an error for case can not delete tag")

// }

// func TestChangeStatusCantAddTag(t *testing.T) {
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
// 	mockRepoLeague.On("WriteLeaguesJSON", leagues).Return(nil)
// 	mockRepoTags.On("DeleteTag", "League 3").Return(nil)
// 	mockRepoTags.On("AddTag", "League 3").Return(fmt.Errorf("cant add tag"))
// 	_, err := service.ChangeStatus("League 3")

// 	assert.Error(err, "Expected an error for case can not add tag")

// }

// func TestChangeStatusLeagueNotFound(t *testing.T) {
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
// 	mockRepoLeague.On("WriteLeaguesJSON", leagues).Return(nil)
// 	mockRepoTags.On("DeleteTag", "League 1").Return(nil)
// 	mockRepoTags.On("AddTag", "League 1").Return(nil)
// 	_, err := service.ChangeStatus("League 4")

// 	assert.Error(err, "Expected an error for case league not found")

// }