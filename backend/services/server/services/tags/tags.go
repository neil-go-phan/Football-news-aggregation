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
	newTag := &entities.Tag{
		TagName: newTags,
	}
	return s.tagRepo.Create(newTag)
}

func (s *tagsService) DeleteTag(tagName string) error {
	return s.tagRepo.Delete(tagName)
}

func (s *tagsService) ListTagsName() ([]string, error) {
	tagNames := make([]string,0)
	tags, err := s.tagRepo.List() 
	if err != nil {
		return tagNames, err
	}

	for _, tag := range *tags {
		tagNames = append(tagNames, tag.TagName)
	}
	return tagNames, nil
}

func (s *tagsService) GetTagsByTagNames(tagNames []string) (*[]entities.Tag, error)  {
	return s.tagRepo.GetTagsByTagNames(tagNames)
}

func (s *tagsService) Get(tagName string) (*entities.Tag,error)  {
	return s.tagRepo.Get(tagName)
}