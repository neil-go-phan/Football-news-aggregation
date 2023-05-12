package statsitem

import (
	"server/entities"
	mock "server/services/statsItem/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstOrCreate(t *testing.T) {
	mockRepo:= new(mock.MockStatsItemlRepository)
	service := NewStatsItemService(mockRepo)
	assert := assert.New(t)

	item := &entities.StatisticsItem{}
	mockRepo.On("FirstOrCreate", item).Return(nil)

	err := service.FirstOrCreate(item)
	assert.Nil(err)

}
