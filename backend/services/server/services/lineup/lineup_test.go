package lineupservice

import (
	"server/entities"
	mock "server/services/lineup/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGetOrCreate(t *testing.T) {
	mockRepoLineup:= new(mock.MockLineupRepository)
	service := NewLineupService(mockRepoLineup)
	assert := assert.New(t)

	lineup := &entities.MatchLineUp{}
	mockRepoLineup.On("FirstOrCreate", lineup).Return(lineup, nil)

	got, err := service.GetOrCreate(lineup)
	assert.Nil(err)

	assert.Equal(lineup, got)
}

func TestGetLineUps(t *testing.T) {
	mockRepoLineup:= new(mock.MockLineupRepository)
	service := NewLineupService(mockRepoLineup)
	assert := assert.New(t)

	lineup1 := &entities.MatchLineUp{
		Model: gorm.Model{
			ID: uint(1),
		},
	}
	lineup2 := &entities.MatchLineUp{
		Model: gorm.Model{
			ID: uint(2),
		},
	}
	mockRepoLineup.On("Get", uint(1)).Return(lineup1, nil)
	mockRepoLineup.On("Get", uint(2)).Return(lineup2, nil)
	got1,got2,  err := service.GetLineUps(uint(1), uint(2))

	assert.Nil(err)
	assert.Equal(got1, lineup1)
	assert.Equal(got2, lineup2)
}
