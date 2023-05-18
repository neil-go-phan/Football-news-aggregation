package repository

import (
	"server/entities"

	"gorm.io/gorm"
)

type ConfigCrawlerRepo struct {
	DB *gorm.DB
}

func NewConfigCrawlerRepo(db *gorm.DB) *ConfigCrawlerRepo {
	return &ConfigCrawlerRepo{
		DB: db,
	}
}

func (repo *ConfigCrawlerRepo) Upsert(configCrawler *entities.ConfigCrawler) error {
	if repo.DB.Model(&configCrawler).Where("url = ?", configCrawler.Url).Updates(&configCrawler).RowsAffected == 0 {
		repo.DB.Create(&configCrawler)
	}
	return nil
}

func (repo *ConfigCrawlerRepo) List() (*[]entities.ConfigCrawler, error) {
	configCrawlers := make([]entities.ConfigCrawler, 0)
	err := repo.DB.Find(&configCrawlers).Error
	if err != nil {
		return nil, err
	}
	return &configCrawlers, nil
}

func (repo *ConfigCrawlerRepo) Get(url string) (*entities.ConfigCrawler, error) {
	configCrawler := new(entities.ConfigCrawler)
	err := repo.DB.Where(map[string]interface{}{"url": url}).Find(&configCrawler).Error
	if err != nil {
		return nil, err
	}
	return configCrawler, nil
}

func (repo *ConfigCrawlerRepo) Delete(url string) error {
	configCrawler := new(entities.ConfigCrawler)
	err := repo.DB.Unscoped().Where("url = ?", url).Delete(&configCrawler).Error
	if err != nil {
		return err
	}
	return nil
}
