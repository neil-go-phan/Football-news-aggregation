package repository

import (
	"server/entities"

	"gorm.io/gorm"
)

type CrawlerRepo struct {
	DB *gorm.DB
}

func NewCrawlerRepo(db *gorm.DB) *CrawlerRepo {
	return &CrawlerRepo{
		DB: db,
	}
}

func (repo *CrawlerRepo) Upsert(crawler *entities.Crawler) error {
	if repo.DB.Model(&crawler).Where("url = ?", crawler.Url).
	Updates(entities.Crawler{
		Url: crawler.Url,
		ArticleDiv: crawler.ArticleDiv,
		ArticleTitle: crawler.ArticleTitle,
		ArticleDescription: crawler.ArticleDescription,
		ArticleLink: crawler.ArticleLink,
		NextPage: crawler.NextPage,
		NetxPageType: crawler.NetxPageType,
	}).RowsAffected == 0 {
		crawler.RunEveryMin = 5
		repo.DB.Create(&crawler)
	}
	return nil
}

func (repo *CrawlerRepo) List() (*[]entities.Crawler, error) {
	configCrawlers := make([]entities.Crawler, 0)
	err := repo.DB.Find(&configCrawlers).Error
	if err != nil {
		return nil, err
	}
	return &configCrawlers, nil
}

func (repo *CrawlerRepo) Get(url string) (*entities.Crawler, error) {
	configCrawler := new(entities.Crawler)
	err := repo.DB.Where(map[string]interface{}{"url": url}).Find(&configCrawler).Error
	if err != nil {
		return nil, err
	}
	return configCrawler, nil
}

func (repo *CrawlerRepo) Delete(url string) error {
	configCrawler := new(entities.Crawler)
	err := repo.DB.Unscoped().Where("url = ?", url).Delete(&configCrawler).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *CrawlerRepo) UpdateRunEveryTime(crawler *entities.Crawler) error {
	err := repo.DB.Model(&crawler).Update("run_every_min", crawler.RunEveryMin).Error
	if err != nil {
		return err
	}
	return nil
}