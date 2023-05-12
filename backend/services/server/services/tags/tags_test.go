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

	newTag := &entities.Tag{
		TagName: "tag1",
	}

	mockRepoTags.On("Create", newTag).Return(nil)

	err := service.AddTag("tag1")

	assert.Nil(err, "no error")
}

func TestDeleteTag(t *testing.T) {
	mockRepoTags := new(mock.MockTagRepository)
	service := NewTagsService(mockRepoTags)
	assert := assert.New(t)

	mockRepoTags.On("Delete", "tag1").Return(nil)

	err := service.DeleteTag("tag1")

	assert.Nil(err, "no error")
}

func TestGetTag(t *testing.T) {
	mockRepoTags := new(mock.MockTagRepository)
	service := NewTagsService(mockRepoTags)
	assert := assert.New(t)
	tag := &entities.Tag{
		TagName: "tag1",
	}
	mockRepoTags.On("Get", "tag1").Return(tag, nil)

	got, err := service.Get("tag1")

	assert.Nil(err, "no error")
	assert.Equal(tag, got)
}

func TestGetTagsByTagNames(t *testing.T) {
	mockRepoTags := new(mock.MockTagRepository)
	service := NewTagsService(mockRepoTags)
	assert := assert.New(t)
	tagNames := []string{"tag1", "tag2"}
	tag := &[]entities.Tag{
		{TagName: "tag1"},
		{TagName: "tag2"},
	}
	mockRepoTags.On("GetTagsByTagNames", tagNames).Return(tag, nil)

	got, err := service.GetTagsByTagNames(tagNames)

	assert.Nil(err, "no error")
	assert.Equal(tag, got)
}

func TestListTagsName(t *testing.T) {
	mockRepoTags := new(mock.MockTagRepository)
	service := NewTagsService(mockRepoTags)
	assert := assert.New(t)
	tagNames := []string{"tag1", "tag2"}
	tag := &[]entities.Tag{
		{TagName: "tag1"},
		{TagName: "tag2"},
	}
	mockRepoTags.On("List").Return(tag, nil)

	got, err := service.ListTagsName()

	assert.Nil(err, "no error")
	assert.Equal(tagNames, got)
}
