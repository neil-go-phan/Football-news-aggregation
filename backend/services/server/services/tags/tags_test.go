package tagsservices

import (
	"testing"

	"server/entities"
	mock "server/services/tags/mock"

	"github.com/stretchr/testify/assert"
)

func TestAddTag(t *testing.T) {
	mockRepoTags := new(mock.MockTagRepository)
	service := NewTagsService(mockRepoTags)
	assert := assert.New(t)

	mockRepoTags.On("AddTag", "tag1").Return(nil)

	err := service.AddTag("tag1")

	assert.Nil(err, "no error")
}

func TestDeleteTag(t *testing.T) {
	mockRepoTags := new(mock.MockTagRepository)
	service := NewTagsService(mockRepoTags)
	assert := assert.New(t)

	mockRepoTags.On("DeleteTag", "tag1").Return(nil)

	err := service.DeleteTag("tag1")

	assert.Nil(err, "no error")
}

func TestListTag(t *testing.T) {
	mockRepoTags := new(mock.MockTagRepository)
	service := NewTagsService(mockRepoTags)
	assert := assert.New(t)

	want := entities.Tags{
		Tags: []string{"tag1", "tag2", "tag3"},
	}
	mockRepoTags.On("ListTags").Return(want)

	got := service.ListTags()

	assert.Equal(want, got)
}