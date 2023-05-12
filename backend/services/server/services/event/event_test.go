package eventservice

import (
	"server/entities"
	mock "server/services/event/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstOrCreate(t *testing.T) {
	mockRepoEvent := new(mock.MockEventRepository)
	service := NewEventService(mockRepoEvent)
	assert := assert.New(t)

	event := &entities.MatchEvent{}
	mockRepoEvent.On("FirstOrCreate", event).Return(nil)

	err := service.FirstOrCreate(event)
	assert.Nil(err)

}
