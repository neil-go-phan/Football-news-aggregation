package repository

import (
	"server/entities"
	"time"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ArticleRepo struct {
	DB *gorm.DB
}

func NewArticleRepo(db *gorm.DB) *ArticleRepo {
	ArticleRepo := &ArticleRepo{
		DB: db,
	}
	return ArticleRepo
}

func (repo *ArticleRepo) FirstOrCreate(article *entities.Article) error {
	err := repo.DB.FirstOrCreate(article, entities.Article{Title: article.Title, Link: article.Link}).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *ArticleRepo) Delete(id uint) error {
	err := repo.DB.Unscoped().Delete(&entities.Article{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *ArticleRepo) GetArticlesByIDs(ids []uint) ([]entities.Article, error) {
	articles := make([]entities.Article, 0)
	err := repo.DB.Preload("Tags").Where("id IN ?", ids).Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (repo *ArticleRepo) AddTag(ids []uint, newTag *entities.Tag) error {
	log.Println("ids len", len(ids))
	articles := make([]entities.Article, 0)
	// get article need to add tag
	err := repo.DB.Where("id IN ?", ids).Find(&articles).Error
	if err != nil {
		return err
	}
	log.Println("articles len", len(articles))
	for _, article := range articles {
		err = repo.DB.Model(&article).Association("Tags").Append([]entities.Tag{*newTag})
		if err != nil {
			log.Error(err)
		}
	}

	return nil
}

func (repo *ArticleRepo) RemoveTag(ids []uint, newTag *entities.Tag) error {
	articles := make([]entities.Article, 0)
	// get article need to remove tag
	err := repo.DB.Where("id IN ?", ids).Find(&articles).Error
	if err != nil {
		return err
	}

	for _, article := range articles {
		err = repo.DB.Model(&article).Association("Tags").Delete([]entities.Tag{*newTag})
		if err != nil {
			log.Error(err)
		}
	}

	return nil
}

func (repo *ArticleRepo) GetCrawledArticleToday() (int64, error) {
	today := time.Now().Format("2006-01-02")
	var count int64

	err := repo.DB.Model(&entities.Article{}).Where("created_at BETWEEN ? AND ?", today+" 00:00:00", today+" 23:59:59").Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (repo *ArticleRepo) GetTotalCrawledArticle() (int64, error) {
	var count int64
	err := repo.DB.Model(&entities.Article{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
