package playerservice

import (
	"server/entities"
	mock "server/services/player/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstOrCreate(t *testing.T) {
	mockRepo:= new(mock.MockPlayerRepository)
	service := NewPlayerService(mockRepo)
	assert := assert.New(t)

	item := &entities.Player{}
	mockRepo.On("FirstOrCreate", item).Return(nil)

	err := service.FirstOrCreate(item)
	assert.Nil(err)

}
