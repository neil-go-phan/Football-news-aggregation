package overviewitem

import (
	"server/entities"
	mock "server/services/overviewItem/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstOrCreate(t *testing.T) {
	mockRepo:= new(mock.MockOverviewItemRepository)
	service := NewOverviewItemService(mockRepo)
	assert := assert.New(t)

	item := &entities.OverviewItem{}
	mockRepo.On("FirstOrCreate", item).Return(nil)

	err := service.FirstOrCreate(item)
	assert.Nil(err)

}
