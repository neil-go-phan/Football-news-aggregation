package tagsservices

import (
	"server/entities"
	"server/repository"
)

type TagsService struct {
	tagRepo repository.TagRepository
}

func NewTagsService(tagRepo repository.TagRepository) *TagsService {
	tag := &TagsService{
		tagRepo: tagRepo,
	}
	return tag
}

func (s *TagsService) AddTag(newTags string) error {
	newTag := &entities.Tag{
		TagName: newTags,
	}
	return s.tagRepo.Create(newTag)
}

func (s *TagsService) DeleteTag(tagName string) error {
	return s.tagRepo.Delete(tagName)
}

func (s *TagsService) ListTagsName() ([]string, error) {
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

func (s *TagsService) GetTagsByTagNames(tagNames []string) (*[]entities.Tag, error)  {
	return s.tagRepo.GetTagsByTagNames(tagNames)
}

func (s *TagsService) Get(tagName string) (*entities.Tag,error)  {
	return s.tagRepo.Get(tagName)
}