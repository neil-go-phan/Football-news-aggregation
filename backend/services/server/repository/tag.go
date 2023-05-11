package repository

import (
	"server/entities"

	"gorm.io/gorm"
)

type tagRepo struct {
	DB *gorm.DB
}

func NewTagRepo(db *gorm.DB) *tagRepo {
	return &tagRepo{
		DB: db,
	}
}

func (repo *tagRepo) Create(tag *entities.Tag) error {
	err := repo.DB.Create(tag).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *tagRepo) List() (*[]entities.Tag, error) {
	tags := make([]entities.Tag, 0)
	err := repo.DB.Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return &tags, nil
}

func (repo *tagRepo) Get(tagName string) (*entities.Tag,error) {
	tag := new(entities.Tag)
	err := repo.DB.Where(map[string]interface{}{"tag_name": tagName}).Find(&tag).Error
	if err != nil {
		return nil, err
	}
	return tag, nil
}

func (repo *tagRepo) Delete(tagName string) error {
	tag := new(entities.Tag)
	err := repo.DB.Unscoped().Where("tag_name = ?", tagName).Delete(&tag).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *tagRepo) GetTagsByTagNames(tagNames []string) (*[]entities.Tag, error) {
	tags := make([]entities.Tag, 0)
	err := repo.DB.Where("tag_name IN ?", tagNames).Find(&tags).Error
	if err != nil {
		return nil ,err
	}
	return &tags, nil
}