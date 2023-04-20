package tagsservices

import (
	"server/entities"
	"server/repository"
)

type tagsService struct {
	tagRepo repository.TagRepository
}

func NewTagsService(tagRepo repository.TagRepository) *tagsService {
	tag := &tagsService{
		tagRepo: tagRepo,
	}
	return tag
}

func (s *tagsService) AddTag(newTags string) error {
	return s.tagRepo.AddTag(newTags)
}

func (s *tagsService) DeleteTag(tag string) error {
	return s.tagRepo.DeleteTag(tag)
}

func (s *tagsService) ListTags() entities.Tags {
	return s.tagRepo.ListTags()
}
